package backend

import (
  "fmt"
  "time"
)

type MockBackend struct {}

func (b *MockBackend) GetLatestVersion(namespace string, name string, provider string) (string, error) {
  return "0.0.1", nil
}

func (b *MockBackend) GetDownloadURL(namespace string, name string, provider string, version string) (string, error) {
  return fmt.Sprintf("https://localhost/downloads/%s", name), nil
}

func (b *MockBackend) GetModulesList(namespace string, name string, provider string) ([]Module, error) {
  return b.GetModuleVersions(namespace, name, provider)
}

func (b *MockBackend) GetModuleVersions(namespace string, name string, provider string) ([]Module, error) {
  var modules []Module
  t, _ := time.Parse(time.RFC3339Nano, "2013-06-05T14:10:43.678Z")
  m := &Module{
    Id: "hashicorp/consul/aws/0.0.1",
    Owner: "gruntwork-team",
    Namespace: "hashicorp",
    Name: "consul",
    Version: "0.0.1",
    Provider: "aws",
    Description: "A Terraform Module for how to run Consul on AWS using Terraform and Packer",
    Source: "https://github.com/hashicorp/terraform-aws-consul",
    Published_at: t,
    Downloads: 5,
    Verified: false,
  }
  modules = append(modules, *m)
  m = &Module{
    Id: "hashicorp/consul/aws/0.0.1",
    Owner: "foobar",
    Namespace: "hashicorp",
    Name: "anuga",
    Version: "0.0.1",
    Provider: "aws",
    Description: "A Terraform Module for nothingin particular",
    Source: "https://github.com/hashicorp/foobar",
    Published_at: t,
    Downloads: 15,
    Verified: false,
  }
  modules = append(modules, *m)
  return modules, nil
}

func (b *MockBackend) GetModuleDetails(namespace string, name string, provider string, version string) (Module, error) {
  t, _ := time.Parse(time.RFC3339Nano, "2013-06-05T14:10:43.678Z")
  module := &Module{
    Id: "hashicorp/consul/aws/0.0.1",
    Owner: "gruntwork-team",
    Namespace: "hashicorp",
    Name: "consul",
    Version: "0.0.1",
    Provider: "aws",
    Description: "A Terraform Module for how to run Consul on AWS using Terraform and Packer",
    Source: "https://github.com/hashicorp/terraform-aws-consul",
    Published_at: t,
    Downloads: 5,
    Verified: false,
  }

  return *module, nil
}

func NewMockBackend() (*MockBackend) {
  return &MockBackend{}
}
