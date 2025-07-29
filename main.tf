resource "aws_acm_certificate" "cert" {
  domain_name               = var.domain_name
  validation_method         = var.validation_method
  key_algorithm             = var.key_algorithm
  subject_alternative_names = var.subject_alternative_names != null ? var.subject_alternative_names : []

  dynamic "options" {
    for_each = try(var.options, [])
    content {
      certificate_transparency_logging_preference = lookup(var.options, "certificate_transparency_logging_preference", "ENABLED")
    }
  }
  dynamic "validation_option" {
    for_each = try(var.validation_option[*], [])
    content {
      domain_name       = lookup(var.validation_option, "domain_name", var.domain_name)
      validation_domain = lookup(var.validation_option, "validation_domain", var.domain_name)
    }
  }

  tags = merge(
    {
      name = join("-", [var.domain_name, "acm-cert"])
    },
    var.tags
  )

  lifecycle {
    create_before_destroy = true
  }
}
