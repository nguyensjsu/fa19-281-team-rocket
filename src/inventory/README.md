# How to run ?
1. - Create mongo cluster
   - setup mongo cluster
   - Setup username password (admin/Welcome_1)
   - Insert inital documents
   - db.items.insertMany( [
{
  "InventoryId": 1,
  "Quantity": 50,
  "Price": 10,
  "Name": "Burger",
  "Image": "https://cdn.newsapi.com.au/image/v1/dfd657934bf7bf648edbdc19670bf977?width=650",
  "Description": "description of burger",
  "Category": "main-course"
},
{
  "InventoryId": 2,
  "Quantity": 50,
  "Price": 5,
  "Name": "French Fries",
  "Image": "https://www.corriecooks.com/wp-content/uploads/2018/10/Instant-Pot-French-Fries-new-500x500.jpg",
  "Description": "description of french fries",
  "Category": "sides"
},
{
  "InventoryId": 3,
  "Quantity": 50,
  "Price": 2,
  "Name": "Coke",
  "Image": "https://images2.minutemediacdn.com/image/upload/c_crop,h_843,w_1500,x_0,y_70/f_auto,q_auto,w_1100/v1555172501/shape/mentalfloss/iStock-487787108.jpg",
  "Description": "description of coke",
  "Category": "sides"
}
]);

2. Change server.go to point to mongo db host

make docker-build
make docker-run
make docker-push

pull image at docker host
 sudo docker pull jojojoseph09/go-inventory
run docker container at host
 sudo docker run --name go-inventory -td -p 3000:3000 jojojoseph09/go-inventory


3. Add api gateway in front of docker host

Sample JSON for inventory item : 
{
  "category": "main-course",
  "description": "description of burger",
  "inventoryid": 17,
  "name": "Burger",
   "image": "https://cdn.newsapi.com.au/image/v1/dfd657934bf7bf648edbdc19670bf977?width=650"
  "price": 170,
  "quantity": 40
}
