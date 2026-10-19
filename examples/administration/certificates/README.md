# TLS Certificate Authority

This example demonstrates creating, updating, and deleting a `smc_tls_certificate_authority` resource.
A local test CA certificate (`ca/ca.crt`) is imported into the SMC as a trusted CA element.

The test CA is generated on demand by `ensure_ca.sh` (invoked automatically by `make plan`,
`make apply` and `make validate`) and removed by `make destroy`. The `ca/` directory is
git-ignored — it is a throwaway dev-only credential.
