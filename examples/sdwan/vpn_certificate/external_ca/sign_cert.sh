#!/usr/bin/env bash
# sign_cert.sh — sign a VPN certificate CSR with a local test CA.
#
# Usage:
#   ./sign_cert.sh <csr_file> <output_cert_file>
#
# The test CA lives in ./ca/ (git-ignored) and is created on demand by
# ensure_ca.sh. Keep ca/ca.key private; ca/ca.crt is imported into the SMC
# by the example configuration.

set -euo pipefail

CSR_FILE="${1:?Usage: $0 <csr_file> <output_cert_file>}"
OUT_FILE="${2:?Usage: $0 <csr_file> <output_cert_file>}"

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
CA_DIR="$SCRIPT_DIR/ca"
CA_KEY="$CA_DIR/ca.key"
CA_CRT="$CA_DIR/ca.crt"

# The CA is created by the Makefile (ensure_ca.sh) before terraform runs and
# the SMC CA element is imported from it in the same apply. Do NOT recreate it
# here: a regenerated CA would no longer match the element in the SMC and the
# certificate import would fail with "Impossible to find the signing
# certificate authority".
if [ ! -f "$CA_KEY" ] || [ ! -f "$CA_CRT" ]; then
    echo "[sign_cert] ERROR: $CA_DIR is missing. Run the example via make (apply runs ensure_ca.sh)," >&2
    echo "[sign_cert]        then re-apply so the SMC CA element and the local CA stay in sync." >&2
    exit 1
fi

# Sign the CSR (valid 10 years).
# The extension file is required: a bare `openssl x509 -req` produces an
# X.509 v1 certificate without extensions, and the SMC needs the
# Authority Key Identifier to match the certificate to its signing CA.
echo "[sign_cert] Signing $CSR_FILE ..."
EXT_FILE="$(mktemp)"
trap 'rm -f "$EXT_FILE"' EXIT
cat > "$EXT_FILE" <<'EOF'
basicConstraints = CA:FALSE
subjectKeyIdentifier = hash
authorityKeyIdentifier = keyid,issuer
keyUsage = critical, digitalSignature, keyEncipherment
EOF
openssl x509 -req \
    -in "$CSR_FILE" \
    -CA "$CA_CRT" \
    -CAkey "$CA_KEY" \
    -CAcreateserial \
    -out "$OUT_FILE" \
    -days 3650 \
    -sha512 \
    -extfile "$EXT_FILE"
echo "[sign_cert] Signed certificate written to $OUT_FILE"
