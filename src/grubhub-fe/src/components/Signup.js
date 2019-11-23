import React, { Component } from "react";
import { Button, Form, FormGroup, Input, Label } from "reactstrap";
import "../App.css";
import axios from "axios";

class Signup extends Component {
  constructor(props) {
    super(props);
    this.state = {
      userEmail: "",
      password: "",
      name: ""
    };

    this.emailHandler = this.emailHandler.bind(this);
    this.pwdHandler = this.pwdHandler.bind(this);
    this.nameHandler = this.nameHandler.bind(this);
    this.signupHandler = this.signupHandler.bind(this);
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

  nameHandler = e => {
    this.setState({
      name: e.target.value
    });
  };

  signupHandler = e => {
    const data = {
      email: this.state.userEmail,
      password: this.state.password,
      name: this.state.name
    };

    axios
      .post(
        "https://1px6zgas05.execute-api.us-west-2.amazonaws.com/prod/signup",
        data
      )
      .then(response => {
        console.log("Status Code : ", response.status);
        if (response.status === 200) {
          localStorage.setItem("name", this.state.name);
          this.props.history.push("/login");
          //localStorage.setItem("emailId", this.state.userEmail);
        } else {
          window.alert("Invalid params!!");
        }
      });
  };

  render() {
    return (
      <Form className="login-page">
        <h1 className="text-center" style={{ color: "red" }}>
          <span className="font-weight-bold">SpartanHub</span>
        </h1>
        <h2 className="text-center">Sign up</h2>
        <FormGroup>
          <Label>Name</Label>
          <Input placeholder="Name" onChange={this.nameHandler}></Input>
        </FormGroup>
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
          onClick={this.signupHandler}
        >
          Submit
        </Button>
        <div className="text-center pt-3">Already a user?</div>
        <div className="text-center">
          <a href="/login">Login</a>
        </div>
      </Form>
    );
  }
}

export default Signup;
