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

https://github.com/nguyensjsu/fa19-281-team-rocket/tree/master/src/payments/src/payments

# Result

- Topic creation
![Image](https://github.com/nguyensjsu/fa19-281-team-rocket/blob/master/extra_credit/SNS/screenshots/payments_topic.png)

- Subscription Confirmation email
![Image](https://github.com/nguyensjsu/fa19-281-team-rocket/blob/master/extra_credit/SNS/screenshots/subscribe_confirmation.png)

- Payment successful notification 
![Image](https://github.com/nguyensjsu/fa19-281-team-rocket/blob/master/extra_credit/SNS/screenshots/payment_success.png)

# Resources

https://aws.amazon.com/sns/?whats-new-cards.sort-by=item.additionalFields.postDateTime&whats-new-cards.sort-order=desc

https://docs.aws.amazon.com/sdk-for-go/api/service/sns/#SNS.CreateTopic

https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/using-sns-with-go-sdk.html

