# Goal
To run cart service in an EKS cluster

# Steps
## Create IAM roles and policies:
 - Cluster Admin
 ![eks1](https://github.com/nguyensjsu/fa19-281-team-rocket/blob/master/extra_credit/cluster%20admin.png)
 - Cluster user
 ![eks2](https://github.com/nguyensjsu/fa19-281-team-rocket/blob/master/extra_credit/cluster%20user.png)
## Create VPC and subnets
Created VPC,subnets and security groups using CloudFormation stack 
https://amazon-eks.s3-us-west-2.amazonaws.com/cloudformation/2019-11-15/amazon-eks-vpc-sample.yaml
![eks3]()
## Create cluster
From the console,create a cluster and provide the VPC,subnets and security groups created by CloudFormation stack.
## Create Kubeconfig file locally
To set the context locally, execute the command
````
aws eks --region us-west-2 update-kubeconfig --name rocketCluster
````
## Create Node Groups
Using the cloud formation stack, create node groups
https://amazon-eks.s3-us-west-2.amazonaws.com/cloudformation/2019-11-15/amazon-eks-nodegroup-role.yaml
Save the Node Instance Role created by this stack.



