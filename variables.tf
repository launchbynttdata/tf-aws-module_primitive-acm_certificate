variable "domain_name" {
  description = "The primary FQDN for the certificate."
  type        = string
  validation {
    condition     = can(regex("^[a-zA-Z0-9][a-zA-Z0-9.-]+[a-zA-Z0-9]$", var.domain_name))
    error_message = "The domain_name must be a valid FQDN."
  }
}

variable "key_algorithm" {
  description = "The key algorithm to use for the certificate. Default is 'RSA_2048'."
  type        = string
  default     = "RSA_2048"
  validation {
    condition     = contains(["RSA_2048", "EC_prime256v1", "EC_secp384r1", "EC_secp521r1"], var.key_algorithm)
    error_message = "Key algorithm must be one of 'RSA_2048', 'EC_prime256v1', 'EC_secp384r1', or 'EC_secp521r1'."
  }
}

variable "validation_method" {
  description = "The validation method for the certificate. Default is 'DNS'."
  type        = string
  default     = "DNS"
  validation {
    condition     = var.validation_method == "DNS" || var.validation_method == "EMAIL"
    error_message = "Validation method must be either 'DNS' or 'EMAIL'."
  }
}

variable "options" {
  description = "Options for the ACM certificate, such as certificate transparency logging preference."
  type        = map(string)
  default = {
    "certificate_transparency_logging_preference" = "ENABLED"
  }
  validation {
    condition     = contains(keys(var.options), "certificate_transparency_logging_preference")
    error_message = "Options must include 'certificate_transparency_logging_preference'."
  }
  validation {
    condition     = var.options["certificate_transparency_logging_preference"] == "ENABLED" || var.options["certificate_transparency_logging_preference"] == "DISABLED"
    error_message = "certificate_transparency_logging_preference must be either 'ENABLED' or 'DISABLED'."
  }
}

variable "validation_option" {
  description = "A map of validation options for the certificate, such as DNS records."
  type        = map(string)
  default     = null
}

variable "subject_alternative_names" {
  description = "A list of subject alternative names for the certificate."
  type        = list(string)
  default     = []
}

variable "tags" {
  description = "A map of tags to assign to the resource."
  type        = map(string)
  default     = {}
}
