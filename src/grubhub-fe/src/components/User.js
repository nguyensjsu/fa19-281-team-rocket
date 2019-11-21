import React, { Component } from "react";
import { Button, Form, FormGroup, Input, Label, Table } from "reactstrap";
import "../App.css";

class User extends Component {
  state = {};
  render() {
    return (
      <Form className="login-page">
        <h1 className="text-center" style={{ color: "red" }}>
          <span className="font-weight-bold">R-Tea</span>
        </h1>
        <h2 className="text-center">Welcome</h2>
        <Table>
          <tbody>
            <tr>
              <th scope="row">1</th>
              <td>Mark</td>
              <td>Otto</td>
              <td>@mdo</td>
            </tr>
            <tr>
              <th scope="row">2</th>
              <td>Jacob</td>
              <td>Thornton</td>
              <td>@fat</td>
            </tr>
            <tr>
              <th scope="row">3</th>
              <td>Larry</td>
              <td>the Bird</td>
              <td>@twitter</td>
            </tr>
          </tbody>
        </Table>
      </Form>
    );
  }
}

export default User;
