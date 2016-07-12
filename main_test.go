package awsconfig

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
)

func TestCreateNewConfig(t *testing.T) {
	cases := []struct {
		log          bool
		region, role string
	}{
		{false, "", ""},
		{true, "", ""},
		{false, "us-east-1", ""},
		{true, "us-east-1", "bla-bla"},
	}

	for _, v := range cases {
		*logging = v.log
		*region = v.region
		*role = v.role
		conf := New()
		if *conf.Region != v.region {
			t.Errorf("Region actual: %s, expected: %s", *conf.Region, v.region)
		}
		if v.log && conf.LogLevel.Value() != aws.LogDebugWithHTTPBody {
			t.Error("Log level doesn't match to expected")
		}
		if !v.log && conf.LogLevel.Value() != aws.LogOff {
			t.Error("Log level doesn't match to expected")
		}
		creds, _ := conf.Credentials.Get()
		if v.role == "" && creds.ProviderName != "EnvProvider" {
			t.Errorf("Credentials provider actual: %s, expected: %s", creds.ProviderName, "EnvProvider")
		}
		if v.role != "" && creds.ProviderName != "StaticProvider" {
			t.Errorf("Credentials provider actual: %s, expected: %s", creds.ProviderName, "StaticProvider")
		}
	}
}
