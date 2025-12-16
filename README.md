# Terraform Provider puidv7

Generate and work with [puidv7](https://puidv7.dev) identifiers in Terraform.

A puidv7 is a prefixed UUIDv7 which:
- Is encoded using Crockford Base32 encoding
- Is always lowercase
- Does not contain any hyphens
- Is prefixed with a 3-character alphabetic (a-z) prefix

## Installation

```hcl
terraform {
  required_providers {
    puidv7 = {
      source = "puidv7/puidv7"
    }
  }
}

provider "puidv7" {}
```

## Usage

### Generate a new puidv7

```hcl
resource "puidv7_id" "account" {
  prefix = "acc"
}

output "account_id" {
  value = puidv7_id.account.id
  # Example: acc06bgm7733st2576nx5jht4ecjw
}

output "account_uuid" {
  value = puidv7_id.account.uuid
  # Example: 01970a1c-e31e-7422-9cd5-e9651d11cc97
}
```

The ID is generated once and stored in state. It will only change if you destroy and recreate the resource.

### Encode an existing UUID

```hcl
data "puidv7_encode" "example" {
  uuid   = "01970a1c-e31e-7422-9cd5-e9651d11cc97"
  prefix = "acc"
}

output "encoded" {
  value = data.puidv7_encode.example.id
  # acc06bgm7733st2576nx5jht4ecjw
}
```

### Decode a puidv7 to UUID

```hcl
data "puidv7_decode" "example" {
  id     = "acc06bgm7733st2576nx5jht4ecjw"
  prefix = "acc"  # Optional: validates the prefix matches
}

output "uuid" {
  value = data.puidv7_decode.example.uuid
  # 01970a1c-e31e-7422-9cd5-e9651d11cc97
}
```

## Building

```bash
go build -o terraform-provider-puidv7
```

## Local Development

To use the provider locally without publishing:

1. Build the provider
2. Create `~/.terraformrc`:

```hcl
provider_installation {
  dev_overrides {
    "puidv7/puidv7" = "/path/to/terraform-provider-puidv7"
  }
  direct {}
}
```

## License

puidv7 is licensed under the Apache License, Version 2.0.
Copyright 2025 Nadrama Pty Ltd.
See the [LICENSE](./LICENSE) file for details.
