package store

import (
	"fmt"
	"time"
)

type MockStore struct{}

func (b *MockStore) GetLatestVersion(namespace string, name string, provider string) (string, error) {
	return "0.0.1", nil
}

func (b *MockStore) GetDownloadURL(namespace string, name string, provider string, version string) (string, error) {
	return fmt.Sprintf("https://localhost/downloads/%s", name), nil
}

func (b *MockStore) GetModulesList(namespace string, name string, provider string) ([]Module, error) {
	return b.GetModuleVersions(namespace, name, provider)
}

func (b *MockStore) GetModuleVersions(namespace string, name string, provider string) ([]Module, error) {
	var modules []Module
	t, _ := time.Parse(time.RFC3339Nano, "2013-06-05T14:10:43.678Z")
	m := &Module{
		Id:           "hashicorp/consul/aws/0.0.1",
		Owner:        "gruntwork-team",
		Namespace:    "hashicorp",
		Name:         "consul",
		Version:      "0.0.1",
		Provider:     "aws",
		Description:  "A Terraform Module for how to run Consul on AWS using Terraform and Packer",
		Source:       "https://github.com/hashicorp/terraform-aws-consul",
		Published_at: t,
		Downloads:    5,
		Verified:     false,
	}
	modules = append(modules, *m)
	m = &Module{
		Id:           "hashicorp/consul/aws/0.0.2",
		Owner:        "gruntwork-team",
		Namespace:    "hashicorp",
		Name:         "consul",
		Version:      "0.0.2",
		Provider:     "aws",
		Description:  "A Terraform Module for how to run Consul on AWS using Terraform and Packer",
		Source:       "https://github.com/hashicorp/terraform-aws-consul",
		Published_at: t,
		Downloads:    5,
		Verified:     false,
	}
	modules = append(modules, *m)
	m = &Module{
		Id:           "hashicorp/consul/aws/0.0.2",
		Owner:        "gruntwork-team",
		Namespace:    "hashicorp",
		Name:         "consul",
		Version:      "0.0.2",
		Provider:     "foobar",
		Description:  "A Terraform Module for how to run Consul on Foobar using Terraform and Packer",
		Source:       "https://github.com/hashicorp/terraform-aws-consul",
		Published_at: t,
		Downloads:    5,
		Verified:     false,
	}
	modules = append(modules, *m)
	return modules, nil
}

func (b *MockStore) GetModuleDetails(namespace string, name string, provider string, version string) (Module, error) {
	t, _ := time.Parse(time.RFC3339Nano, "2013-06-05T14:10:43.678Z")
	module := &Module{
		Id:           "hashicorp/consul/aws/0.0.1",
		Owner:        "gruntwork-team",
		Namespace:    "hashicorp",
		Name:         "consul",
		Version:      "0.0.1",
		Provider:     "aws",
		Description:  "A Terraform Module for how to run Consul on AWS using Terraform and Packer",
		Source:       "https://github.com/hashicorp/terraform-aws-consul",
		Published_at: t,
		Downloads:    5,
		Verified:     false,
	}

	return *module, nil
}

func (b *MockStore) ValidateAccessToken(token string) (string, error) {
	if token == "reallylongstringthatstotallygoingtostandoutinthelistofheaders" {
		return "mockuser", nil
	} else {
		return "", fmt.Errorf("Invalid token %s", token)
	}
}

func NewMockStore() *MockStore {
	return &MockStore{}
}
