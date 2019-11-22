import React, {Component} from 'react';
import { Row, Col, Container } from 'reactstrap';
import Header from './Header';

class Order extends Component{
    render(){
        return(
            <Container fluid>
            <Header/>
            <div class="col-md-3"  style = { { left : "250px", top : "100px" ,width : "250px",color : "green", border: "1px solid green", padding :"10px"}  }>

            <h3> Order placed</h3>
            </div>
            </Container>
        
        )
    }
}

export default Order;