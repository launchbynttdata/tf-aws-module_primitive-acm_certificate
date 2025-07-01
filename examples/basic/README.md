# Basic

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 5.14 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | 5.100.0 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_aws_acm_certificate"></a> [aws\_acm\_certificate](#module\_aws\_acm\_certificate) | ../.. | n/a |
| <a name="module_cert_validation_records"></a> [cert\_validation\_records](#module\_cert\_validation\_records) | terraform.registry.launch.nttdata.com/module_primitive/dns_record/aws | ~> 1.0 |

## Resources

| Name | Type |
|------|------|
| [aws_acm_certificate_validation.cert_validation](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/acm_certificate_validation) | resource |
| [aws_route53_zone.sandbox](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/route53_zone) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_domain_name"></a> [domain\_name](#input\_domain\_name) | The primary FQDN for the certificate. | `string` | n/a | yes |
| <a name="input_route53_zone_id"></a> [route53\_zone\_id](#input\_route53\_zone\_id) | The Route53 zone ID where the certificate will be validated. | `string` | `""` | no |
| <a name="input_key_algorithm"></a> [key\_algorithm](#input\_key\_algorithm) | The key algorithm to use for the certificate. Default is 'RSA\_2048'. | `string` | `"RSA_2048"` | no |
| <a name="input_validation_method"></a> [validation\_method](#input\_validation\_method) | The validation method for the certificate. Default is 'DNS'. | `string` | `"DNS"` | no |
| <a name="input_options"></a> [options](#input\_options) | Options for the ACM certificate, such as certificate transparency logging preference. | `map(string)` | <pre>{<br/>  "certificate_transparency_logging_preference": "ENABLED"<br/>}</pre> | no |
| <a name="input_validation_option"></a> [validation\_option](#input\_validation\_option) | A map of validation options for the certificate, such as DNS records. | `map(string)` | `null` | no |
| <a name="input_subject_alternative_names"></a> [subject\_alternative\_names](#input\_subject\_alternative\_names) | A list of subject alternative names for the certificate. | `list(string)` | `[]` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to assign to the resource. | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_certificate_arn"></a> [certificate\_arn](#output\_certificate\_arn) | The ARN of the ACM certificate. |
<!-- END_TF_DOCS -->
