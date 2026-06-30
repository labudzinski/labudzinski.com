#!/usr/bin/env bash
# Weryfikuje podpis PGP jednego posta (treść markdown po front matter).
set -euo pipefail

POST="${1:?usage: $0 content/posts/example.md}"
ASC="${POST}.asc"

python3 - "$POST" "$ASC" <<'PY'
import subprocess
import sys

post_path, asc_path = sys.argv[1:3]
with open(post_path, encoding="utf-8") as f:
    text = f.read()

if not text.startswith("---\n"):
    raise SystemExit("missing front matter")

rest = text[4:]
end = rest.index("\n---")
body = rest[end + 4 :].lstrip("\n")
canonical = body.rstrip("\n") + "\n"

proc = subprocess.run(
    ["gpg", "--batch", "--verify", asc_path, "-"],
    input=canonical,
    text=True,
    capture_output=True,
)
sys.stdout.write(proc.stdout)
sys.stderr.write(proc.stderr)
raise SystemExit(proc.returncode)
PY
