# Getting Started

This is an example demonstrating how to start using the provider on a
simple example with a single firewall. See [here](https://registry.terraform.io/providers/forcepoint-nsp/fp-ngfw-smc/latest/docs)

Pre-requisites:
- git installed
- terraform (version >=14) installed

Steps:

## Get an API key

- Open the SMC console and navigate to `Administration/Access Rights/API Clients`
- Create a new API key. For now just use "unrestricted permissions". 
- Copy the API key (a string that looks like this: mSs8tBu9naqyLGh9j5YiV7kA)

## Enable the SMC API

- Navigate to Dashboard/Server-Devices
- Edit the properties of your Management Server (secondary click/Properties).
- Go to "Client API" tab and click enable. For now just use HTTP.

## Clone the Examples github repository

```bash
git clone https://github.com/Forcepoint/terraform-provider-fp-ngfw-smc
```

## Export environment variables

```bash
export TF_VAR_url=http://localhost:8082
export TF_VAR_api_key="mSs8tBu9naqyLGh9j5YiV7kA"
export TF_VAR_api_version="7.4"
```

## Apply the configuration

```bash
cd terraform-provider-fp-ngfw-smc/examples/getting_started/
terraform init
terraform plan
terraform apply
```

## Check the result

You can now reopen the SMC console. It should have created:
- a single firewall with a physical interface
- a default route
- a policy that allows all the traffic to the firewall

It should have performed the following actions:
- bind a license to the firewall
- initiate an initial contact with the firewall and save the file
  containing the OTP to /tmp/${local.tf_single_fw_name}_initial_contact.txt

## Destroy the configuration

Delete all the created resources:

```bash
terraform destroy
```
