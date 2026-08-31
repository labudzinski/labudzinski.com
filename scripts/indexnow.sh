#!/usr/bin/env bash
# Submit live sitemap URLs to IndexNow (Bing and participating engines).
set -euo pipefail

HOST="${INDEXNOW_HOST:-labudzinski.com}"
KEY="${INDEXNOW_KEY:-8f4c2a91e6b07d3f5c1a9e8b4d6f0c2a}"
BASE="https://${HOST}"
KEY_LOCATION="${BASE}/${KEY}.txt"
SITEMAP="${BASE}/sitemap.xml"
ENDPOINT="${INDEXNOW_ENDPOINT:-https://api.indexnow.org/indexnow}"

xml=""
for attempt in 1 2 3 4 5; do
  if xml="$(curl -fsSL --retry 2 "$SITEMAP")"; then
    break
  fi
  echo "indexnow: oczekiwanie na sitemap (próba ${attempt})" >&2
  sleep 15
done
if [ -z "$xml" ]; then
  echo "indexnow: nie udało się pobrać ${SITEMAP}" >&2
  exit 1
fi
urls="$(printf '%s' "$xml" | python3 -c "
import sys, re, json
xml = sys.stdin.read()
locs = re.findall(r'<loc>(.*?)</loc>', xml)
print(json.dumps(locs))
")"

python3 - "$ENDPOINT" "$HOST" "$KEY" "$KEY_LOCATION" "$urls" <<'PY'
import json, sys, urllib.request

endpoint, host, key, key_location, urls_json = sys.argv[1:6]
urls = json.loads(urls_json)
if not urls:
    raise SystemExit("indexnow: sitemap nie zawiera URL-i")

payload = json.dumps({
    "host": host,
    "key": key,
    "keyLocation": key_location,
    "urlList": urls,
}).encode()

req = urllib.request.Request(
    endpoint,
    data=payload,
    headers={"Content-Type": "application/json; charset=utf-8"},
    method="POST",
)
with urllib.request.urlopen(req, timeout=30) as resp:
    print(f"indexnow: {resp.status} {len(urls)} URLs")
PY
