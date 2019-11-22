import React, { useState } from "react";
import { Button, Modal, ModalHeader, ModalBody, ModalFooter } from "reactstrap";
import NumericInput from "react-numeric-input";

const ModalExample = props => {
  const { buttonLabel, className, item } = props;

  const [modal, setModal] = useState(false);

  const toggle = () => setModal(!modal);
  var Quantity = 0;

  return (
    <div>
      <Button color="primary" onClick={toggle}>
        {buttonLabel}
      </Button>
      <Modal isOpen={modal} toggle={toggle} className={className}>
        <ModalHeader toggle={toggle}>{item.Name}</ModalHeader>
        <ModalBody>
          Select Quantity {"  "}{" "}
          <NumericInput
            min={0}
            max={100}
            value={0}
            onChange={value => (item["qnty"] = value)}
          />
        </ModalBody>
        <ModalFooter>
          <Button
            color="danger"
            onClick={() => {
              props.onAddToCartClicked(item);
              toggle();
            }}
          >
            Add to Cart
          </Button>{" "}
        </ModalFooter>
      </Modal>
    </div>
  );
};

export default ModalExample;
