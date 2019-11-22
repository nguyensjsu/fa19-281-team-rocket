import React, { Component, useState } from "react";
import { Card, CardImg, CardText, CardBody, CardTitle } from "reactstrap";
import {
  Row,
  Col,
  Container,
  Badge,
  Button,
  Modal,
  ModalHeader,
  ModalBody,
  ModalFooter
} from "reactstrap";
import "../App.css";
import axios from "axios";
import Header from "./Header";
import ModalExample from "./modal";
import Footer from "./Footer";

class Item extends Component {
  constructor(props) {
    super(props);
  }

  render() {
    console.log(this.props);

    return (
      <Col sm="3">
        <React.Fragment>
          <Card className="box">
            <CardImg
              class="card-img-top"
              top
              width="100%"
              src={this.props.item.Image}
              alt="Card image cap"
            />
            <CardBody>
              <Badge color="secondary float-right m-1" pill>
                {this.props.item.Category}
              </Badge>
              <CardTitle className="font-weight-bold">
                {this.props.item.Name}
              </CardTitle>
              {/* <CardSubtitle>{this.props.item.Category}</CardSubtitle> */}

              <p className="text-left">{this.props.item.Description}</p>
              <CardText className="text-right font-weight-bold">
                ${" "}
                {this.props.item.Price.toLocaleString(navigator.language, {
                  minimumFractionDigits: 2
                })}
              </CardText>
              {/* <Button>Add to Cart</Button> */}
              <ModalExample
                buttonLabel="Add to Cart"
                className="buttonLabel"
                item={this.props.item}
                onAddToCartClicked={this.props.onAddToCartClicked}
              />
            </CardBody>
          </Card>
        </React.Fragment>
      </Col>
    );
  }
}

class Inventory extends Component {
  state = {
    items: []
  };

  handleAddToCartClicked = item => {
    console.log("Inventory", item);
    //Post to cart
    axios.post(
      "https://xy0os460h9.execute-api.us-west-2.amazonaws.com/prod/addToCart",
      {
        InventoryID: item.InventoryId,
        Quantity: item.qnty,
        Item: item.Name,
        Price: item.Price,
        UserEmail: "sam.mam@gmail.com"
      }
    );
    //updateCart();
  };

  componentDidMount() {
    axios
      .get(
        `https://9tz63arnh5.execute-api.us-west-2.amazonaws.com/prod/inventory`
      )
      .then(res => {
        const items = res.data;
        console.log(items);
        this.setState({ items });
      });
  }

  render() {
    console.log("State : " + this.state.items);
    return (
      <Container fluid>
        <Header />
        <Row>
          {this.state.items.map(item => (
            <Item
              item={item}
              onAddToCartClicked={this.handleAddToCartClicked}
            />
          ))}
        </Row>
        <Footer />
      </Container>
    );
  }
}

export default Inventory;
