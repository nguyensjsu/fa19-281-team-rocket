import React, { Component } from "react";
import { Button, Form, FormGroup, Input, Label, Table } from "reactstrap";
import "../App.css";
import axios from "axios";
import Header from "./Header";

class User extends Component {
  state = {
    orders: []
  };

  componentWillMount() {
    let email = localStorage.getItem("emailId");
    axios
      .get("http://34.219.240.229:8080/allOrdersByEmail/" + email)
      .then(response => {
        console.log("Status Code : ", response.status);
        if (response.status === 200) {
          this.setState({
            orders: response.data
          });
        }
      });
  }
  render() {
    let orders = this.state.orders;
    let name = localStorage.getItem("name");
    if (orders != null) {
      return (
     
        
        <div>
          <Header />
          <Form className="login-page">
            
            <h1 className="text-center" style={{ color: "red" }}>
              <span className="font-weight-bold">R-Tea</span>
            </h1>
            <h2 className="text-center">Welcome {name}</h2>
          </Form>
          <div className="text-center">
            <Table class="table table-hover">
              <thead>
                <tr>
                  <th>#</th>
                  <th>OrderID</th>
                  <th>Status</th>
                  <th>Time</th>
                </tr>
              </thead>
              <tbody>
                {this.state.orders.map(order => (
                  <tr>
                    <th scope="row">*</th>
                    <td>{order["_id"]}</td>
                    <td>{order["orderStatus"] ? "Success" : "Failed"}</td>
                    <td>{Date(order["orderPlacedTime"].toString())}</td>
                  </tr>
                ))}
              </tbody>
            </Table>
          </div>
        </div>
      );
    } else {
      let name = localStorage.getItem("name");
      return (
        <div>
        <Header />
        <Form className="login-page">
          
          <h1 className="text-center" style={{ color: "red" }}>
            <span className="font-weight-bold">R-Tea</span>
          </h1>
          <h2 className="text-center">Welcome {name}</h2>
          <h3 className="text-center">No Orders Found!!!</h3>
        </Form>
        </div>
      );
    }
  }
}

export default User;
