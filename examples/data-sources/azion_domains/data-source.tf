terraform {
  required_providers {
    azion = {
      source  = "github.com/actions/azion"
    }
  }
}

provider "azion" {
  api_token  = "<token>"
}

data "azion_domains" "dev" {
}

output "dev_domains" {
  value = data.azion_domains.dev
}