output "certificate_arn" {
  description = "The ARN of the ACM certificate."
  value       = aws_acm_certificate.cert.arn
}
output "domain_validation_options" {
  description = "The domain validation options for the ACM certificate."
  value       = aws_acm_certificate.cert.domain_validation_options
}
