terraform {
  required_providers {
    azion = {
      source  = "hashicorp.com/dev/azion"
    }
  }
}

provider "azion" {
  api_token  = "<token>"
}

data "azion_records" "dev" {
  zone_id = 2638
}

output "dev_records" {
  value = data.azion_records.dev
}