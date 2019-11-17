import React, { Component } from "react";
import axios from 'axios';
import {Navbar,Nav,UncontrolledDropdown,DropdownToggle,DropdownItem,DropdownMenu} from "reactstrap"
import "../App.css";
import { ROOT_URL } from '../config/URLSettings';


class Header extends Component{

    constructor(props){
        super(props);
        this.state ={
            quantity : 0,
            item : "",
            itemSubTotal : 0,
            price : 0
        }
    }
    
    componentDidMount()
     {
            axios.get(ROOT_URL+`cartItems/sam.mam@gmail.com`)
            .then((response)=>{
                var data = response.data
                console.log("Item",data[0].item)
                this.setState({quantity:data[0].quantity})
                this.setState({item:data[0].item})
                this.setState({price :data[0].price })
                var totalSubtotal =  data[0].price * data[0].quantity
                this.setState({itemSubTotal:totalSubtotal})

            })
    }
    render() {
        return (
                <div className="header">
                    <Navbar color="light" light expand="md">
                        <h1 style={{ color: "red" }}>
                        <span className="font-weight-bold">GrubHub</span>
                        </h1>
                        <Nav className="ml-auto" navbar>
                        <UncontrolledDropdown nav inNavbar>
                                
                                <DropdownToggle nav caret>
                                    Options
                                </DropdownToggle>
                                <DropdownMenu right>
                                    <h6 className="text-center">Your Orders</h6>
                                    <DropdownItem divider />
                                    <DropdownItem >
                                        {this.state.quantity} &nbsp;&nbsp;
                                        <span className="text-primary">  {this.state.item} </span> &nbsp;&nbsp;
                                        ${this.state.price}
                                    </DropdownItem>
                                    <DropdownItem divider />
                                    <DropdownItem >
                                        Item Subtotal &nbsp;&nbsp;&nbsp; ${this.state.itemSubTotal}
                                    </DropdownItem>
                                   
                                    <DropdownItem divider />
                                    <DropdownItem >
                                         Sign Out
                                    </DropdownItem>
                                </DropdownMenu>
                                </UncontrolledDropdown>
                        </Nav>
                    </Navbar>
                </div>
        )
    }
}

export default Header;