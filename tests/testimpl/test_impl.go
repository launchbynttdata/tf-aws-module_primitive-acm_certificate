package testimpl

import (
	"context"
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

const (
	describeCertificateError = "Failure during DescribeCertificate: %v"
	certificateNotExist      = "Expected certificate does not exist!"
)

func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	// Get AWS config - this is our single source of truth for region and configuration
	awsConfig := GetAWSConfig(t)
	awsACMClient := acm.NewFromConfig(awsConfig)
	awsRegion := awsConfig.Region

	t.Run("TestIsDeployed", func(t *testing.T) {
		certificateArn := terraform.Output(t, ctx.TerratestTerraformOptions(), "certificate_arn")
		t.Logf("Looking for certificate ARN: %s in region: %s", certificateArn, awsRegion)

		out, err := awsACMClient.DescribeCertificate(context.TODO(), &acm.DescribeCertificateInput{
			CertificateArn: aws.String(certificateArn),
		})

		if err != nil {
			t.Fatalf(describeCertificateError, err)
		}

		require.NotNil(t, out.Certificate, certificateNotExist)
		assert.Equal(t, certificateArn, *out.Certificate.CertificateArn, "Certificate ARN does not match!")

		t.Logf("Certificate validated in region: %s", awsRegion)
	})

	t.Run("TestCommonNameisCorrect", func(t *testing.T) {
		certificateArn := terraform.Output(t, ctx.TerratestTerraformOptions(), "certificate_arn")
		out, err := awsACMClient.DescribeCertificate(context.TODO(), &acm.DescribeCertificateInput{
			CertificateArn: aws.String(certificateArn),
		})

		if err != nil {
			t.Fatalf(describeCertificateError, err)
		}

		require.NotNil(t, out.Certificate, certificateNotExist)
		assert.Equal(t, "terratest.sandbox.launch.nttdata.com", *out.Certificate.DomainName, "Common Name does not match!")
	})

	t.Run("TestSANsAreCorrect", func(t *testing.T) {
		certificateArn := terraform.Output(t, ctx.TerratestTerraformOptions(), "certificate_arn")
		out, err := awsACMClient.DescribeCertificate(context.TODO(), &acm.DescribeCertificateInput{
			CertificateArn: aws.String(certificateArn),
		})

		if err != nil {
			t.Fatalf(describeCertificateError, err)
		}

		require.NotNil(t, out.Certificate, certificateNotExist)
		expectedSANs := []string{"www.terratest.sandbox.launch.nttdata.com", "terratest.sandbox.launch.nttdata.com"}
		assert.ElementsMatch(t, expectedSANs, out.Certificate.SubjectAlternativeNames, "SANs do not match!")
	})
}

func GetAWSACMClient(t *testing.T) *acm.Client {
	awsACMClient := acm.NewFromConfig(GetAWSConfig(t))
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
