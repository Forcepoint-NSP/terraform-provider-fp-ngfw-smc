# Config Package

This package provides configuration loading and access functions for the Terraform SMC provider.

## Configuration File

The configuration file is named `tfsmc.conf.json` and should be placed in the local directory (current working directory).

## Configuration Format

The configuration supports two formats for the `hcl2` setting:

### Boolean Format

Enable or disable HCL2 for all resource types:

```json
{
  "hcl2": true
}
```

or

```json
{
  "hcl2": false
}
```

### Array Format

Enable HCL2 for specific resource types only:

```json
{
  "hcl2": ["smc_single_fw", "smc_routing_node"]
}
```

## Usage

### Loading Configuration

Load the configuration at application startup:

```go
import "github.com/terraform-providers/terraform-provider-smc/internal/config"

func main() {
    // Load config from tfsmc.conf.json in current directory
    if err := config.LoadConfig(); err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }
    
    // ... rest of your code
}
```

### Checking HCL2 Status

Check if HCL2 is enabled for a specific resource type:

```go
import "github.com/terraform-providers/terraform-provider-smc/internal/config"

func processResource(resourceType string) {
    if config.IsHcl2Enabled(resourceType) {
        // Use HCL2 processing
        fmt.Printf("Using HCL2 for %s\n", resourceType)
    } else {
        // Use legacy processing
        fmt.Printf("Using legacy processing for %s\n", resourceType)
    }
}
```

