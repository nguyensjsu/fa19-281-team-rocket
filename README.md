# Team Hackathon Project
## Topic of Project
SpartanHub - SAAS Application for ordering food from a restaurant. Deployed on Amazon Cloud and Heroku.
* Technologies used : ReactJS, Go, MongoDB, Docker, Kubernetes, Kong, AWS

## Team Members
 * [Deepika Yannamani](https://github.com/deepikay912)
 * [Harshraj Mahesh](https://github.com/harshrajm)
 * [Jojo Joseph](https://github.com/jojojoseph94)
 * [Megha Lnu](https://github.com/megha-31)
 * [Vaishnavi Ramesh](https://github.com/iivaishnavii)

## Contributions
Microservice | Owner
-------------|------
Login | Megha
Orders | Harshraj
Payments | Deepika
Cart | Vaishnavi
Inventory | Jojo

Other services - FrontEnd React - All members contributed to their corresponding frontend pages and other styling.

## Summary of Team Project
Implemented online store for a restaurant using 5 microservices deployed in AWS and frontend deployed in Heroku. The user will be able to signup, login, browse inventory, add items to cart, perform payment operations and check order history.

## High Level Architecture Diagram
![Architecture Diagram](https://github.com/nguyensjsu/fa19-281-team-rocket/blob/master/Architecture%20diagram.jpg)

## Summary of Key Features
 * All microservices are deployed in private subnets and exposed via API Gateways/ Route 53. Each microservice has it's own Mongo cluster for storage of data.
 * For orders microservice, data sharding is performed for orders collection across 2 replica sets. This enables better scaling (horizontal) as application will be able to handle large amounts of orders without hitting DB bottlenecks.
 * The images to be served as part of the website are stored in an S3 bucket and served via CloudFront CDN endpoint.
 * Continuous deployment was setup for frontend application deployed on Heroku. This was achived using Github Actions. For every push that occurs to master repository, the React application subtree is pushed to Heroku master repository using Heroku API key. This ensures that any UI changes pushed to master is automatically deployed.

## Application Flow Gif
![Gif](https://github.com/nguyensjsu/fa19-281-team-rocket/blob/master/journal/SpartanEats.gif)

## Slides
![Slides](https://github.com/nguyensjsu/fa19-281-team-rocket/blob/master/SpartanHub.pptx)

