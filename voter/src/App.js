import React, { Component } from "react";
import roost from "./assets/roost.png";
import k3d from "./assets/k3d.svg";
import kind from "./assets/kind.png";
import minikube from "./assets/minikube.png";
import docker from "./assets/docker.png";
import kubernates from "./assets/kubernates.png";
import "./App.css";

const ballot_endpoint = process.env.REACT_APP_BALLOT_ENDPOINT || ""
class App extends Component {
  constructor(props) {
    super(props);
    // this.handleonCardClick = this.handleonCardClick.bind(this)
    this.state = {
      candidates: [
        {
          Name: "Roost",
          ID: "0",
        },
        {
          Name: "Docker Desktop",
          ID: "1",
        },
        {
          Name: "Minikube",
          ID: "2",
        },
        {
          Name: "K3d",
          ID: "3",
        },
        {
          Name: "Kind",
          ID: "4",
        },
      ],
      vote: "",
      // username: "",
      disabled: false
    };
  }

  // componentDidMount() {
  //   let r = Math.random().toString(36).substring(7);
  //   this.setState({ username: r });
  // }

  componentDidUpdate(prevProps, prevState) {
    if(prevState.vote != this.state.vote) {
      const data = {
        candidates: [...this.state.candidates], vote: this.state.vote, username: this.state.username
      }
      console.log("state: ", this.state)
      console.log("state: ", data)
      console.log("ballot endpoint is: ", ballot_endpoint)
      console.log("env: ", process.env)
      if(ballot_endpoint === "") {
        console.error("ballot endpoint is not set");
      } else {
        fetch(`http://${ballot_endpoint}`, {
          method: "POST",
          body: JSON.stringify(data),
        })
      .then(response => response.json())
      .then(response => {console.log(response)});
      }
    }
  }

  render() {
    const handleonCardClick = (e) => {
      if(this.state.disabled == false) {
        console.log(e.target.className);
        let targetHtml = e.target.innerHTML;
        let targetElement = e.target;
        if (e.target.className == "cardBackgroundContainer") {
          targetHtml = e.target.parentElement.children[1].innerHTML;
          targetElement = e.target.parentElement.children[1];
        } else if (
          e.target.className == "cardBackground" ||
          e.target.className == "cardBackgroundImage"
        ) {
          targetHtml = e.target.parentElement.parentElement.children[1].innerHTML;
          targetElement = e.target.parentElement.parentElement.children[1];
        } else if (e.target.className == "card") {
          targetHtml = e.target.children[1].innerHTML;
          targetElement = e.target.children[1];
        } else if (e.target.className == "image") {
          targetHtml =
            e.target.parentElement.parentElement.parentElement.children[1]
              .innerHTML;
          targetElement =
            e.target.parentElement.parentElement.parentElement.children[1];
        }
        this.state.candidates.forEach((candidate) => {
          if (candidate.Name == targetHtml) {
            targetElement.parentElement.classList.add("selectedCard");
            this.setState({ vote: candidate.ID });
            this.setState({disabled: true})
          }
        });
      }
    };
    const CustomCard = (candidate) => {
      return (
        <div className="card" onClick={(e) => handleonCardClick(e)}>
          <div className="cardBackgroundContainer">
            <div className="cardBackground"></div>
            <div className="cardBackgroundImage">
              {candidate.ID == 0 ? (
                <img
                  src={roost}
                  width="150px"
                  height="150px"
                  className="image"
                />
              ) : null}
              {candidate.ID == 1 ? (
                <img
                  src={docker}
                  width="150px"
                  height="150px"
                  className="image"
                />
              ) : null}
              {candidate.ID == 2 ? (
                <img
                  src={minikube}
                  width="150px"
                  height="150px"
                  className="image"
                />
              ) : null}
              {candidate.ID == 3 ? (
                <img src={k3d} width="150px" height="150px" className="image" />
              ) : null}
              {candidate.ID == 4 ? (
                <img
                  src={kind}
                  width="150px"
                  height="150px"
                  className="image"
                />
              ) : null}
            </div>
          </div>
          <div className="cardContent">{candidate.Name}</div>
        </div>
      );
    };

    return (
      <div className="App">
        <div className="logo">
        <img src={kubernates} width="70px" height="70px" />
        </div>
        <div className="heading">Question ?</div>
        <div className="cardContainer">
          {this.state.candidates.map((candidate, index) => {
            return CustomCard(candidate, index);
          })}
        </div>
      </div>
    );
  }
}

export default App;
