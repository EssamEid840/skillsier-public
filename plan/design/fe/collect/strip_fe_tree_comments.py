#!/usr/bin/env python3
"""
drop_comment_only_tree_lines.py
--------------------------------
Remove ONLY the lines that contain *no file/folder name* and are just
tree guides followed by a '#' comment, e.g.:

    │   │   │   │   │   │   │       # BE: utility/status   <-- will be REMOVED
    │   │   │   │   │   ├── status/                        <-- will be KEPT
    │   │   │   │   │   │   ├── current/                   <-- will be KEPT
    │   │   │   │   │   │   │   └── page.tsx  # comment    <-- will be KEPT

Default scope:
- Only processes inside fenced code blocks (```), so headings/prose stay intact.

Options:
- --everywhere        : Process the whole file (not just inside ``` fences).
- --drop-guide-only   : Also drop lines that are ONLY tree guides (no text, no '#').
- -o, --output        : Write to a file; otherwise prints to stdout.

Usage:
  python drop_comment_only_tree_lines.py combined-fe-folder-strucure.md -o cleaned.md
  python drop_comment_only_tree_lines.py combined-fe-folder-strucure.md --everywhere
"""

import argparse
import re
import sys
from pathlib import Path

# Leading "tree guide" tokens seen in these docs (Unicode + ASCII)
GUIDE_PREFIX_RE = re.compile(r'^[\s│├┤└┌┐┘┬┴┼─—\+\|\-]*')

def payload_after_guides(line: str) -> str:
    """Return the content after removing leading tree guides and whitespace."""
    return GUIDE_PREFIX_RE.sub("", line).lstrip()

def is_comment_only_tree_line(line: str) -> bool:
    """
    True iff, after removing leading tree guides/whitespace,
    the line starts with '#' (i.e., it's only a comment line).
    """
    payload = payload_after_guides(line)
    return payload.startswith("#")

def is_guide_only(line: str) -> bool:
    """True iff, after removing guides, nothing remains (blank payload)."""
    return payload_after_guides(line) == ""

def process(text: str, only_inside_code_fences: bool = True, drop_guide_only: bool = False) -> str:
    out = []
    in_code = False

    for raw in text.splitlines():
        line = raw.rstrip("\n")

        # Toggle on triple-backtick fences
        if line.lstrip().startswith("```"):
            out.append(line)
            in_code = not in_code
            continue

        if in_code or not only_inside_code_fences:
            if is_comment_only_tree_line(line):
                # Drop comment-only tree line
                continue
            if drop_guide_only and is_guide_only(line):
                # Optionally drop pure guide/spacing lines
                continue
            # Keep everything else verbatim (folders/files with or without comments)
            out.append(line)
        else:
            out.append(line)

    # Preserve trailing newline if original had one
    return "\n".join(out) + ("\n" if text.endswith("\n") else "")

def main():
    ap = argparse.ArgumentParser(description="Drop comment-only tree lines; keep file/folder lines (with their comments).")
    ap.add_argument("input", nargs="?", help="Input file path (reads stdin if omitted)")
    ap.add_argument("-o", "--output", help="Output file path (prints to stdout if omitted)")
    ap.add_argument("--everywhere", action="store_true",
                    help="Process lines outside of ``` code fences as well.")
    ap.add_argument("--drop-guide-only", action="store_true",
                    help="Also drop lines that are only tree guides with no text.")
    args = ap.parse_args()

    data = Path(args.input).read_text(encoding="utf-8") if args.input else sys.stdin.read()

    cleaned = process(
        data,
        only_inside_code_fences=not args.everywhere,
        drop_guide_only=args.drop_guide_only,
    )

    if args.output:
        Path(args.output).write_text(cleaned, encoding="utf-8")
        print(f"Wrote: {args.output}")
    else:
        sys.stdout.write(cleaned)

if __name__ == "__main__":
    main()
