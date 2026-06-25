# Custom Domain

This is an example demonstrating the usage of the Custom Domain
configuration. 

workflow:

- create an alias provider with indication of the domain
- create the domain using smc_admin_domain
- create resources in the newly created domain

Note that all the resources created in the newly domain must add an
explicit dependency to the domain creation. You may want to place all
the domain resources in a terraform module to simplify this.
