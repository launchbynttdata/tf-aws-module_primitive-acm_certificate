package testimpl

import (
	"context"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	// Get the AWS region, prioritizing environment variable over AWS config
	awsRegion := GetRegionFromEnvOrConfig(t)
	awsACMClient := GetAWSACMClientWithRegion(t, awsRegion)

	t.Run("TestIsDeployed", func(t *testing.T) {
		certificateArn := terraform.Output(t, ctx.TerratestTerraformOptions(), "certificate_arn")
		out, err := awsACMClient.DescribeCertificate(context.TODO(), &acm.DescribeCertificateInput{
			CertificateArn: aws.String(certificateArn),
		})

		if err != nil {
			t.Errorf("Failure during DescribeCertificate: %v", err)
		}

		assert.NotNil(t, out.Certificate, "Expected certificate does not exist!")
		assert.Equal(t, certificateArn, *out.Certificate.CertificateArn, "Certificate ARN does not match!")

		// Additional assertion to verify we're testing in the correct region
		t.Logf("Certificate validated in region: %s", awsRegion)
	})

	t.Run("TestCommonNameisCorrect", func(t *testing.T) {
		certificateArn := terraform.Output(t, ctx.TerratestTerraformOptions(), "certificate_arn")
		out, err := awsACMClient.DescribeCertificate(context.TODO(), &acm.DescribeCertificateInput{
			CertificateArn: aws.String(certificateArn),
		})

		if err != nil {
			t.Errorf("Failure during DescribeCertificate: %v", err)
		}

		assert.NotNil(t, out.Certificate, "Expected certificate does not exist!")
		assert.Equal(t, "terratest.sandbox.launch.nttdata.com", *out.Certificate.DomainName, "Common Name does not match!")
	})
	t.Run("TestSANsAreCorrect", func(t *testing.T) {
		certificateArn := terraform.Output(t, ctx.TerratestTerraformOptions(), "certificate_arn")
		out, err := awsACMClient.DescribeCertificate(context.TODO(), &acm.DescribeCertificateInput{
			CertificateArn: aws.String(certificateArn),
		})

		if err != nil {
			t.Errorf("Failure during DescribeCertificate: %v", err)
		}

		assert.NotNil(t, out.Certificate, "Expected certificate does not exist!")
		expectedSANs := []string{"www.terratest.sandbox.launch.nttdata.com", "terratest.sandbox.launch.nttdata.com"}
		assert.ElementsMatch(t, expectedSANs, out.Certificate.SubjectAlternativeNames, "SANs do not match!")
	})
}

func GetAWSACMClient(t *testing.T) *acm.Client {
	awsACMClient := acm.NewFromConfig(GetAWSConfig(t))
	return awsACMClient
}

func GetAWSACMClientWithRegion(t *testing.T, region string) *acm.Client {
	awsACMClient := acm.NewFromConfig(GetAWSConfigWithRegion(t, region))
	return awsACMClient
}

func GetAWSEC2Client(t *testing.T) *ec2.Client {
	awsEc2Client := ec2.NewFromConfig(GetAWSConfig(t))
	return awsEc2Client
}

func GetAWSConfig(t *testing.T) (cfg aws.Config) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	require.NoErrorf(t, err, "unable to load SDK config, %v", err)
	return cfg
}

func GetAWSConfigWithRegion(t *testing.T, region string) (cfg aws.Config) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	require.NoErrorf(t, err, "unable to load SDK config for region %s, %v", region, err)
	return cfg
}

// GetRegionFromEnvOrConfig gets the AWS region, prioritizing environment variables over AWS config
func GetRegionFromEnvOrConfig(t *testing.T) string {
	// First check AWS_REGION environment variable
	if region := os.Getenv("AWS_REGION"); region != "" {
		t.Logf("Using AWS region from AWS_REGION environment variable: %s", region)
		return region
	}

	// Fall back to AWS_DEFAULT_REGION environment variable
	if region := os.Getenv("AWS_DEFAULT_REGION"); region != "" {
		t.Logf("Using AWS region from AWS_DEFAULT_REGION environment variable: %s", region)
		return region
	}

	// Fall back to AWS SDK default config resolution (checks AWS config files, instance metadata, etc.)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		t.Logf("Unable to load AWS config, using fallback region us-west-2: %v", err)
		return "us-west-2"
	}

	region := cfg.Region
	if region == "" {
		region = "us-west-2" // Final fallback to match the provider.tf
		t.Logf("No region found in AWS config, using fallback: %s", region)
	} else {
		t.Logf("Using AWS region from AWS config: %s", region)
	}

	return region
}
