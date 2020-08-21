import React, { Component } from "react";
import Home from './Home'
import Result from './Result'
import { Router, Link } from "@reach/router";
import "./App.css";

const ballot_endpoint = process.env.REACT_APP_BALLOT_ENDPOINT || ""
class App extends Component {
  constructor(props) {
    super(props);
  }

  render() {
    
    return (
      <div className="App">
        <Router>
          <Home path="/" />
          <Result path="result" />
        </Router>
      </div>
    );
  }
}

export default App;
