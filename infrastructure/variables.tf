variable "location" {
  description = "Location of the resource"
  type        = string
  default     = "westeurope"
}

variable "environment" {
  description = "Name of the environment of the resources"
  type        = string
}

variable "default_tags" {
  type        = map(string)
  description = "set of default tags to apply everywhere"
}