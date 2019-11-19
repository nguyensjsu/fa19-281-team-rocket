import React, { Component } from "react";
import {  Card, CardImg, CardText, CardBody, CardTitle, CardSubtitle, Button } from "reactstrap";
import { Row, Col, Container } from 'reactstrap';
import "../App.css";
import axios from 'axios';
import Header from './Header'


class Item extends Component {
    constructor(props){
        super(props);
    }
    render() {
        console.log(this.props)
        return (
        <Col sm="4">
        <Card>
        <CardImg top width="100%" src={this.props.item.Image} alt="Card image cap" />
        <CardBody>
        <CardTitle>{this.props.item.Name}</CardTitle>
        <CardSubtitle>{this.props.item.Category}</CardSubtitle>
        <CardBody>{this.props.item.Description}</CardBody>
        <CardText>{this.props.item.Price} Dollars</CardText>
         <Button>Add to Cart</Button>
        </CardBody>
      </Card>
      </Col>
        );
    }
}

class Inventory extends Component {
    state = {
        items: []
    }

    componentDidMount() {
        axios.get(`https://9tz63arnh5.execute-api.us-west-2.amazonaws.com/prod/inventory`)
          .then(res => {
            const items = res.data;
            console.log(items);
            this.setState({ items });
          })
    }

    render() {
        console.log("State : " + this.state.items)
        return (
           
            <Container fluid>
             <Header/>
            <Row>
                {this.state.items.map(item => <Item item={item}/>)}
            </Row>
            </Container>
        )
    }

}

export default Inventory;