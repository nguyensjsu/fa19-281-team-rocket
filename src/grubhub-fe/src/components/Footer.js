import React from 'react';
import { Card,Container,Row,Col } from 'reactstrap';


function Footer(/*props*/) {
  return (
    <footer style={{paddingTop: '20px'}}>
        <hr></hr>
      <Container>
        <Row>
            <Col sm="4"><p style={{textAlign: 'center'}}>Privacy policy</p></Col>
            <Col sm="4"><p style={{textAlign: 'center'}}>Terms & Conditions</p></Col>
            <Col sm="4"><p style={{textAlign: 'center'}} ><a href="https://github.com/nguyensjsu/fa19-281-team-rocket" style={{color: 'black'}}>Source Code</a></p></Col>
        </Row>
        <div className="text-center small copyright">
          © Team Rocket 2019
        </div>
      </Container>
    </footer>
  );
}

export default Footer;