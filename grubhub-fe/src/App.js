import React, { Component } from "react";
import "./App.css";
import { BrowserRouter } from "react-router-dom";
import Main from "./components/Main";
import "bootstrap/dist/css/bootstrap.min.css";

class App extends Component {
  render() {
    return (
      <BrowserRouter>
        <div>
          <Main />
        </div>
      </BrowserRouter>
    );
  }
}

export default App;
