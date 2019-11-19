import React, { Component } from "react";
import { Route } from "react-router-dom";
import Login from "./Login";
import Signup from "./Signup";
import User from "./User";
import Inventory from "./Item";
import Header from "./Header";
import Payment from "./Payment";

class Main extends Component {
  state = {};
  render() {
    return (
      <div>
        <Route path="/header" component={Header} />
        <Route path="/login" component={Login} />
        <Route path="/signup" component={Signup} />
        <Route path="/user" component={User} />
        
        
        <Route path="/inventory" component={Inventory} />
        <Route path="/payment" component={Payment} />
        
      </div>
    );
  }
}

export default Main;
