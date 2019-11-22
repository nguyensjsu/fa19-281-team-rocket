import React, { Component } from "react";
import axios from "axios";
import { Table } from "reactstrap";
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

class Header extends Component {
  constructor(props) {
    super(props);
    this.state = {
      quantity: 0,
      item: "",
      itemSubTotal: 0,
      price: 0,
      orderedItems: []
    };
  }

  componentDidMount() {
    axios.get(ROOT_URL + `cartItems/sam.mam@gmail.com`).then(response => {
      var data = response.data;
      console.log("Item", data[0].item);
      // this.setState({quantity:data[0].quantity})
      // this.setState({item:data[0].item})
      // this.setState({price :data[0].price })
      // var totalSubtotal =  data[0].price * data[0].quantity
      // this.setState({itemSubTotal:totalSubtotal})
      var itemSubTotal = 0;
      data.map(v => {
        var intprice = parseInt(v.price);
        var intquantity = parseInt(v.quantity);
        console.log(intprice);
        console.log(intquantity);
        itemSubTotal += intprice * intquantity;
      });
      this.setState({ itemSubTotal });
      this.setState({ orderedItems: data });
    });
  }

  getCartItems = e => {
    console.log("Calling getCart items api");
    axios.get(ROOT_URL + `cartItems/sam.mam@gmail.com`).then(response => {
      var data = response.data;
      console.log("Item", data[0].item);
      // this.setState({quantity:data[0].quantity})
      // this.setState({item:data[0].item})
      // this.setState({price :data[0].price })
      // var totalSubtotal =  data[0].price * data[0].quantity
      // this.setState({itemSubTotal:totalSubtotal})
      var itemSubTotal = 0;
      data.map(v => {
        var intprice = parseInt(v.price);
        var intquantity = parseInt(v.quantity);
        console.log(intprice);
        console.log(intquantity);
        itemSubTotal += intprice * intquantity;
      });
      this.setState({ itemSubTotal });
      this.setState({ orderedItems: data });
    });
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
                  <Button color="success">Proceed to Checkout</Button>
                </DropdownItem>
              </DropdownMenu>
            </UncontrolledDropdown>
          </Nav>
        </Navbar>
      </div>
    );
  }
}

export default Header;
