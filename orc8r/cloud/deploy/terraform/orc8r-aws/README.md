# orc8r-aws

A terraform module to create an Orchestrator deployment in AWS. This module
will create the infrastructure components (EKS, RDS, Elasticsearch, etc.)
required to run the application.

## Providers

| Name | Version |
|------|---------|
| aws | >= 2.6.0 |
| tls | ~> 2.1 |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:-----:|
| cluster\_name | Name for the orc8r EKS cluster. | `string` | `"orc8r"` | no |
| deploy\_elasticsearch | Flag to deploy AWS Elasticsearch service as the datasink for aggregated logs. | `bool` | `false` | no |
| deploy\_elasticsearch\_service\_linked\_role | Flag to deploy AWS Elasticsearch service linked role with cluster. If you've already created an ES service linked role for another cluster, you should set this to false. | `bool` | `true` | no |
| efs\_project\_name | Project name for EFS file system | `string` | `"orc8r"` | no |
| eks\_map\_roles | EKS IAM role mapping. Note that by default, the creator of the cluster will be in the system:master group. | <pre>list(<br>    object({<br>      rolearn  = string,<br>      username = string,<br>      groups   = list(string),<br>    })<br>  )</pre> | `[]` | no |
| eks\_worker\_additional\_sg\_ids | Additional security group IDs to attach to EKS worker nodes. | `list(string)` | `[]` | no |
| eks\_worker\_group\_key | If specified, the worker nodes for EKS will use this EC2 keypair. | `string` | n/a | yes |
| eks\_worker\_groups | Worker group configuration for EKS. Default value is 1 worker group consisting of 3 t3.small instances. | <pre>list(<br>    object({<br>      name                 = string,<br>      instance_type        = string,<br>      asg_desired_capacity = number,<br>      asg_min_size         = number,<br>      asg_max_size         = number,<br>      autoscaling_enabled  = bool,<br>    })<br>  )</pre> | <pre>[<br>  {<br>    "asg_desired_capacity": 3,<br>    "asg_max_size": 3,<br>    "asg_min_size": 1,<br>    "autoscaling_enabled": false,<br>    "instance_type": "t3.large",<br>    "name": "wg-1"<br>  }<br>]</pre> | no |
| elasticsearch\_az\_count | AZ count for ES. | `number` | `2` | no |
| elasticsearch\_dedicated\_master\_count | Number of dedicated ES master nodes. | `number` | n/a | yes |
| elasticsearch\_dedicated\_master\_enabled | Enable/disable dedicated master nodes for ES. | `bool` | `false` | no |
| elasticsearch\_dedicated\_master\_type | Instance type for ES dedicated master nodes. | `string` | n/a | yes |
| elasticsearch\_domain\_name | Name for the AWS Elasticsearch domain. | `string` | n/a | yes |
| elasticsearch\_domain\_tags | Extra tags for the ES domain. | `map` | `{}` | no |
| elasticsearch\_ebs\_enabled | Use EBS for ES storage. See https://aws.amazon.com/elasticsearch-service/pricing/ to check if your chosen instance types can use non-EBS storage. | `bool` | `false` | no |
| elasticsearch\_ebs\_iops | IOPS for ES EBS volumes. | `number` | n/a | yes |
| elasticsearch\_ebs\_volume\_size | Size in GB to allocate for ES EBS data volumes. | `number` | n/a | yes |
| elasticsearch\_ebs\_volume\_type | EBS volume type for ES data volumes. | `string` | n/a | yes |
| elasticsearch\_instance\_count | Number of instances to allocate for ES domain. | `number` | n/a | yes |
| elasticsearch\_instance\_type | AWS instance type for ES domain. | `string` | n/a | yes |
| elasticsearch\_version | ES version for Elasticsearch domain. | `string` | `"7.1"` | no |
| global\_tags | n/a | `map` | `{}` | no |
| nms\_db\_engine\_version | MySQL engine version for NMS DB. | `string` | `"5.7"` | no |
| nms\_db\_identifier | Identifier for the RDS instance for NMS. | `string` | `"nmsdb"` | no |
| nms\_db\_instance\_class | RDS instance type for NMS DB. | `string` | `"db.m4.large"` | no |
| nms\_db\_name | DB name for NMS RDS instance. | `string` | `"magma"` | no |
| nms\_db\_password | Password for the NMS DB. | `string` | n/a | yes |
| nms\_db\_storage\_gb | Capacity in GB to allocate for NMS RDS instance. | `number` | `16` | no |
| nms\_db\_username | Username for default DB user for NMS DB. | `string` | `"magma"` | no |
| orc8r\_db\_engine\_version | Postgres engine version for Orchestrator DB. | `string` | `"9.6.15"` | no |
| orc8r\_db\_identifier | Identifier for the RDS instance for Orchestrator. | `string` | `"orc8rdb"` | no |
| orc8r\_db\_instance\_class | RDS instance type for Orchestrator DB. | `string` | `"db.m4.large"` | no |
| orc8r\_db\_name | DB name for Orchestrator RDS instance. | `string` | `"orc8r"` | no |
| orc8r\_db\_password | Password for the Orchestrator DB. | `string` | n/a | yes |
| orc8r\_db\_storage\_gb | Capacity in GB to allocate for Orchestrator RDS instance. | `number` | `64` | no |
| orc8r\_db\_username | Username for default DB user for Orchestrator DB. | `string` | `"orc8r"` | no |
| orc8r\_domain\_name | Base domain name for AWS Route 53 hosted zone | `string` | n/a | yes |
| region | AWS region to deploy Orchestrator components into. The chosen region must provide EKS. | `string` | n/a | yes |
| secretsmanager\_orc8r\_secret | AWS Secretmanager secret to store Orchestrator secrets. | `string` | n/a | yes |
| vpc\_cidr | CIDR block for the VPC. | `string` | `"10.10.0.0/16"` | no |
| vpc\_database\_subnets | CIDR blocks for the VPC's database subnets. | `list(string)` | <pre>[<br>  "10.10.200.0/24",<br>  "10.10.201.0/24",<br>  "10.10.202.0/24"<br>]</pre> | no |
| vpc\_extra\_tags | Tags to add to the VPC. | `map` | `{}` | no |
| vpc\_name | Name for the VPC that will contain all the Orchestrator components. | `string` | `"orc8r_vpc"` | no |
| vpc\_private\_subnets | CIDR blocks for the VPC's private subnets. | `list(string)` | <pre>[<br>  "10.10.100.0/24",<br>  "10.10.101.0/24",<br>  "10.10.102.0/24"<br>]</pre> | no |
| vpc\_public\_subnets | CIDR blocks for the VPC's public subnets. | `list(string)` | <pre>[<br>  "10.10.0.0/24",<br>  "10.10.1.0/24",<br>  "10.10.2.0/24"<br>]</pre> | no |

## Outputs

| Name | Description |
|------|-------------|
| efs\_file\_system\_id | ID of the EFS file system created for k8s persistent volumes. |
| efs\_provisioner\_role\_arn | ARN of the IAM role for the EFS provisioner. |
| eks\_cluster\_id | Cluster ID for the EKS cluster |
| eks\_config\_map\_aws\_auth | A k8s configmap to allow authentication to the EKS cluster. |
| es\_endpoint | Endpoint of the ES cluster if deployed. |
| external\_dns\_role\_arn | IAM role ARN for external-dns |
| kubeconfig | kubectl config file to access the EKS cluster |
| nms\_db\_host | Hostname of the NMS RDS instance |
| nms\_db\_name | Database name for NMS RDS instance |
| nms\_db\_pass | NMS DB password |
| nms\_db\_user | Database username for NMS RDS instance |
| orc8r\_db\_host | Hostname of the Orchestrator RDS instance |
| orc8r\_db\_name | Database name for Orchestrator RDS instance |
| orc8r\_db\_pass | Orchestrator DB password |
| orc8r\_db\_port | Database connection port for Orchestrator RDS instance |
| orc8r\_db\_user | Database username for Orchestrator RDS instance |
| orc8r\_domain\_name | Base domain name for Orchestrator application components. |
| route53\_nameservers | Route53 zone nameservers for external DNS configuration. |
| route53\_zone\_id | Route53 zone ID for Orchestrator domain name |
| secretsmanager\_secret\_name | Name of the secretsmanager secret for deployment secrets |

