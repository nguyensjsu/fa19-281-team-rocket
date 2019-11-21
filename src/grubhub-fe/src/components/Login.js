import React, { Component } from "react";
import { Button, Form, FormGroup, Input, Label } from "reactstrap";
import "../App.css";
import axios from "axios";

class Login extends Component {
  constructor(props) {
    super(props);
    this.state = {
      userEmail: "",
      password: ""
    };

    this.emailHandler = this.emailHandler.bind(this);
    this.pwdHandler = this.pwdHandler.bind(this);
    this.loginHandler = this.loginHandler.bind(this);
  }
  emailHandler = e => {
    this.setState({
      userEmail: e.target.value
    });
  };

  pwdHandler = e => {
    this.setState({
      password: e.target.value
    });
  };

  loginHandler = e => {
    const data = {
      email: this.state.userEmail,
      password: this.state.password
    };

    axios
      .post(
        "https://1px6zgas05.execute-api.us-west-2.amazonaws.com/prod/login",
        data
      )
      .then(response => {
        console.log("Status Code : ", response.status);
        if (response.status === 200) {
          localStorage.setItem("emailId", this.state.userEmail);
          this.props.history.push("/user");
        } else {
          window.alert("Invalid Login!!");
        }
      });
  };

  render() {
    return (
      <Form className="login-page">
        <h1 className="text-center" style={{ color: "red" }}>
          <span className="font-weight-bold">R-Tea</span>
        </h1>
        <h2 className="text-center">Login</h2>
        <FormGroup>
          <Label>Email</Label>
          <Input
            type="email"
            placeholder="Email"
            onChange={this.emailHandler}
          ></Input>
        </FormGroup>
        <FormGroup>
          <Label>Password</Label>
          <Input
            type="password"
            placeholder="Password"
            onChange={this.pwdHandler}
          ></Input>
        </FormGroup>
        <Button
          className="btn-lg btn-dark btn-block"
          onClick={this.loginHandler}
        >
          Login
        </Button>
        <div className="text-center pt-3">New to Rocket Restaurent?</div>
        <div className="text-center">
          <a href="/signup">Signup</a>
        </div>
      </Form>
    );
  }
}

export default Login;
