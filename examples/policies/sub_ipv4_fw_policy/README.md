# Sub IPv4 FW Policy Example

Shows use of sub-policy (jump) in firewall policies for modular rule structure.

- Creates two networks, a sub-policy with FTP allow, and a main policy with SSH jump rule referencing the sub-policy

## Usage

Customize network and rules as desired. Configure authentication, then:

```
terraform init
terraform apply
```
