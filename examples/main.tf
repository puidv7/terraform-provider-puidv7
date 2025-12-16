terraform {
  required_providers {
    puidv7 = {
      source = "puidv7/puidv7"
    }
  }
}

provider "puidv7" {}

# Generate a new puidv7 ID
resource "puidv7_id" "account" {
  prefix = "acc"
}

output "account_id" {
  value = puidv7_id.account.id
}

output "account_uuid" {
  value = puidv7_id.account.uuid
}

# Encode an existing UUID to puidv7 format
data "puidv7_encode" "example" {
  uuid   = "01970a1c-e31e-7422-9cd5-e9651d11cc97"
  prefix = "acc"
}

output "encoded_id" {
  value = data.puidv7_encode.example.id
}

# Decode a puidv7 back to UUID
data "puidv7_decode" "example" {
  id     = "acc06bgm7733st2576nx5jht4ecjw"
  prefix = "acc"
}

output "decoded_uuid" {
  value = data.puidv7_decode.example.uuid
}
