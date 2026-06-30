#!/usr/bin/env bash
# Generuje parę kluczy OpenPGP dla dominik@labudzinski.com i eksportuje klucz publiczny.
#
# Po uruchomieniu:
#   1. Skopiuj static/public-key.asc do repozytorium (już zapisany przez skrypt).
#   2. Dodaj sekrety w GitHub → Settings → Secrets and variables → Actions:
#      - PGP_PRIVATE_KEY  — zawartość pliku private-key.asc (wygenerowany poniżej)
#      - PGP_PASSPHRASE   — hasło klucza (puste, jeśli klucz bez hasła)
#   3. Uruchom workflow lub lokalnie: make sign-posts
#
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
EMAIL="dominik@labudzinski.com"
NAME="Dominik Łabudziński"
BATCH_FILE="$(mktemp)"
PRIVATE_EXPORT="${ROOT}/private-key.asc"
PUBLIC_EXPORT="${ROOT}/static/public-key.asc"

cleanup() {
  rm -f "$BATCH_FILE"
}
trap cleanup EXIT

cat >"$BATCH_FILE" <<EOF
Key-Type: EDDSA
Key-Curve: Ed25519
Name-Real: ${NAME}
Name-Email: ${EMAIL}
Expire-Date: 0
EOF

echo "Generowanie klucza PGP dla ${EMAIL}..."
echo "Podaj hasło klucza (zalecane) lub zostaw puste dla klucza CI bez hasła:"
gpg --batch --generate-key "$BATCH_FILE"

mkdir -p "${ROOT}/static"
gpg --armor --export "$EMAIL" >"$PUBLIC_EXPORT"
gpg --armor --export-secret-keys "$EMAIL" >"$PRIVATE_EXPORT"
chmod 600 "$PRIVATE_EXPORT"

FPR="$(gpg --fingerprint --with-colons "$EMAIL" | awk -F: '/^fpr:/ {print $10; exit}')"
FPR_FMT="$(echo "$FPR" | tr '[:lower:]' '[:upper:]' | sed -E 's/(..)/\1 /g' | sed 's/ $//')"

echo
echo "Klucz publiczny:  ${PUBLIC_EXPORT}"
echo "Klucz prywatny:   ${PRIVATE_EXPORT}  (NIE commituj tego pliku)"
echo "Fingerprint:      ${FPR_FMT}"
echo
echo "GitHub Secrets:"
echo "  PGP_PRIVATE_KEY = zawartość ${PRIVATE_EXPORT}"
echo "  PGP_PASSPHRASE  = hasło klucza (jeśli ustawione)"
echo
echo "Weryfikacja lokalna:"
echo "  export PGP_PRIVATE_KEY=\"\$(cat ${PRIVATE_EXPORT})\""
echo "  export PGP_PASSPHRASE='twoje-haslo'"
echo "  make sign-posts"
