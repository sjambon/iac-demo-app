# Modules

Terraform modules zijn herbruikbare eenheden van Terraform-configuraties die helpen bij het organiseren en hergebruiken van infrastructuurcode. Een module bestaat uit één of meerdere Terraform-bestanden en heeft meestal dezelfde structuur als onze bovenliggende map, zie [readme](../README.md#opstelling).

Om een module uit een subfolder te gebruiken:

```hcl
module "my-function-module-name" {
  source = "./function_app"
  ...
}
```

#### functies en meta-argumenten die handig zijn
- [coalesce](https://developer.hashicorp.com/terraform/language/functions/coalesce)
- [merge](https://developer.hashicorp.com/terraform/language/functions/merge)
- [count](https://developer.hashicorp.com/terraform/language/meta-arguments/count) *handig voor conditioneel resource blocks te maken*
- [lower](https://developer.hashicorp.com/terraform/language/functions/lower)
- [replace](https://developer.hashicorp.com/terraform/language/functions/replace)
- [substr](https://developer.hashicorp.com/terraform/language/functions/substr)
