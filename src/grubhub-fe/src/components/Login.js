import React, { Component } from "react";
import { Button, Form, FormGroup, Input, Label } from "reactstrap";
import "../App.css";

class Login extends Component {
  state = {};
  render() {
    return (
     
      <Form className="login-page">
        <h1 className="text-center" style={{ color: "red" }}>
          <span className="font-weight-bold">GrubHub</span>
        </h1>
        <h2 className="text-center">Login</h2>
        <FormGroup>
          <Label>Email</Label>
          <Input type="email" placeholder="Email"></Input>
        </FormGroup>
        <FormGroup>
          <Label>Password</Label>
          <Input type="password" placeholder="Password"></Input>
        </FormGroup>
        <Button className="btn-lg btn-dark btn-block">Login</Button>
        <div className="text-center pt-3">New to GrubHub?</div>
        <div className="text-center">
          <a href="/signup">Signup</a>
        </div>
      </Form>
    );
  }
}

export default Login;
