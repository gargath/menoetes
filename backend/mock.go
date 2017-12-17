package backend

import (
  "fmt"
)

type MockBackend struct {}

func (b *MockBackend) GetLatestVersion(namespace string, name string, provider string) (string, error) {
  return "0.0.1", nil
}

func (b *MockBackend) GetDownloadURL(namespace string, name string, provider string, version string) (string, error) {
  return fmt.Sprintf("https://localhost/downloads/%s", name), nil
}

func (b *MockBackend) GetModuleVersions(namespace string, name string, provider string) (string, error) {
  d := `
  {
     "modules": [
        {
           "source": "hashicorp/consul/aws",
           "versions": [
              {
                 "version": "0.0.1",
                 "submodules" : [
                    {
                       "path": "modules/consul-cluster",
                       "providers": [
                          {
                             "name": "aws",
                             "version": ""
                          }
                       ],
                       "dependencies": []
                    },
                    {
                       "path": "modules/consul-security-group-rules",
                       "providers": [
                          {
                             "name": "aws",
                             "version": ""
                          }
                       ],
                       "dependencies": []
                    },
                    {
                       "providers": [
                          {
                             "name": "aws",
                             "version": ""
                          }
                       ],
                       "dependencies": [],
                       "path": "modules/consul-iam-policies"
                    }
                 ],
                 "root": {
                    "dependencies": [],
                    "providers": [
                       {
                          "name": "template",
                          "version": ""
                       },
                       {
                          "name": "aws",
                          "version": ""
                       }
                    ]
                 }
              }
           ]
        }
     ]
  }
  `
  return d, nil
}

func (b *MockBackend) GetModuleDetails(namespace string, name string, provider string, version string) (string, error) {
  d := `
  {
    "id": "hashicorp/consul/aws/0.0.1",
    "owner": "gruntwork-team",
    "namespace": "hashicorp",
    "name": "consul",
    "version": "0.0.1",
    "provider": "aws",
    "description": "A Terraform Module for how to run Consul on AWS using Terraform and Packer",
    "source": "https://github.com/hashicorp/terraform-aws-consul",
    "published_at": "2017-09-14T23:22:44.793647Z",
    "downloads": 113,
    "verified": false,
    "root": {
      "path": "",
      "readme": "# Consul AWS Module\n\nThis repo contains a Module for how to deploy a [Consul]...",
      "empty": false,
      "inputs": [
        {
          "name": "ami_id",
          "description": "The ID of the AMI to run in the cluster. ...",
          "default": "\"\""
        },
        {
          "name": "aws_region",
          "description": "The AWS region to deploy into (e.g. us-east-1).",
          "default": "\"us-east-1\""
        }
      ],
      "outputs": [
        {
          "name": "num_servers",
          "description": ""
        },
        {
          "name": "asg_name_servers",
          "description": ""
        }
      ],
      "dependencies": [],
      "resources": []
    },
    "submodules": [
      {
        "path": "modules/consul-cluster",
        "readme": "# Consul Cluster\n\nThis folder contains a [Terraform](https://www.terraform.io/) ...",
        "empty": false,
        "inputs": [
          {
            "name": "cluster_name",
            "description": "The name of the Consul cluster (e.g. consul-stage). This variable is used to namespace all resources created by this module.",
            "default": ""
          },
          {
            "name": "ami_id",
            "description": "The ID of the AMI to run in this cluster. Should be an AMI that had Consul installed and configured by the install-consul module.",
            "default": ""
          }
        ],
        "outputs": [
          {
            "name": "asg_name",
            "description": ""
          },
          {
            "name": "cluster_size",
            "description": ""
          }
        ],
        "dependencies": [],
        "resources": [
          {
            "name": "autoscaling_group",
            "type": "aws_autoscaling_group"
          },
          {
            "name": "launch_configuration",
            "type": "aws_launch_configuration"
          }
        ]
      }
    ],
    "providers": [
      "aws",
      "azurerm"
    ],
    "versions": [
      "0.0.1"
    ]
  }
  `
  return d, nil
}

func NewMockBackend() (*MockBackend) {
  return &MockBackend{}
}
