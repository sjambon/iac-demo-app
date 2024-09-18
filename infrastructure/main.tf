locals {
  default_tags = merge(var.default_tags, {
    environment = var.environment,
    name = "steven"
  })
}

resource "azurerm_resource_group" "storage-group" {
  name     = "iac-demo-${var.environment}"
  location = var.location
}

resource "azurerm_storage_account" "storage-account" {
  name                     = "saniac"
  resource_group_name      = azurerm_resource_group.storage-group.name
  location                 = azurerm_resource_group.storage-group.location
  account_tier             = "Standard"
  account_replication_type = "LRS"

  tags = local.default_tags
}