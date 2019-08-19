EKS Cluster Setup
-----------------

Uses terraform module https://github.com/terraform-aws-modules/terraform-aws-eks

To setup Kubernetes Cluster using EKS

```
tfenv install 0.12.0
tfenv use 0.12.0
aws-vault exec `aws-profile-name` -- terraform init
aws-vault exec `aws-profile-name` -- terraform plan -out=eks-1.tfplan
aws-vault exec `aws-profile-name` -- terraform apply eks-1.tfplan
```


Terraform will take a while to apply so you might want to increase the aws-vault timeout `--no-session`  will create unlimited timeout for apply
