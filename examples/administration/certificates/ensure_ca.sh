#!/usr/bin/env bash
# ensure_ca.sh — create the local test CA for an example if it does not exist.
#
# Runs in the current example directory ($PWD) and creates ./ca/ca.key and
# ./ca/ca.crt. Examples reference the certificate with
# file("${path.module}/ca/ca.crt"), so the Makefile runs this script before
# any terraform command that evaluates the configuration; `make destroy`
# removes the ca/ directory afterwards. The directory is git-ignored — the
# CA is a throwaway dev-only credential regenerated on demand.
#
# The SMC requires CA certificates to carry the X509v3 key identifier
# extensions, so they are declared explicitly instead of relying on
# openssl defaults.

set -euo pipefail

CA_DIR="$PWD/ca"
CA_KEY="$CA_DIR/ca.key"
CA_CRT="$CA_DIR/ca.crt"

if [ -f "$CA_KEY" ] && [ -f "$CA_CRT" ]; then
    exit 0
fi

echo "[ensure_ca] Creating test CA in $CA_DIR (valid 100 years)..."
mkdir -p "$CA_DIR"
openssl req -newkey rsa:4096 -nodes \
    -keyout "$CA_KEY" \
    -x509 -days 36500 \
    -out "$CA_CRT" \
    -subj "/CN=Forcepoint Test CA/O=Forcepoint/OU=Testing" \
    -addext "basicConstraints = critical, CA:TRUE" \
    -addext "keyUsage = critical, keyCertSign, cRLSign" \
    -addext "subjectKeyIdentifier = hash" \
    -addext "authorityKeyIdentifier = keyid:always, issuer"
chmod 600 "$CA_KEY"
echo "[ensure_ca] Test CA created: $CA_CRT"
