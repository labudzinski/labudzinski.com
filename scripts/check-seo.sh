#!/usr/bin/env bash
# ponytail: one assert that home sitemap lastmod tracks the latest post.
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"
hugo --minify --gc >/dev/null
python3 - <<'PY'
import re, sys
from pathlib import Path
xml = Path("public/sitemap.xml").read_text()
posts = re.findall(r"<loc>https://labudzinski.com/posts/[^<]+/</loc><lastmod>([^<]+)</lastmod>", xml)
home = re.search(r"<loc>https://labudzinski.com/</loc><lastmod>([^<]+)</lastmod>", xml)
if not home:
    sys.exit("check-seo: missing home url in sitemap")
if not posts:
    sys.exit("check-seo: no posts in sitemap")
latest = max(posts)
if home.group(1) != latest:
    sys.exit(f"check-seo: home lastmod {home.group(1)} != latest post {latest}")
html = Path("public/index.html").read_text()
if "max-image-preview:large" not in html:
    sys.exit("check-seo: robots meta missing max-image-preview")
if 'hreflang=x-default' not in html and 'hreflang="x-default"' not in html:
    sys.exit("check-seo: missing hreflang x-default")
print(f"check-seo: ok (home lastmod={home.group(1)})")
PY
