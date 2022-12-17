variable "project" {
  type = string
}

variable "region" {
  type = string
}

variable "zone" {
  type = string
}

variable "env_vars_staging" {
  type = list(object({
    value = string
    name  = string
  }))
}

variable "env_vars_production" {
  type = list(object({
    value = string
    name  = string
  }))
}
