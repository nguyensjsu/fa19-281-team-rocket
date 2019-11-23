# Instructions to Run 
1. Setup Mongo Cluster 
2. Setup Docker Hosts 
3. For Orders API to run 
````
  * make docker-build 
  * make docker-run 
  * make docker-push 
```` 
4. Available endpoints 
```` 
  * Create new order        :(POST: /newOrder) 
  * Get order by Id         :(GET: /order/{id}) 
  * Get orders status of Id :(GET: /orderStatus/{id}) 
  * Get all orders          :(GET: /orders) 
  * Get all orders by email :(GET: /allOrdersByEmail/{uEmail}) 
  * Get all arders by Id    :(GET: /allOrdersByStatus/{status}) 
  * Delete by order Id      :(DELETE: /deleteOrder/{id}) 
  * Update order status     :(PUT: /updateOrderStatus/{id}/{status}) 
````
