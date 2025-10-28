#!/usr/bin/env python3
import re, sys, json, csv
from collections import Counter
from pathlib import Path

LINE_RE = re.compile(r'^(?P<prefix>[ \t│]*)(?P<edge>[├└])── (?P<name>[^#\n]+?)(?:\s{2,}#.*)?$')
INVALID = {"│", "…", "...", "—", "─", ""}

def depth(prefix: str) -> int:
    return (len(prefix.replace("\t", "    ")) // 4) + 1

def canonicalize(p: str) -> str:
    parts = [x for x in p.split("/") if x]
    i = 0
    while i < len(parts) and parts[i] == "fe":
        i += 1
    if i > 1:
        parts = ["fe"] + parts[i:]
    return "/".join(parts)

def parse_files(text: str):
    lines = text.splitlines()
    try:
        root_idx = next(i for i, ln in enumerate(lines) if ln.strip() == "fe/")
    except StopIteration:
        raise SystemExit("No 'fe/' root line found.")
    files, parents = [], {0: "fe"}
    for ln in lines[root_idx+1:]:
        m = LINE_RE.match(ln.rstrip())
        if not m: 
            continue
        name = m.group("name").rstrip()
        if name in INVALID:
            continue
        is_dir = name.endswith("/")
        name = name[:-1] if is_dir else name
        full = f'{parents.get(depth(m.group("prefix"))-1, "fe")}/{name}'
        full = re.sub(r"/{2,}", "/", full)
        if is_dir:
            parents[depth(m.group("prefix"))] = full
        else:
            files.append(full)
    return files

def write(path: Path, content: str): path.write_text(content, encoding="utf-8")
def write_csv(path: Path, rows, header):
    with open(path, "w", newline="", encoding="utf-8") as f:
        w = csv.writer(f); w.writerow(header); w.writerows(rows)

def main(tree_path: str):
    text = Path(tree_path).read_text(encoding="utf-8")
    files = parse_files(text)
    raw = Counter(files)
    canon = Counter(canonicalize(p) for p in files)

    dups_raw = sorted([p for p, c in raw.items() if c > 1])
    dups_canon = sorted([p for p, c in canon.items() if c > 1])

    summary = {
        "total_files_raw": len(files),
        "unique_files_raw": len(raw),
        "duplicate_file_paths_raw": len(dups_raw),
        "total_files_canonical": len(files),
        "unique_files_canonical": len(canon),
        "duplicate_file_paths_canonical": len(dups_canon),
    }

    print(json.dumps(summary, indent=2))
    write(Path("duplicate-summary.txt"), json.dumps(summary, indent=2))
    write(Path("duplicate-files.txt"), "\n".join(dups_raw) if dups_raw else "(no duplicate file paths found)")
    write(Path("duplicate-files-canonical.txt"), "\n".join(dups_canon) if dups_canon else "(no canonical duplicate file paths found)")
    write_csv(Path("duplicate-files.csv"), [(p, raw[p]) for p in dups_raw], ["file_path","count"])
    write_csv(Path("duplicate-files-canonical.csv"), [(p, canon[p]) for p in dups_canon], ["file_path","count"])

if __name__ == "__main__":
    path = sys.argv[1] if len(sys.argv) > 1 else "merged-fe-tree.md"
    main(path)
