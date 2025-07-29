# tf-aws-module_primitive-acm_certificate

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![License: CC BY-NC-ND 4.0](https://img.shields.io/badge/License-CC_BY--NC--ND_4.0-lightgrey.svg)](https://creativecommons.org/licenses/by-nc-nd/4.0/)

## Overview

This Terraform module provisions an AWS Certificate Manager (ACM) SSL/TLS certificate for securing domain communications. The module supports both single domain and multi-domain certificates with Subject Alternative Names (SANs), and provides flexible validation methods including DNS and email validation.

Key features include:

- **Primary domain certificate creation** with optional Subject Alternative Names
- **Configurable validation methods** (DNS recommended for automation)
- **Multiple key algorithms** including RSA_2048, RSA_4096, EC_prime256v1, and EC_secp384r1
- **Certificate Transparency (CT) logging control** - enabled by default for compliance and security monitoring

### Certificate Transparency

Certificate Transparency is a security mechanism that logs all SSL/TLS certificates to public, append-only logs. This module enables CT logging by default (`certificate_transparency_logging_preference = "ENABLED"`).

**When to enable CT logging:**

- Required for publicly-trusted certificates (browsers may reject certificates without CT)
- Enhances security by allowing monitoring for unauthorized certificate issuance
- Meets compliance requirements for many security frameworks
- Recommended for production environments

**When you might disable CT logging:**

- Internal/private certificates that don't need public visibility
- Testing environments where certificate visibility is not desired
- Specific compliance requirements that prohibit public certificate disclosure

The module creates certificates that are automatically validated and can be used with AWS services like Application Load Balancers, CloudFront distributions, and API Gateway.

## Usage

A sample variable file `example.tfvars` is available in the root directory which can be used to test this module. User needs to follow the below steps to execute this module

1. Update the `example.tfvars` to manually enter values for all fields marked within `<>` to make the variable file usable
2. Create a file `provider.tf` with the below contents

   ```angular2html
    provider "aws" {
      profile = "<profile_name>"
      region  = "<region_name>"
    }
    ```

   If using `SSO`, make sure you are logged in `aws sso login --profile <profile_name>`
3. Make sure terraform binary is installed on your local. Use command `type terraform` to find the installation location. If you are using `asdf`, you can run `asfd install` and it will install the correct terraform version for you. `.tool-version` contains all the dependencies.
4. Run the `terraform` to provision infrastructure on AWS

    ```angular2html
    # Initialize
    terraform init
    # Plan
    terraform plan -var-file example.tfvars
    # Apply (this is create the actual infrastructure)
    terraform apply -var-file example.tfvars -auto-approve
    ```

## Pre-Commit hooks

[.pre-commit-config.yaml](.pre-commit-config.yaml) file defines certain `pre-commit` hooks that are relevant to terraform, golang and common linting tasks. There are no custom hooks added.

`commitlint` hook enforces commit message in certain format. The commit contains the following structural elements, to communicate intent to the consumers of your commit messages:

- **fix**: a commit of the type `fix` patches a bug in your codebase (this correlates with PATCH in Semantic Versioning).
- **feat**: a commit of the type `feat` introduces a new feature to the codebase (this correlates with MINOR in Semantic Versioning).
- **BREAKING CHANGE**: a commit that has a footer `BREAKING CHANGE:`, or appends a `!` after the type/scope, introduces a breaking API change (correlating with MAJOR in Semantic Versioning). A BREAKING CHANGE can be part of commits of any type.
footers other than BREAKING CHANGE: <description> may be provided and follow a convention similar to git trailer format.
- **build**: a commit of the type `build` adds changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm)
- **chore**: a commit of the type `chore` adds changes that don't modify src or test files
- **ci**: a commit of the type `ci` adds changes to our CI configuration files and scripts (example scopes: Travis, Circle, BrowserStack, SauceLabs)
- **docs**: a commit of the type `docs` adds documentation only changes
- **perf**: a commit of the type `perf` adds code change that improves performance
- **refactor**: a commit of the type `refactor` adds code change that neither fixes a bug nor adds a feature
- **revert**: a commit of the type `revert` reverts a previous commit
- **style**: a commit of the type `style` adds code changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
- **test**: a commit of the type `test` adds missing tests or correcting existing tests

Base configuration used for this project is [commitlint-config-conventional (based on the Angular convention)](https://github.com/conventional-changelog/commitlint/tree/master/@commitlint/config-conventional#type-enum)

If you are a developer using vscode, [this](https://marketplace.visualstudio.com/items?itemName=joshbolduc.commitlint) plugin may be helpful.

`detect-secrets-hook` prevents new secrets from being introduced into the baseline. TODO: INSERT DOC LINK ABOUT HOOKS

In order for `pre-commit` hooks to work properly

- You need to have the pre-commit package manager installed. [Here](https://pre-commit.com/#install) are the installation instructions.
- `pre-commit` would install all the hooks when commit message is added by default except for `commitlint` hook. `commitlint` hook would need to be installed manually using the command below

```shell
pre-commit install --hook-type commit-msg
```

## To test the resource group module locally

1. For development/enhancements to this module locally, you'll need to install all of its components. This is controlled by the `configure` target in the project's [`Makefile`](./Makefile). Before you can run `configure`, familiarize yourself with the variables in the `Makefile` and ensure they're pointing to the right places.

```shell
make configure
```

This adds in several files and directories that are ignored by `git`. They expose many new Make targets.

2. The first target you care about is `env`. This is the common interface for setting up environment variables. The values of the environment variables will be used to authenticate with cloud provider from local development workstation.

`make configure` command will bring down `aws_env.sh` file on local workstation. Developer would need to modify this file, replace the environment variable values with relevant values.

These environment variables are used by `terratest` integration suit.

Then run this make target to set the environment variables on developer workstation.

```shell
make env
```

3. The first target you care about is `check`.

**Pre-requisites**
Before running this target it is important to ensure that, developer has created files mentioned below on local workstation under root directory of git repository that contains code for primitives/segments. Note that these files are `AWS` specific. If primitive/segment under development uses any other cloud provider than AWS, this section may not be relevant.

- A file named `provider.tf` with contents below

```hcl
provider "aws" {
  profile = "<profile_name>"
  region  = "<region_name>"
}
```

- A file named `terraform.tfvars` which contains key value pairs of variables used.

Note that since these files are added in `gitignore` they would not be checked in into primitive/segment's git repo.

After creating these files, for running tests associated with the primitive/segment, run

```shell
make check
```

If `make check` target is successful, developer is good to commit the code to primitive/segment's git repo.

`make check` target

- runs `terraform commands` to `lint`,`validate` and `plan` terraform code.
- runs `conftests`. `conftests` make sure `policy` checks are successful.
- runs `terratest`. This is integration test suit.
- runs `opa` tests
-
<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.5 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 5.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | 5.100.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_acm_certificate.cert](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/acm_certificate) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_domain_name"></a> [domain\_name](#input\_domain\_name) | The primary FQDN for the certificate. | `string` | n/a | yes |
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
| <a name="output_domain_validation_options"></a> [domain\_validation\_options](#output\_domain\_validation\_options) | The domain validation options for the ACM certificate. |
<!-- END_TF_DOCS -->
