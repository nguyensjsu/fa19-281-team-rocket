# Goal
To run cart service in an EKS cluster

# Steps
## Create IAM roles and policies:
 - Cluster Admin
 ![eks1](https://github.com/nguyensjsu/fa19-281-team-rocket/blob/master/extra_credit/EKS/images/cluster%20admin.png)
 - Cluster user
 ![eks2](https://github.com/nguyensjsu/fa19-281-team-rocket/blob/master/extra_credit/EKS/images/cluster%20user.png)
## Create VPC and subnets
Created VPC,subnets and security groups using CloudFormation stack 
https://amazon-eks.s3-us-west-2.amazonaws.com/cloudformation/2019-11-15/amazon-eks-vpc-sample.yaml
![eks3](https://github.com/nguyensjsu/fa19-281-team-rocket/blob/master/extra_credit/EKS/images/vpc.png)
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

![eks4](https://github.com/nguyensjsu/fa19-281-team-rocket/blob/master/extra_credit/EKS/images/node%20groups.png)

## Map Master to Worker Nodes
From the console,add the nodes to cluster. Provide the nodeInstance role and mention the number of nodes needed. I have provided 3 nodes. An autoscaling group is created by EKS for the worker nodes.

![eks5](https://github.com/nguyensjsu/fa19-281-team-rocket/blob/master/extra_credit/EKS/images/mmwn.png)
![eks6](https://github.com/nguyensjsu/fa19-281-team-rocket/blob/master/extra_credit/EKS/images/mmwn2.png)

## Create a deployment to run in the pods
Apply the deployment to EKS using this command :   
````
Kubectl apply -f deployment.yaml
````

![eks7](https://github.com/nguyensjsu/fa19-281-team-rocket/blob/master/extra_credit/EKS/images/ss1.png)

## Expose the service using Load Balancer

````
kubectl expose deployment cart-deployment --type=LoadBalancer --port=8000 --target-port=8000 --name=cart-load-balancer
````

![eks8](https://github.com/nguyensjsu/fa19-281-team-rocket/blob/master/extra_credit/EKS/images/ss2.png)

Test the endpoint using curl

![eks9](https://github.com/nguyensjsu/fa19-281-team-rocket/blob/master/extra_credit/EKS/images/ss3.png)
