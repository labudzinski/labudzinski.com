---
date: 2026-06-30T13:30:59+02:00
draft: false
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJMBAABCgA2FiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpD2CMYHGRvbWluaWtA
  bGFidWR6aW5za2kuY29tAAoJEFzelGpBYMIp5d4QAJjb+3fVFYHG4qdbfOrcJAzX
  fXcYvdbBfIPiBlfxzVo0WKESI7jP6meoQJnK/dO1uVB2YtzyCULyi6rmlYakVi0f
  GbM4NSBHWNiKQgwFBya/1xqkuQxQ2DVjq/KHb2zyw9uWj8C7Iczu0npmPPl8CsDQ
  hVJFz6QWsGtbT4joNFdTYGE9cc1mJ7GxBpxkVBEDLplctxRVdGExXHKXQm/97t5U
  s+J+ZCzyGDmi2M1xXLlKZSAc/fVnWXvmJnuvdcWytTiFTZx+d/EFWuTRsSJ9D7/c
  jWdTKFREeCG5UghbBbL1QefXoKnY5hneeW4NANgoXPaYocmfDVEhldcg4E+nlHcS
  e/Y7Zy+tCNVEPrPGlYUh58ISV6ReAAuhmGbuJ1mdOFyjce12URXNqDxlgwZJ5uTW
  43Wo7aIQZBpbC5HhpKEqGmoI7aZRdhBhzPNw0y/csDXNmMvg0O6B1YwmlJ3yKzg0
  JzvJbnrhkUKf2WVEUgCHR8Nsq5h9Unj/ICYzdopNa2dJgmrKsAevPpvdNCuXcEGA
  PUX0GldjFE6yt+aKcYN1BIJyzzUl8r6XfK4z5evs+rOiIZMmgLx3K/Zh7JL0TrKe
  Xk1W7M+OUs+l4hm/kAKajC75EuUHL031I13B3gskkV8daMloBxnDY/AhlMoc8+KI
  1zdp+J7y2rPOP7qgCxeN
  =ZWxM
  -----END PGP SIGNATURE-----
tags:
  - backup
  - restic
  - macOS
  - hetzner
title: Mój backup
toc: false
---

Od dłuższego czasu miałem wrażenie, że temat kopii zapasowych odkładam na później bardziej niż powinienem. Pracuję na MacBooku, mam dużo projektów, trochę prywatnych plików, trochę rzeczy rozrzuconych po różnych katalogach i zawsze zakładałem, że „jakoś to będzie”. W praktyce oznaczało to brak spójnego systemu i poleganie na tym, że nic się nie stanie.

Ostatnio uporządkowałem ten temat i w końcu mam działający backup, który działa automatycznie i niezależnie ode mnie. Zdecydowałem się na Restic i zewnętrzny storage od Hetznera. Całość jest szyfrowana i działa przyrostowo, więc nie kopiuję wszystkiego od nowa, tylko zmieniające się fragmenty danych.

Najważniejsze dla mnie było to, żeby ten system był prosty w utrzymaniu. Nie chciałem rozbudowanych paneli, klikanych interfejsów ani usług, które trzeba regularnie kontrolować. Wystarczy mi jeden skrypt i jedno miejsce, w którym wiem, że są moje dane.

Poniżej zostawiam aktualny skrypt bash, którego używam. To jest wersja, którą uruchamiam automatycznie przez launchd na macOS. Pokazuje mi na bieżąco, co się dzieje, zapisuje logi i wysyła powiadomienie, kiedy backup się kończy lub kiedy coś pójdzie nie tak.

```bash
#!/usr/bin/env bash
#
# Automated restic backup for macOS (launchd).
# Configurable via environment variables; see readonly defaults below.
#
set -Eeuo pipefail
IFS=$'\n\t'

umask 077

readonly RESTIC_DIR="${RESTIC_DIR:-${HOME}/.restic}"
readonly PASSWORD_FILE="${RESTIC_PASSWORD_FILE:-${RESTIC_DIR}/password}"
readonly REPO="${RESTIC_REPOSITORY:-sftp:hetzner-backup:/home/backups/macbook}"
readonly EXCLUDE_FILE_DEFAULT="${RESTIC_EXCLUDE_FILE:-${RESTIC_DIR}/excludes.txt}"
readonly LOCK_DIR="${RESTIC_LOCK_DIR:-${RESTIC_DIR}/backup.lock.d}"
readonly LOG_FILE="${RESTIC_LOG_FILE:-${RESTIC_DIR}/backup.log}"
readonly LAST_CHECK_FILE="${RESTIC_DIR}/last-check"
readonly LOG_MAX_BYTES="${RESTIC_LOG_MAX_BYTES:-10485760}"

readonly KEEP_DAILY="${RESTIC_KEEP_DAILY:-7}"
readonly KEEP_WEEKLY="${RESTIC_KEEP_WEEKLY:-4}"
readonly KEEP_MONTHLY="${RESTIC_KEEP_MONTHLY:-6}"
readonly CHECK_INTERVAL_DAYS="${RESTIC_CHECK_INTERVAL_DAYS:-7}"

readonly -a BACKUP_PATHS=(
  "${HOME}/Documents"
  "${HOME}/IdeaProjects"
  "${HOME}/.ssh"
  "${HOME}/.config"
)

export PATH="/opt/homebrew/bin:/usr/local/bin:/usr/bin:/bin"
export RESTIC_PASSWORD_FILE="${PASSWORD_FILE}"

VALID_BACKUP_PATHS=()
EXCLUDE_FILE="${EXCLUDE_FILE_DEFAULT}"

USE_COLOR=false
if [[ -t 1 ]]; then
  USE_COLOR=true
fi

readonly C_GREEN=$'\033[0;32m'
readonly C_RED=$'\033[0;31m'
readonly C_YELLOW=$'\033[1;33m'
readonly C_BLUE=$'\033[0;34m'
readonly C_RESET=$'\033[0m'

log() {
  local level="$1"
  local message="$2"
  local prefix=""
  local plain=""
  local display=""
  local timestamp=""

  case "${level}" in
    info)  prefix="[INFO] " ;;
    warn)  prefix="[WARN] " ;;
    error) prefix="[ERROR] " ;;
  esac

  plain="${prefix}${message}"
  display="${plain}"

  if [[ "${USE_COLOR}" == true ]]; then
    case "${level}" in
      info)  display="${C_BLUE}${plain}${C_RESET}" ;;
      warn)  display="${C_YELLOW}${plain}${C_RESET}" ;;
      error) display="${C_RED}${plain}${C_RESET}" ;;
      ok)    display="${C_GREEN}${message}${C_RESET}" ;;
    esac
  fi

  timestamp="$(date '+%F %T')"
  printf '%s\n' "${display}"
  printf '%s %s\n' "${timestamp}" "${plain}" >>"${LOG_FILE}"
}

die() {
  log error "$1"
  exit "${2:-1}"
}

escape_applescript() {
  local value="$1"
  value="${value//\\/\\\\}"
  value="${value//\"/\\\"}"
  printf '%s' "${value}"
}

notify() {
  local title message
  title="$(escape_applescript "$1")"
  message="$(escape_applescript "$2")"

  if command -v osascript >/dev/null 2>&1; then
    osascript -e "display notification \"${message}\" with title \"${title}\"" >/dev/null 2>&1 || true
  fi
}

rotate_log_if_needed() {
  [[ -f "${LOG_FILE}" ]] || return 0

  local size=0
  size="$(wc -c <"${LOG_FILE}" | tr -d '[:space:]')"

  if (( size > LOG_MAX_BYTES )); then
    mv -f "${LOG_FILE}" "${LOG_FILE}.1"
    : >"${LOG_FILE}"
    log info "Rotated log file (${LOG_MAX_BYTES} byte limit)"
  fi
}

validate_password_file() {
  [[ -r "${PASSWORD_FILE}" ]] || die "Password file not readable: ${PASSWORD_FILE}"

  local mode=""
  mode="$(stat -f '%OLp' "${PASSWORD_FILE}" 2>/dev/null || stat -c '%a' "${PASSWORD_FILE}" 2>/dev/null || true)"
  if [[ -n "${mode}" && "${mode}" != "600" && "${mode}" != "400" ]]; then
    die "Password file permissions must be 600 or 400, got ${mode}: ${PASSWORD_FILE}"
  fi
}

preflight() {
  command -v restic >/dev/null 2>&1 || die "restic not found in PATH"

  mkdir -p "${RESTIC_DIR}"
  chmod 700 "${RESTIC_DIR}" 2>/dev/null || true

  validate_password_file

  if [[ -n "${EXCLUDE_FILE}" && ! -f "${EXCLUDE_FILE}" ]]; then
    log warn "Exclude file not found, continuing without excludes: ${EXCLUDE_FILE}"
    EXCLUDE_FILE=""
  fi
}

collect_backup_paths() {
  VALID_BACKUP_PATHS=()
  local path=""

  for path in "${BACKUP_PATHS[@]}"; do
    if [[ -e "${path}" ]]; then
      VALID_BACKUP_PATHS+=("${path}")
    else
      log warn "Skipping missing path: ${path}"
    fi
  done

  ((${#VALID_BACKUP_PATHS[@]} > 0)) || die "No valid backup paths found"
}

acquire_lock() {
  if mkdir "${LOCK_DIR}" 2>/dev/null; then
    printf '%s\n' "$$" >"${LOCK_DIR}/pid"
    return 0
  fi

  if [[ -f "${LOCK_DIR}/pid" ]]; then
    local old_pid=""
    old_pid="$(<"${LOCK_DIR}/pid")"
    if [[ "${old_pid}" =~ ^[0-9]+$ ]] && ! kill -0 "${old_pid}" 2>/dev/null; then
      log warn "Removing stale lock (pid ${old_pid})"
      rm -rf "${LOCK_DIR}"
      mkdir "${LOCK_DIR}" 2>/dev/null || die "Backup already running"
      printf '%s\n' "$$" >"${LOCK_DIR}/pid"
      return 0
    fi
  fi

  die "Backup already running (lock: ${LOCK_DIR})"
}

release_lock() {
  rm -rf "${LOCK_DIR}"
}

on_exit() {
  release_lock
}

on_interrupt() {
  log error "Backup interrupted"
  notify "Restic Backup" "Backup interrupted"
  exit 130
}

should_run_check() {
  if (( CHECK_INTERVAL_DAYS < 0 )); then
    return 1
  fi
  if (( CHECK_INTERVAL_DAYS == 0 )); then
    return 0
  fi
  if [[ ! -f "${LAST_CHECK_FILE}" ]]; then
    return 0
  fi

  local last_run now age_days
  last_run="$(<"${LAST_CHECK_FILE}")"
  now="$(date +%s)"
  age_days=$(( (now - last_run) / 86400 ))
  (( age_days >= CHECK_INTERVAL_DAYS ))
}

mark_check_completed() {
  date +%s >"${LAST_CHECK_FILE}"
}

run_restic() {
  local -a cmd=(restic -r "${REPO}" "$@")
  log info "Running: ${cmd[*]}"
  "${cmd[@]}"
}

run_backup() {
  local -a args=(
    backup
    --verbose=1
    --tag "host-$(hostname -s)"
    --tag automated
  )

  collect_backup_paths
  args+=("${VALID_BACKUP_PATHS[@]}")

  if [[ -n "${EXCLUDE_FILE}" ]]; then
    args+=(--exclude-file="${EXCLUDE_FILE}")
  fi

  run_restic "${args[@]}"
}

run_forget_prune() {
  run_restic forget \
    --keep-daily "${KEEP_DAILY}" \
    --keep-weekly "${KEEP_WEEKLY}" \
    --keep-monthly "${KEEP_MONTHLY}" \
    --prune
}

run_check() {
  run_restic check --read-data-subset=5%
  mark_check_completed
}

main() {
  local stage_failed=false

  trap on_interrupt INT TERM
  trap on_exit EXIT

  preflight
  rotate_log_if_needed
  acquire_lock

  log ok "====================================="
  log info "Starting backup: $(date)"
  log info "Repository: ${REPO}"
  log ok "====================================="

  log warn "[1/3] Running backup..."
  if ! run_backup; then
    log error "Backup stage failed"
    notify "Backup Failed" "Backup stage failed"
    exit 1
  fi

  log warn "[2/3] Applying retention policy..."
  if ! run_forget_prune; then
    log error "Retention stage failed"
    notify "Backup Warning" "Backup succeeded, but retention policy failed"
    stage_failed=true
  fi

  if should_run_check; then
    log warn "[3/3] Verifying repository..."
    if ! run_check; then
      log error "Repository check failed"
      notify "Backup Warning" "Backup succeeded, but repository check failed"
      stage_failed=true
    fi
  else
    log info "[3/3] Skipping repository check (interval: ${CHECK_INTERVAL_DAYS} days)"
  fi

  if [[ "${stage_failed}" == true ]]; then
    exit 2
  fi

  log ok "====================================="
  log ok "Backup completed successfully"
  log ok "====================================="
  notify "Backup Completed" "Backup completed successfully"
}

main "$@"
```

Nie traktuję tego jako rozwiązania idealnego. Raczej jako coś, co wreszcie daje mi spokój w tle. Kopie zapasowe przestały być tematem, o którym muszę pamiętać, a stały się procesem, który po prostu działa.
Być może jest ktoś, kto potrafiłby napisać lepszy skrypt, ale na razie ten mi wystarcza.
