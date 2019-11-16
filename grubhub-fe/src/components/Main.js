import React, { Component } from "react";
import { Route } from "react-router-dom";
import Login from "./Login";
import Signup from "./Signup";
import User from "./User";

class Main extends Component {
  state = {};
  render() {
    return (
      <div>
        <Route path="/login" component={Login} />
        <Route path="/signup" component={Signup} />
        <Route path="/user" component={User} />
      </div>
    );
  }
}

export default Main;
