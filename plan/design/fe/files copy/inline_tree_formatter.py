#!/usr/bin/env python3
import sys
from typing import List

SEP = " | "

def inline_tree(text: str, sep: str = SEP) -> str:
    lines = text.splitlines()
    out: List[str] = []
    last_idx = -1

    def is_path_line(s: str) -> bool:
        # Treat any line containing a tree-connector and a filename as a path line.
        return ("├── " in s or "└── " in s)

    for raw in lines:
        line = raw.rstrip("\n")
        if is_path_line(line):
            # Bring any trailing "# ..." on the same line, if present.
            if "#" in line:
                before, after = line.split("#", 1)
                base = before.rstrip()
                first_comment = after.strip()
                composed = f"{base}  # {first_comment}" if first_comment else base
            else:
                composed = line
            out.append(composed)
            last_idx = len(out) - 1
        elif "#" in line:
            # Continuation comment line — append to the most recent path line.
            if last_idx >= 0:
                _, after = line.split("#", 1)
                extra = after.strip()
                if extra:
                    out[last_idx] = f"{out[last_idx]}{sep}{extra}"
            else:
                # No prior path line; keep as-is.
                out.append(line)
        else:
            # Non-comment/connector line (e.g., folder headings) — keep it.
            out.append(line)

    return "\n".join(out) + ("\n" if out and not out[-1].endswith("\n") else "")

def main():
    if len(sys.argv) < 3:
        print("Usage: inline_tree_formatter.py <input.txt> <output.txt> [separator]", file=sys.stderr)
        sys.exit(1)
    src = sys.argv[1]
    dst = sys.argv[2]
    sep = sys.argv[3] if len(sys.argv) > 3 else SEP
    with open(src, "r", encoding="utf-8") as f:
        text = f.read()
    result = inline_tree(text, sep=sep)
    with open(dst, "w", encoding="utf-8") as f:
        f.write(result)

if __name__ == "__main__":
    main()
