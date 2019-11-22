import React, { Component } from "react";
import axios from "axios";
import { Table } from "reactstrap";
import { Redirect } from "react-router-dom";
import { Route, withRouter } from "react-router-dom";
import {
  Navbar,
  Nav,
  UncontrolledDropdown,
  DropdownToggle,
  DropdownItem,
  DropdownMenu,
  Button
} from "reactstrap";
import "../App.css";
import { ROOT_URL } from "../config/URLSettings";
import Payment from "./Payment";

class Header extends Component {
  constructor(props) {
    super(props);
    this.state = {
      quantity: 0,
      item: "",
      itemSubTotal: 0,
      price: 0,
      orderedItems: [],
      redirectToPayments: false
    };
  }

  componentDidMount() {
    // axios.get(ROOT_URL + `cartItems/sam.mam@gmail.com`).then(response => {
    //   var data = response.data;
    //   console.log("Item", data);
    //   var itemSubTotal = 0;
    //   data.map(v => {
    //     var intprice = parseInt(v.price);
    //     var intquantity = parseInt(v.quantity);
    //     console.log(intprice);
    //     console.log(intquantity);
    //     itemSubTotal += intprice * intquantity;
    //   });
    //   this.setState({ itemSubTotal });
    //   this.setState({ orderedItems: data });
    // });
  }

  getCartItems = e => {
    console.log("Calling getCart items api");
    axios.get(ROOT_URL + `cartItems/`+localStorage.getItem("emailId")).then(response => {
      var data = response.data;
      console.log("Item", data);
      var itemSubTotal = 0;
      if(data !== null )
      {
        data.map(v => {
          var intprice = parseInt(v.price);
          var intquantity = parseInt(v.quantity);
          console.log(intprice);
          console.log(intquantity);
          itemSubTotal += intprice * intquantity;
        });
       
        this.setState({ orderedItems: data });
      }
      this.setState({ itemSubTotal });
      
     
    });
  };

  handleCheckout = orderedItems => {
    console.log("orderedItems", orderedItems);
    //localStorage.setItem("emailId"
    //this.props.history.push("/payment");
    this.setState({ redirectToPayments: true });
    this.setState({ orderedItems });
    //return <Redirect to="/payment" />;
  };

  render() {
    let items = this.state.orderedItems.map(oitem => {
      return (
        <DropdownItem>
          {oitem.quantity} &nbsp;&nbsp;
          <span className="text-primary"> {oitem.item} </span> &nbsp;&nbsp; $
          {oitem.price}
        </DropdownItem>
      );
    });
    return (
      <React.Fragment>
        {this.state.redirectToPayments && (
          <Redirect
            to={{
              pathname: "/payment",
              state: {
                orderedItems: this.state.orderedItems,
                itemSubTotal: this.state.itemSubTotal
              }
            }}
          />
        )}
        <div className="header">
          <Navbar color="" light expand="md">
            <h1 style={{ color: "red" }}>
              <span className="font-weight-bold">GrubHub</span>
            </h1>
            <Nav className="ml-auto" navbar>
              <UncontrolledDropdown nav inNavbar>
                <DropdownToggle nav caret onClick={this.getCartItems}>
                  Options
                </DropdownToggle>
                <DropdownMenu right>
                  <h6 className="text-center">Your Orders</h6>
                  <DropdownItem divider />

                  <tbody>{items}</tbody>

                  <DropdownItem divider />
                  <DropdownItem>
                    Item Subtotals&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; $
                    {this.state.itemSubTotal}
                  </DropdownItem>
                  <DropdownItem divider />
                  <DropdownItem>
                    <Button
                      onClick={() =>
                        this.handleCheckout(this.state.orderedItems)
                      }
                      color="success"
                    >
                      Proceed to Checkout
                    </Button>
                  </DropdownItem>
                </DropdownMenu>
              </UncontrolledDropdown>
            </Nav>
          </Navbar>
        </div>
      </React.Fragment>
    );
  }
}

export default Header;
