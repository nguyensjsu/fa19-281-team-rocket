# How to run ?
  *  1. Create mongo cluster
     2. setup mongo cluster
     3. Setup username password (admin/Welcome_1)
     4. Insert inital documents
     ```
     db.items.insertMany( [ {
     "InventoryId": 1, "Quantity": 50, "Price": 10, "Name": "Burger",
     "Image": "https://dr8kj9p1zyz6i.cloudfront.net/burger.jpg",
     "Description": "description of burger", "Category": "main-course" },
     { "InventoryId": 2, "Quantity": 50, "Price": 5, "Name": "French Fries",
     "Image": "https://dr8kj9p1zyz6i.cloudfront.net/frenchfries.jpg",
     "Description": "description of french fries", "Category": "sides" },
     { "InventoryId": 3, "Quantity": 50, "Price": 2, "Name": "Coke",
     "Image": "https://dr8kj9p1zyz6i.cloudfront.net/cola.jpg",
     "Description": "description of coke", "Category": "sides" } ]);
     ```

* Change server.go to point to mongo db host
  ```
  make docker-build
  make docker-run
  make docker-push
  ```

  pull image and run at docker host
  ```
  sudo docker pull jojojoseph09/go-inventory
  sudo docker run --name go-inventory -td -p 3000:3000 jojojoseph09/go-inventory
  ```


* Add api gateway in front of docker host

* Make sure files are uploaded to s3 bucket.


```Sample JSON for inventory item : 
{
  "category": "main-course",
  "description": "description of burger",
  "inventoryid": 17,
  "name": "Burger",
   "image": "https://dr8kj9p1zyz6i.cloudfront.net/burger.jpg",
  "price": 170,
  "quantity": 40
}```
