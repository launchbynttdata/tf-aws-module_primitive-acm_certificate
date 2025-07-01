data "aws_route53_zone" "sandbox" {
  zone_id      = var.route53_zone_id
  private_zone = false
}
