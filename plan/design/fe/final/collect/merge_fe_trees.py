#!/usr/bin/env python3
"""
merge_fe_trees.py
-----------------

Merge multiple Markdown "folder structure" trees into a single continuous tree rooted at `fe/`.

Updates in this version:
- **Multi-segment path support**: Lines like
    `apps/web/src/app/[locale]/(dashboard)/`
  are parsed as a full path (split on '/'), and children (e.g., `search/`) are
  correctly placed under that path hierarchy.
- Preserves: canonical placement, folders-first sorting, deduplication,
  comment merging, file/folder collision handling, and platform variant retention.

Usage:
    python merge_fe_trees.py --base base.md a.md b.md -o merged-fe-tree.md
"""

from __future__ import annotations

import argparse
import re
import sys
from dataclasses import dataclass, field
from typing import Dict, List, Optional, Tuple


# ---------- Helpers: ordered unique list (preserve first-seen order) ----------

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

    def __contains__(self, item: str):
        return item in self._set

    def to_list(self) -> List[str]:
        return list(self._list)


# ---------- Node models ----------

@dataclass
class FileNode:
    name: str
    comments: OrderedSet = field(default_factory=OrderedSet)
    conflict: bool = False  # if a file-vs-folder name collision happened


@dataclass
class DirNode:
    name: str
    comments: OrderedSet = field(default_factory=OrderedSet)
    dirs: Dict[str, "DirNode"] = field(default_factory=dict)
    files: Dict[str, FileNode] = field(default_factory=dict)

    def get_or_add_dir(self, name: str) -> "DirNode":
        key = name
        if key not in self.dirs:
            self.dirs[key] = DirNode(name=name)
        return self.dirs[key]

    def get_or_add_file(self, name: str) -> FileNode:
        key = name
        if key not in self.files:
            self.files[key] = FileNode(name=name)
        return self.files[key]


# ---------- Canonical placement ----------

TOP_SYNONYMS = {
    # ci/cd & hooks
    ".github": ".github",
    ".husky": ".husky",
    ".vscode": ".vscode",
    # shared libs
    "pkg": "packages",
    "pkgs": "packages",
    "libs": "packages",
    "lib": "packages",
    "shared": "packages",
    "packages": "packages",
    # applications
    "apps": "apps",
    "web": ("apps", "web"),
    "mobile": ("apps", "mobile"),
    # other top-levels
    "docs": "docs",
    "documentation": "docs",
    "scripts": "scripts",
    "automation": "scripts",
    "config": "config",
    "configuration": "config",
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


def canonicalize_path(parts: List[str], enable: bool = True) -> List[str]:
    """
    Ensure path starts with 'fe' and normalize top-level placements.
    """
    if not parts:
        return ["fe"]

    parts = [p for p in parts if p and p != "."]

    # Drop root aliases like 'fe' if present
    if parts and parts[0] in FE_ALIASES:
        parts = parts[1:]

    out: List[str] = ["fe"]
    if not enable:
        return out + parts

    if not parts:
        return out

    head = parts[0]

    # apps/web or apps/mobile already
    if head == "apps" and len(parts) >= 2 and parts[1] in ("web", "mobile"):
        return out + ["apps", parts[1]] + parts[2:]

    # direct web or mobile
    if head in ("web", "mobile"):
        return out + ["apps", head] + parts[1:]

    mapped = TOP_SYNONYMS.get(head)
    if mapped is None:
        return out + [head] + parts[1:]

    if isinstance(mapped, tuple):
        return out + list(mapped) + parts[1:]

    return out + [mapped] + parts[1:]


# ---------- Parsing of Markdown tree files ----------

# Match a tree line with optional box-drawing prefix.
ITEM_RE = re.compile(
    r"""
    ^(?P<prefix>[\s│\u2502\u250C\u2514\u251C\u2500\u2510\u2524\u2534\u252C\u253C]*)
    (?:(?:├── |└── )\s*)?
    (?P<rest>.*\S.*)$
    """,
    re.VERBOSE,
)

def split_name_and_inline_comment(rest: str) -> Tuple[str, List[str], bool]:
    """
    Split the item text into name and comment(s).
    Returns: (name, [comment_lines], is_dir)
    """
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

    # Directory if ends with "/"
    is_dir = name_part.endswith("/")
    name = name_part[:-1] if is_dir else name_part
    name = name.strip()

    return name, inline_comments, is_dir


def parse_markdown_tree(md_text: str) -> List[Tuple[List[str], bool, List[str]]]:
    """
    Parse a markdown tree into a list of entries:
        - (parts, is_dir, comment_lines)

    Enhancements:
    - If an item "name" contains '/', we split it into path segments
      so 'apps/web/src/app/[locale]/(dashboard)/' creates nested dirs:
      ['apps','web','src','app','[locale]','(dashboard)']
    """
    entries: List[Tuple[List[str], bool, List[str]]] = []
    stack: List[List[str]] = []     # path parts per depth
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

        # Determine depth by counting groups of "│   " or 4 spaces in prefix
        prefix = m.group("prefix")
        depth = 0
        i = 0
        while i < len(prefix):
            if prefix.startswith("│   ", i):
                depth += 1
                i += 4
            elif prefix.startswith("    ", i):
                depth += 1
                i += 4
            else:
                i += 1

        name, inline_comments, is_dir = split_name_and_inline_comment(rest)

        # NEW: split multi-segment paths (e.g., "apps/web/src/...")
        name_segments = [seg for seg in name.split("/") if seg]

        parent = stack[depth - 1] if depth > 0 and depth - 1 < len(stack) else []
        parts = list(parent) + name_segments

        entries.append((parts, is_dir, inline_comments))
        last_idx = len(entries) - 1

        set_depth(depth, parts)

    return entries


# ---------- Merge logic ----------

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
        if name in parent.files:
            file_node = parent.files.pop(name)
            dir_node = parent.get_or_add_dir(name)
            conflict_file = dir_node.get_or_add_file(name)
            conflict_file.comments.add("conflict: file also existed")
            for c in file_node.comments:
                conflict_file.comments.add(c)
        dir_node = parent.get_or_add_dir(name)
        for c in comments:
            dir_node.comments.add(c)
    else:
        if name in parent.dirs:
            dir_node = parent.dirs[name]
            conflict_file = dir_node.get_or_add_file(name)
            conflict_file.conflict = True
            conflict_file.comments.add("conflict: file also existed")
            for c in comments:
                conflict_file.comments.add(c)
        else:
            f = parent.get_or_add_file(name)
            for c in comments:
                f.comments.add(c)


def merge_entries_into_tree(entries: List[Tuple[List[str], bool, List[str]]],
                            root: DirNode,
                            canonicalize: bool = True):
    for parts, is_dir, comments in entries:
        canon = canonicalize_path(parts, enable=canonicalize)
        add_entry(root, canon, is_dir, comments)


# ---------- Rendering ----------

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

    def draw_dir(node: DirNode, prefix: str, is_root: bool = False):
        name = node.name + "/"
        if is_root:
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

    # Ensure we render under 'fe'
    if root.name != "fe":
        wrapper = DirNode(name="fe")
        wrapper.dirs[root.name] = root
        root = wrapper

    fake_root = DirNode(name="fe")
    fake_root.dirs = root.dirs
    fake_root.files = root.files
    fake_root.comments = root.comments

    draw_dir(fake_root, prefix="", is_root=True)

    if as_markdown:
        return "```\n" + "\n".join(lines) + "\n```"
    else:
        return "\n".join(lines)


# ---------- CLI ----------

def read_text(path: str) -> str:
    with open(path, "r", encoding="utf-8") as f:
        return f.read()


def main():
    ap = argparse.ArgumentParser(description="Merge multiple Markdown folder structure trees into one 'fe/' tree.")
    ap.add_argument("--base", required=True, help="Path to the BASE tree markdown file (included in full, merged first).")
    ap.add_argument("-o", "--output", help="Write merged tree to this file (markdown code block). If omitted, prints to stdout.")
    ap.add_argument("--no-markdown", action="store_true", help="Do not wrap output in a markdown code block.")
    ap.add_argument("--no-canonicalize", action="store_true", help="Disable canonical placement normalization (rare).")
    ap.add_argument("additives", nargs="*", help="Additional markdown tree files to merge (order matters).")

    args = ap.parse_args()

    try:
        base_text = read_text(args.base)
    except Exception as e:
        print(f"ERROR: failed to read base file '{args.base}': {e}", file=sys.stderr)
        sys.exit(2)

    all_entries: List[Tuple[List[str], bool, List[str]]] = []
    base_entries = parse_markdown_tree(base_text)
    all_entries.extend(base_entries)

    for add_path in args.additives:
        try:
            add_text = read_text(add_path)
        except Exception as e:
            print(f"ERROR: failed to read additive file '{add_path}': {e}", file=sys.stderr)
            sys.exit(2)
        add_entries = parse_markdown_tree(add_text)
        all_entries.extend(add_entries)

    root = DirNode(name="fe")
    merge_entries_into_tree(entries=all_entries, root=root, canonicalize=(not args.no_canonicalize))

    tree_str = draw_tree(root, as_markdown=(not args.no_markdown))

    if args.output:
        with open(args.output, "w", encoding="utf-8") as f:
            f.write(tree_str)
        print(f"Wrote merged tree to: {args.output}")
    else:
        print(tree_str)


if __name__ == "__main__":
    main()
