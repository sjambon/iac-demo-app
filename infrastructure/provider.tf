terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 4.1.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "~> 3.6.3"
    }
    archive = {
      source  = "hashicorp/archive"
      version = "2.6.0"
    }
  }
  required_version = ">= 1.9.0"
}

# Configure the Azure provider
provider "azurerm" {
  features {}
}

