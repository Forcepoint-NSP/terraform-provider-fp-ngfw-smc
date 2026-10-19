# VPN Certificate Authority

This example demonstrates creating, updating, and deleting a `smc_vpn_certificate_authority` resource.
A local test CA certificate (`ca/ca.crt`) is imported into the SMC as a VPN CA element,
usable for validating externally signed internal gateway (VPN) certificates.

The test CA is generated on demand by `ensure_ca.sh` (invoked automatically by `make plan`,
`make apply` and `make validate`) and removed by `make destroy`. The `ca/` directory is
git-ignored — it is a throwaway dev-only credential.

Note: to sign a gateway certificate via the `generate_certificate` operation, the CA must be
imported as a `smc_tls_certificate_authority` instead — see `../imported_ca`.
