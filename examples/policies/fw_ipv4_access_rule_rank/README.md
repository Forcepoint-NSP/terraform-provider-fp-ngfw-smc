# IPv4 Access Rule Rank Example

Demonstrates ordering IPv4 access rules in a firewall policy using the `rank` argument. Creates a policy and several access rules with custom order.  

- Defines policy template and five ranked rules (SSH, HTTPS, FTP,
  ICMP, AOL)
- Shows source/destination/service/action per rule
- Automates policy/rule creation with Terraform

## Usage

Adjust rule/order as needed, set authentication, then:

```
terraform init
terraform apply
```
