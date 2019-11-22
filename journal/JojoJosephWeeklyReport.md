# Team Hackathon Project

Jojo Joseph

WEEK 1 (11/09/19 - 11/15/19)

11/10/19: Brainstorming and exploring project ideas
11/11/19: Discussed everyone's ideas and finalised project idea similar to grubhub, divided services among team
11/13/19: Reading up on golang, created api spec in swaggerhub
11/15/19: Created inventory api with CRUD operations

Planned Next Week
Task1 : integrate with other services
Task2 : figure out how to update count without concurrency issues

Problems
Update inventory works only if you send full json body.
This is fine but when you need to update inventory as part of a purchase, this endpoint cannot be used.
We need a new endpoint which can update count concurrently without having issues.

Time to Spend
Task1 : 5 hours
Task2 : ??

WEEK 2 (11/16/19 - 11/23/19)

11/16/19: Deployed mongo cluster, made docker images for the inventory microservice, deployed and tested the api
11/22/19: Created AMI for launch configuration, made auto scaling group and attached it to the target group so that
          Application scales based on load.
          Created mongo shards for orders API in Harsh's AWS account. Sharding appears to be working fine for orders.
          Created Architecture diagram for the system.
