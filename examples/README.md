# Terraform SMC Provider Examples

This directory contains various examples demonstrating how to use the Terraform SMC (Secure Management Center) provider.

## Prerequisites

1. **SMC Server**: You need access to a Forcepoint SMC server
2. **API Access**: Ensure you have valid API credentials
3. **Terraform**: Install Terraform (version compatible with the provider)
4. **direnv**: Install [direnv](https://direnv.net/) for environment variable management

## Quick Start

### 1. Choose an Example

Navigate to any example directory:
- `fw_cluster/` - Firewall cluster configuration
- `fw_policy/` - Firewall policy management
- `host/` - Host management
- `import/` - Import existing resources
- `single_fw/` - Single firewall configuration

### 2. Configure Environment Variables

The simplest approach is to create a `.envrc` file in the `examples/` directory. This file will be shared by all examples using `direnv`.

Create a `.envrc` file in the `examples/` directory with the following content:

```bash
export TF_LOG_PATH=/tmp/tf.log
export TF_LOG_PROVIDER=DEBUG
export TF_VAR_url="http://localhost:8082"
export TF_VAR_api_key="xxxxx"
export TF_VAR_api_version="7.3"
export TF_LOG_SDK=OFF
```

Replace `xxxxx` with your actual SMC API key and adjust the API version as needed.

Then allow direnv to load the environment:

```bash
cd examples
direnv allow
```

The configuration is now set once and shared by all examples.

### 3. Initialize and Apply

```bash
# Navigate to your chosen example
cd <example-directory>

# Initialize Terraform with the local provider
make clean

# Review the planned changes
make plan

# Apply the configuration
make apply

# Clean up when done
make destroy
```

## Available Make Targets

Each example directory includes a Makefile with the following targets:

- `make help` - Show available commands
- `make init` - Initialize Terraform with local plugin
- `make plan` - Show planned changes
- `make apply` - Apply the configuration
- `make destroy` - Destroy the infrastructure
- `make clean` - Clean Terraform files and state

## Examples Overview

### Host Management (`host/`)
Demonstrates how to manage host objects in SMC, including creating and configuring host entries.

### Single Firewall (`single_fw/`)
Shows configuration of a single firewall instance with basic policies and settings.

### Firewall Cluster (`fw_cluster/`)
Example of setting up and managing a firewall cluster configuration.

### Firewall Policy (`fw_policy/`)
Demonstrates firewall policy creation and management.

### Import (`import/`)
Shows how to import existing SMC resources into Terraform state for management.

## Configuration Variables

All examples use the same provider configuration variables, set via environment variables:

| Environment Variable | Description | Example | Required |
|---------------------|-------------|---------|----------|
| `TF_VAR_api_key` | SMC API authentication key | `"your-api-key"` | Yes |
| `TF_VAR_url` | SMC API endpoint URL | `"http://localhost:8082"` | No |
| `TF_VAR_api_version` | SMC API version | `"7.3"` | No |

Additional logging variables:

| Environment Variable | Description | Example |
|---------------------|-------------|---------|
| `TF_LOG_PATH` | Path to Terraform log file | `/tmp/tf.log` |
| `TF_LOG_PROVIDER` | Provider log level | `DEBUG` |
| `TF_LOG_SDK` | SDK log level | `OFF` |

## Security Notes

- **Never commit** `.envrc` files containing sensitive data to version control
- The `.envrc` file is already included in `.gitignore`
- Use `direnv` to automatically load environment variables when entering the examples directory
- For production deployments, consider using secure secret management solutions


## Troubleshooting

### Common Issues

1. **Connection Error**: Verify the SMC server URL and ensure the server is accessible
2. **Authentication Error**: Check that your API key is valid and has appropriate permissions
3. **Version Mismatch**: Ensure the API version matches your SMC server version

### Getting Help

- Check the provider documentation
- Review Terraform logs with `TF_LOG=DEBUG terraform plan`
- Ensure your SMC server is properly configured for API access

## Development

When developing new examples:

1. Follow the same structure as existing examples
2. Include a `provider.tf` file with variable definitions
3. Use the common Makefile by including `../../makefile.common`
4. Ensure your `.envrc` is configured in the `examples/` directory

## Contributing

When adding new examples, please:
- Follow the established patterns
- Update this README with the new example description
- Test the complete workflow using the `.envrc` configuration
