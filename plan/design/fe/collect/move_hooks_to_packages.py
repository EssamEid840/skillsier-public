#!/usr/bin/env python3
"""
move_hooks_to_packages.py
-------------------------

Normalize hooks placement in a commented Markdown "folder structure" tree.

Moves hooks from:
  - fe/apps/<web|mobile>/src/**/<domain>/hooks/**        → fe/packages/hooks/<domain>/**
  - fe/shared/**/<domain>/hooks/**                       → fe/packages/hooks/<domain>/**
  - fe/packages/shared/**/<domain>/hooks/**              → fe/packages/hooks/<domain>/**
And (optionally):
  - fe/packages/<any>/**/<domain>/hooks/**               → fe/packages/hooks/<domain>/**  (when --promote-packages all)

Rules
- Keeps anything already in fe/packages/hooks/** intact.
- Detects <domain> as the path segment **immediately before** 'hooks', with fallback that
  skips known scaffolding segments like 'src', 'app', 'pages', 'shared', and strips route wrappers
  like '(dashboard)' or '[locale]'.
- Merges comments **uniquely** (first-seen order) if a destination file exists.
- By default, **does not** move files directly under apps/web/src/hooks/*; you can override.
- Prunes empty source directories after moving.

Usage
    python move_hooks_to_packages.py input.md -o output.md
    python move_hooks_to_packages.py input.md --dry-run
    python move_hooks_to_packages.py input.md -o out.md --promote-shared --promote-packages shared
    python move_hooks_to_packages.py input.md -o out.md --promote-packages all

CLI Options
    --dry-run                   Show planned moves; do not change the tree.
    --no-markdown               Output without ``` fences.
    --keep-web-hooks            Comma-separated filenames to KEEP under apps/web/src/hooks
                                (default: use-keycloak.ts,use-ssr-query.ts).
    --move-web-root-hooks       Also move files directly under apps/<app>/src/hooks (off by default).
    --only-app {web,mobile}     Restrict moves to a specific app.
    --promote-shared            Also move hooks under top-level 'shared/**' or 'packages/shared/**'.
    --promote-packages {none,shared,all}
                                Promote from packages/*:
                                  none    → do not promote from packages/*
                                  shared  → (default) promote only from packages/shared/*
                                  all     → promote from any packages/* (excluding packages/hooks/* itself)
"""

from __future__ import annotations

import argparse
import re
import sys
from dataclasses import dataclass, field
from typing import Dict, List, Optional, Tuple


# ---------- Ordered unique list (preserve first-seen order) ----------

class OrderedSet:
    def __init__(self):
        self._list: List[str] = []
        self._set = set()

    def add(self, item: str):
        if item not in self._set:
            self._set.add(item)
            self._list.append(item)

    def extend(self, items: List[str]):
        for it in items:
            self.add(it)

    def __iter__(self):
        return iter(self._list)

    def __len__(self):
        return len(self._list)

    def to_list(self) -> List[str]:
        return list(self._list)


# ---------- Node models ----------

@dataclass
class FileNode:
    name: str
    comments: OrderedSet = field(default_factory=OrderedSet)


@dataclass
class DirNode:
    name: str
    comments: OrderedSet = field(default_factory=OrderedSet)
    dirs: Dict[str, "DirNode"] = field(default_factory=dict)
    files: Dict[str, FileNode] = field(default_factory=dict)

    def get_or_add_dir(self, name: str) -> "DirNode":
        if name not in self.dirs:
            self.dirs[name] = DirNode(name=name)
        return self.dirs[name]

    def get_or_add_file(self, name: str) -> FileNode:
        if name not in self.files:
            self.files[name] = FileNode(name=name)
        return self.files[name]

    def remove_file(self, name: str):
        if name in self.files:
            del self.files[name]

    def remove_dir_if_empty(self, name: str) -> bool:
        """Remove a child dir if it (recursively) becomes empty. Returns True if removed."""
        if name not in self.dirs:
            return False
        d = self.dirs[name]
        # Clean children first
        for sub in list(d.dirs.keys()):
            d.remove_dir_if_empty(sub)
        if not d.files and not d.dirs:
            del self.dirs[name]
            return True
        return False


# ---------- Canonical placement ----------

TOP_SYNONYMS = {
    ".github": ".github",
    ".husky": ".husky",
    ".vscode": ".vscode",
    "pkg": "packages",
    "pkgs": "packages",
    "libs": "packages",
    "lib": "packages",
    "shared": "packages",  # top-level 'shared' canonicalizes under 'packages'
    "packages": "packages",
    "apps": "apps",
    "web": ("apps", "web"),
    "mobile": ("apps", "mobile"),
    "docs": "docs",
    "scripts": "scripts",
    "automation": "scripts",
    "config": "config",
    "infra": "infra",
    "infrastructure": "infra",
    "tools": "tools",
    "examples": "examples",
    "tests": "tests",
    "test": "tests",
    "assets": "assets",
    "public": "public",
}

FE_ALIASES = {"fe", "front-end", "frontend", "skillsier-fe", "skillsier_fe"}

def canonicalize_path(parts: List[str]) -> List[str]:
    """Ensure path starts with 'fe' and normalize top-level placements."""
    parts = [p for p in parts if p and p != "."]
    if parts and parts[0] in FE_ALIASES:
        parts = parts[1:]
    out: List[str] = ["fe"]
    if not parts:
        return out
    head = parts[0]
    if head == "apps" and len(parts) >= 2 and parts[1] in ("web", "mobile"):
        return out + ["apps", parts[1]] + parts[2:]
    if head in ("web", "mobile"):
        return out + ["apps", head] + parts[1:]
    mapped = TOP_SYNONYMS.get(head)
    if mapped is None:
        return out + [head] + parts[1:]
    if isinstance(mapped, tuple):
        return out + list(mapped) + parts[1:]
    return out + [mapped] + parts[1:]


# ---------- Parsing of Markdown tree files ----------

ITEM_RE = re.compile(
    r"""
    ^(?P<prefix>[\s│\u2502\u250C\u2514\u251C\u2500\u2510\u2524\u2534\u252C\u253C]*)
    (?:(?:├── |└── )\s*)?
    (?P<rest>.*\S.*)$
    """,
    re.VERBOSE,
)

def split_name_and_inline_comment(rest: str) -> Tuple[str, List[str], bool]:
    inline_comments: List[str] = []
    name_part = rest
    for sep in ["  #", " #", "\t#"]:
        if sep in rest:
            name_part, comment = rest.split(sep, 1)
            inline = comment.strip()
            if inline.startswith("#"):
                inline = inline.lstrip("#").strip()
            if inline:
                inline_comments.append(inline)
            break
    name_part = name_part.rstrip()
    is_dir = name_part.endswith("/")
    name = name_part[:-1] if is_dir else name_part
    name = name.strip()
    return name, inline_comments, is_dir


def parse_markdown_tree(md_text: str) -> List[Tuple[List[str], bool, List[str]]]:
    """
    Parse a markdown tree into entries of (parts, is_dir, comment_lines).
    Also splits multi-segment names containing '/' into path segments.
    Associates trailing '# ...' lines with the last item.
    """
    entries: List[Tuple[List[str], bool, List[str]]] = []
    stack: List[List[str]] = []
    last_idx: Optional[int] = None

    def set_depth(depth: int, parts: List[str]):
        while len(stack) > depth:
            stack.pop()
        if len(stack) == depth:
            stack.append(parts)
        else:
            stack[depth] = parts

    in_code = False
    for raw_line in md_text.splitlines():
        line = raw_line.rstrip("\n")
        fence = line.strip()
        if fence.startswith("```"):
            in_code = not in_code
            continue
        if not line.strip():
            continue

        m = ITEM_RE.match(line)
        if not m:
            stripped = line.lstrip(" │")
            if stripped.startswith("#") and last_idx is not None:
                extra = stripped.lstrip("#").strip()
                if extra:
                    parts, is_dir, comments = entries[last_idx]
                    comments.append(extra)
                    entries[last_idx] = (parts, is_dir, comments)
            continue

        rest = m.group("rest").strip()
        if rest.startswith("#"):
            if last_idx is not None:
                extra = rest.lstrip("#").strip()
                if extra:
                    parts, is_dir, comments = entries[last_idx]
                    comments.append(extra)
                    entries[last_idx] = (parts, is_dir, comments)
            continue

        prefix = m.group("prefix")
        depth = 0
        i = 0
        while i < len(prefix):
            if prefix.startswith("│   ", i) or prefix.startswith("    ", i):
                depth += 1
                i += 4
            else:
                i += 1

        name, inline_comments, is_dir = split_name_and_inline_comment(rest)
        name_segments = [seg for seg in name.split("/") if seg]
        parent = stack[depth - 1] if depth > 0 and depth - 1 < len(stack) else []
        parts = list(parent) + name_segments

        entries.append((parts, is_dir, inline_comments))
        last_idx = len(entries) - 1
        set_depth(depth, parts)

    return entries


# ---------- Tree helpers ----------

def ensure_path(root: DirNode, path: List[str]) -> DirNode:
    node = root
    for p in path:
        node = node.get_or_add_dir(p)
    return node


def add_entry(root: DirNode, path: List[str], is_dir: bool, comments: List[str]):
    if not path:
        return
    *parent_parts, name = path
    parent = ensure_path(root, parent_parts)
    if is_dir:
        d = parent.get_or_add_dir(name)
        for c in comments:
            d.comments.add(c)
    else:
        f = parent.get_or_add_file(name)
        for c in comments:
            f.comments.add(c)


def sort_ci(names: List[str]) -> List[str]:
    return sorted(names, key=lambda s: s.casefold())


def render_comments_inline_and_block(comment_set: OrderedSet) -> Tuple[str, List[str]]:
    comments = comment_set.to_list()
    if not comments:
        return "", []
    inline = comments[0]
    rest = comments[1:]
    return inline, rest


def draw_tree(root: DirNode, as_markdown: bool = True) -> str:
    lines: List[str] = []

    def emit(line: str):
        lines.append(line)

    def draw_dir(node: DirNode, prefix: str):
        name = node.name + "/"
        base = f"{name}"
        inline, rest = render_comments_inline_and_block(node.comments)
        if inline:
            base += f"  # {inline}"
        emit(base)
        for c in rest:
            emit(f"# {c}")

        dir_names = sort_ci(list(node.dirs.keys()))
        file_names = sort_ci(list(node.files.keys()))
        children = [("dir", n) for n in dir_names] + [("file", n) for n in file_names]

        total = len(children)
        for idx, (kind, name) in enumerate(children):
            last = idx == total - 1
            connector = "└── " if last else "├── "
            if kind == "dir":
                child = node.dirs[name]
                base = f"{prefix}{connector}{child.name}/"
                inline, rest = render_comments_inline_and_block(child.comments)
                if inline:
                    base += f"  # {inline}"
                emit(base)
                cp = f"{prefix}{'    ' if last else '│   '}"
                for c in rest:
                    emit(f"{cp}# {c}")
                draw_subtree(child, cp)
            else:
                file_node = node.files[name]
                base = f"{prefix}{connector}{file_node.name}"
                inline, rest = render_comments_inline_and_block(file_node.comments)
                if inline:
                    base += f"  # {inline}"
                emit(base)
                cp = f"{prefix}{'    ' if last else '│   '}"
                for c in rest:
                    emit(f"{cp}# {c}")

    def draw_subtree(node: DirNode, prefix: str):
        dir_names = sort_ci(list(node.dirs.keys()))
        file_names = sort_ci(list(node.files.keys()))
        children = [("dir", n) for n in dir_names] + [("file", n) for n in file_names]

        total = len(children)
        for idx, (kind, name) in enumerate(children):
            last = idx == total - 1
            connector = "└── " if last else "├── "
            if kind == "dir":
                child = node.dirs[name]
                base = f"{prefix}{connector}{child.name}/"
                inline, rest = render_comments_inline_and_block(child.comments)
                if inline:
                    base += f"  # {inline}"
                emit(base)
                cp = f"{prefix}{'    ' if last else '│   '}"
                for c in rest:
                    emit(f"{cp}# {c}")
                draw_subtree(child, cp)
            else:
                file_node = node.files[name]
                base = f"{prefix}{connector}{file_node.name}"
                inline, rest = render_comments_inline_and_block(file_node.comments)
                if inline:
                    base += f"  # {inline}"
                emit(base)
                cp = f"{prefix}{'    ' if last else '│   '}"
                for c in rest:
                    emit(f"{cp}# {c}")

    draw_dir(root, prefix="")
    return ("```\n" + "\n".join(lines) + "\n```") if as_markdown else "\n".join(lines)


# ---------- Moves logic ----------

ROUTE_WRAPPER_RE = re.compile(r"^[\(\[](.+)[\)\]]$")

SKIP_BEFORE_HOOKS = {"src", "app", "pages", "shared"}
APPS_SET = {"web", "mobile"}

def sanitize_segment(seg: str) -> str:
    """Strip route-group markers like (dashboard) or [locale]."""
    m = ROUTE_WRAPPER_RE.match(seg)
    return m.group(1) if m else seg

def is_under_packages(parts: List[str]) -> bool:
    return len(parts) >= 2 and parts[0] == "fe" and parts[1] == "packages"

def is_under_packages_shared(parts: List[str]) -> bool:
    return len(parts) >= 3 and parts[0] == "fe" and parts[1] == "packages" and parts[2] == "shared"

def is_under_top_shared(parts: List[str]) -> bool:
    return len(parts) >= 2 and parts[0] == "fe" and parts[1] == "shared"

def find_hooks_cases(root: DirNode) -> List[Tuple[List[str], FileNode]]:
    """
    Return a list of (path_parts, file_node) for any FILE under a 'hooks' dir
    that is NOT under fe/packages/** (i.e., apps, shared).
    """
    results: List[Tuple[List[str], FileNode]] = []

    def dfs(node: DirNode, path: List[str]):
        if path and path[-1] == "hooks":
            for fname, fnode in node.files.items():
                results.append((path + [fname], fnode))
        for dname, dnode in node.dirs.items():
            dfs(dnode, path + [dname])

    dfs(root, [])
    results = [(p, f) for (p, f) in results if not is_under_packages(p)]
    return results

def find_hooks_cases_in_packages_shared(root: DirNode) -> List[Tuple[List[str], FileNode]]:
    """Collect files under any 'hooks' dir inside packages/shared or top-level shared."""
    results: List[Tuple[List[str], FileNode]] = []

    def dfs(node: DirNode, path: List[str]):
        if path and path[-1] == "hooks":
            for fname, fnode in node.files.items():
                results.append((path + [fname], fnode))
        for dname, dnode in node.dirs.items():
            dfs(dnode, path + [dname])

    dfs(root, [])
    out = []
    for p, f in results:
        if is_under_packages_shared(p) or is_under_top_shared(p):
            if len(p) >= 3 and p[:3] == ["fe", "packages", "hooks"]:
                continue  # already canonical
            out.append((p, f))
    return out

def locate_dir(root: DirNode, parts: List[str]) -> Optional[DirNode]:
    node = root
    for seg in parts:
        if seg not in node.dirs:
            return None
        node = node.dirs[seg]
    return node

def remove_file_at(root: DirNode, parts: List[str]) -> None:
    *parent, fname = parts
    d = locate_dir(root, parent)
    if not d:
        return
    d.remove_file(fname)
    # prune upwards
    for depth in range(len(parent), 0, -1):
        ancestor_parent = locate_dir(root, parent[:depth-1]) if depth > 1 else root
        ancestor_name = parent[depth-1] if depth > 0 else None
        if ancestor_parent and ancestor_name:
            removed = ancestor_parent.remove_dir_if_empty(ancestor_name)
            if not removed:
                break

def _find_idx(parts: List[str], value: str) -> int:
    try:
        return parts.index(value)
    except ValueError:
        return -1

def _domain_from_before_hooks(parts: List[str], hooks_idx: int) -> Optional[str]:
    """Find a usable <domain> from segments before 'hooks', skipping scaffolding."""
    j = hooks_idx - 1
    while j >= 0 and sanitize_segment(parts[j]) in SKIP_BEFORE_HOOKS:
        j -= 1
    if j < 0:
        return None
    return sanitize_segment(parts[j])

def plan_destination_for_apps(parts: List[str], only_app: Optional[str], move_web_root_hooks: bool) -> Optional[List[str]]:
    """
    fe/apps/<app>/src/**/<domain>/hooks/[...]/file -> fe/packages/hooks/<domain>/[...]/file
    Special case: fe/apps/<app>/src/hooks/* → move only if move_web_root_hooks=True (domain 'web-root').
    """
    if not (len(parts) >= 4 and parts[0] == "fe" and parts[1] == "apps" and parts[2] in APPS_SET):
        return None
    app = parts[2]
    if only_app and app != only_app:
        return None

    idx = _find_idx(parts, "hooks")
    if idx < 0:
        return None

    # Directly under src/hooks?
    if idx >= 1 and parts[idx - 1] == "src":
        if not move_web_root_hooks:
            return None
        domain = "web-root" if app == "web" else f"{app}-root"
    else:
        domain = _domain_from_before_hooks(parts, idx)
        if not domain:
            return None

    rel_after_hooks = parts[idx+1:-1]
    fname = parts[-1]
    return ["fe", "packages", "hooks", domain] + rel_after_hooks + [fname]

def plan_destination_for_shared_or_packages(parts: List[str]) -> Optional[List[str]]:
    """
    For shared and packages/*:
      fe/(packages/)?shared/**/<domain>/hooks/[...]/file
      fe/packages/<any>/**/<domain>/hooks/[...]/file
    → fe/packages/hooks/<domain>/[...]/file
    """
    idx = _find_idx(parts, "hooks")
    if idx < 0:
        return None
    domain = _domain_from_before_hooks(parts, idx)
    if not domain:
        return None
    rel_after_hooks = parts[idx+1:-1]
    fname = parts[-1]
    return ["fe", "packages", "hooks", domain] + rel_after_hooks + [fname]

def merge_file_comments(dst_parent: DirNode, fname: str, src_comments: OrderedSet):
    f = dst_parent.get_or_add_file(fname)
    f.comments.extend(src_comments.to_list())  # unique, preserves first-seen order

def move_hooks(root: DirNode,
               keep_web_hooks: List[str],
               move_web_root_hooks: bool,
               only_app: Optional[str],
               dry_run: bool = False,
               promote_shared: bool = False,
               promote_packages: str = "shared") -> List[Tuple[str, str]]:
    """
    Execute moves and return a list of (src_str, dst_str).
    """
    moves_done: List[Tuple[str, str]] = []

    # --- 1) Apps (web/mobile)
    cases = find_hooks_cases(root)
    for parts, fnode in list(cases):
        parts = canonicalize_path(parts)

        # Skip canonical package location
        if len(parts) >= 3 and parts[:3] == ["fe", "packages", "hooks"]:
            continue

        # Skip kept files in apps/web/src/hooks
        if len(parts) >= 6 and parts[:5] == ["fe", "apps", "web", "src", "hooks"]:
            if fnode.name in keep_web_hooks and not move_web_root_hooks:
                continue

        dest = plan_destination_for_apps(parts, only_app=only_app, move_web_root_hooks=move_web_root_hooks)
        if not dest:
            continue

        *dst_parent_parts, dst_fname = dest
        dst_parent = ensure_path(root, dst_parent_parts)

        if dry_run:
            moves_done.append(("/" + "/".join(parts), "/" + "/".join(dest)))
            continue

        merge_file_comments(dst_parent, dst_fname, fnode.comments)
        remove_file_at(root, parts)
        moves_done.append(("/" + "/".join(parts), "/" + "/".join(dest)))

    # --- 2) Promote from shared (top-level or packages/shared)
    if promote_shared or promote_packages in {"shared", "all"}:
        shared_cases = find_hooks_cases_in_packages_shared(root)
        for parts, fnode in list(shared_cases):
            parts = canonicalize_path(parts)
            dest = plan_destination_for_shared_or_packages(parts)
            if not dest:
                continue
            *dst_parent_parts, dst_fname = dest
            dst_parent = ensure_path(root, dst_parent_parts)

            if dry_run:
                moves_done.append(("/" + "/".join(parts), "/" + "/".join(dest)))
                continue

            merge_file_comments(dst_parent, dst_fname, fnode.comments)
            remove_file_at(root, parts)
            moves_done.append(("/" + "/".join(parts), "/" + "/".join(dest)))

    # --- 3) Promote from ANY packages/* (excluding packages/hooks/*), if requested
    if promote_packages == "all":
        candidates: List[Tuple[List[str], FileNode]] = []

        def dfs(node: DirNode, path: List[str]):
            if path and path[-1] == "hooks":
                for fname, fnode in node.files.items():
                    candidates.append((path + [fname], fnode))
            for dname, dnode in node.dirs.items():
                dfs(dnode, path + [dname])

        dfs(root, [])

        for parts, fnode in list(candidates):
            parts = canonicalize_path(parts)

            # Skip canonical and packages/shared (already handled)
            if len(parts) >= 3 and parts[:3] == ["fe", "packages", "hooks"]:
                continue
            if is_under_packages_shared(parts):
                continue

            # Only consider those under packages/*
            if not (len(parts) >= 2 and parts[0] == "fe" and parts[1] == "packages"):
                continue

            dest = plan_destination_for_shared_or_packages(parts)
            if not dest:
                continue
            *dst_parent_parts, dst_fname = dest
            dst_parent = ensure_path(root, dst_parent_parts)

            if dry_run:
                moves_done.append(("/" + "/".join(parts), "/" + "/".join(dest)))
                continue

            merge_file_comments(dst_parent, dst_fname, fnode.comments)
            remove_file_at(root, parts)
            moves_done.append(("/" + "/".join(parts), "/" + "/".join(dest)))

    return moves_done


# ---------- CLI ----------

def read_text(path: str) -> str:
    with open(path, "r", encoding="utf-8") as f:
        return f.read()

def main():
    ap = argparse.ArgumentParser(description="Move hooks to 'fe/packages/hooks/<domain>' with unique comment merging.")
    ap.add_argument("input", help="Input markdown tree file (with comments)")
    ap.add_argument("-o", "--output", help="Output file (markdown code block). If omitted, prints to stdout.")
    ap.add_argument("--dry-run", action="store_true", help="Show the planned moves but do not alter the tree.")
    ap.add_argument("--no-markdown", action="store_true", help="Do not wrap output in a markdown code block.")
    ap.add_argument("--keep-web-hooks", default="use-keycloak.ts,use-ssr-query.ts",
                    help="Comma-separated filenames to keep in apps/web/src/hooks (default: %(default)s)")
    ap.add_argument("--move-web-root-hooks", action="store_true",
                    help="Also move files directly under apps/<app>/src/hooks (default: keep them).")
    ap.add_argument("--only-app", choices=["web", "mobile"], help="Only move hooks for a specific app (optional).")
    ap.add_argument("--promote-shared", action="store_true",
                    help="Also move hooks under 'shared/**' or 'packages/shared/**' into 'packages/hooks/<domain>'.")
    ap.add_argument("--promote-packages", choices=["none", "shared", "all"], default="shared",
                    help="Promote hooks from packages/*: 'none', 'shared' (default), or 'all'.")

    args = ap.parse_args()

    try:
        md_text = read_text(args.input)
    except Exception as e:
        print(f"ERROR: cannot read input file '{args.input}': {e}", file=sys.stderr)
        sys.exit(2)

    # Parse → build tree (canonicalized)
    entries = parse_markdown_tree(md_text)
    root = DirNode(name="fe")
    for parts, is_dir, comments in entries:
        canon = canonicalize_path(parts)
        add_entry(root, canon, is_dir, comments)

    keep_list = [s.strip() for s in args.keep_web_hooks.split(",") if s.strip()] if args.keep_web_hooks else []
    moves = move_hooks(
        root,
        keep_web_hooks=keep_list,
        move_web_root_hooks=args.move_web_root_hooks,
        only_app=args.only_app,
        dry_run=args.dry_run,
        promote_shared=args.promote_shared,
        promote_packages=args.promote_packages,
    )

    if args.dry_run:
        if moves:
            print("# Planned moves:")
            for src, dst in moves:
                print(f"- {src}  ->  {dst}")
        else:
            print("# No moves planned (nothing matched).")
        sys.exit(0)

    tree_str = draw_tree(root, as_markdown=(not args.no_markdown))
    if args.output:
        with open(args.output, "w", encoding="utf-8") as f:
            f.write(tree_str)
        print(f"Wrote updated tree to: {args.output}")
    else:
        print(tree_str)


if __name__ == "__main__":
    main()
