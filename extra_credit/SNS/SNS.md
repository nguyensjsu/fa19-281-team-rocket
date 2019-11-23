# Goal
The goal is to send Email notification to user after he successfully place an order.

# Implementation
- To acheive this, we have used AWS Amazon Simple Notification Service (SNS), a highly available pub/sub service to send an email notification 
after his payment is successful.
- We have integrated SNS in payments service using it's create, publish, subscribe functionalities.

# Steps
- Create 'payments' topic in AWS SNS dashboard
- Subscribe every user to payments topic after signup using Email protocol in subscribe request, which triggers confirmation email to user.
- After payment is done, publish custom message to payments topic using publishInput request with user's email, which sends email to given user.
- update Docker file/ ENV's accordingly with SNS service details to reach server.

# Code


# Result


# Resources

