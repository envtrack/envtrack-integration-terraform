# Terraform Provider for EnvTrack

This Terraform provider allows you to manage environment variables and their ownership in EnvTrack directly from your Terraform configurations.

## Features

- Push environment variables to EnvTrack
- Track variable ownership across projects and environments
- Integrate with multiple cloud providers and services

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.13.x or later
- [Go](https://golang.org/doc/install) 1.16 or later (for development)

## Installation

To use this provider, add the following code to your Terraform configuration:

```hcl
terraform {
  required_providers {
    envtrack = {
      source = "envtrack/envtrack"
    }
  }
}

provider "envtrack" {
  auth_token = var.envtrack_auth_token
}
```

### Usage

Here's a basic example of how to use the EnvTrack provider:
```hcl
resource "envtrack_track" "example" {
  organization_id = "your-org-id"
  project_id      = "your-project-id"
  environment_id  = "your-environment-id"
  var_identifier  = "unique-identifier"

  input_data = {
    DB_HOST = "localhost"
    DB_PORT = "5432"
  }
}
```

### Development

To build the provider:

- Clone the repository
- Enter the repository directory
- Build the provider using go build command

```sh
git clone https://github.com/your-username/terraform-provider-envtrack.git
cd terraform-provider-envtrack
go build
```

### Local development

To use the provider locally, you can use the `terraform init` command with the `plugin-dir` flag.

Building the provider locally:

```sh
go build -o terraform-provider-envtrack .
mkdir -p ~/.terraform.d/plugins/local.providers/local/envtrack/1.0.4/darwin_arm64/
cp terraform-provider-envtrack ~/.terraform.d/plugins/local.providers/local/envtrack/1.0.3/darwin_arm64/
```

### Documentation

Full documentation is generated using tfplugindocs:
```sh
go generate ./...
```

### Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### License

This project is licensed under the MIT License.

