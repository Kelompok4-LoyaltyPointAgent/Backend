variable "project" {
  type = string
}

variable "region" {
  type = string
}

variable "zone" {
  type = string
}

variable "mysql_root_password" {
  type = string
}

variable "env_vars_staging" {
  type = list(object({
    value = string
    name  = string
  }))
}
