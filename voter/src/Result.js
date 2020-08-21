import React, { Component } from "react";
import roost from "./assets/roost.png";
import k3d from "./assets/k3d.svg";
import kind from "./assets/kind.png";
import minikube from "./assets/minikube.png";
import docker from "./assets/docker.png";
import kubernates from "./assets/kubernates.png";
import "./App.css";
const ballot_endpoint = "roost-controlplane:30080"
const date = new Date();
class Result extends Component {
  constructor(props) {
    super(props);
    this.state = {
      results: [
        {
          id: 0,
          name: "Roost",
          total_votes: 1,
        },
        {
          id: 1,
          name: "Docker",
          total_votes: 5,
        },
        {
          id: 2,
          name: "Minikube",
          total_votes: 4,
        },
        {
          id: 3,
          name: "K3d",
          total_votes: 0,
        },
        {
          id: 4,
          name: "kind",
          total_votes: 0,
        },
      ],
      total: 0,
    };
  }

  componentDidMount() {
    fetch(`http://${ballot_endpoint}`, {
      method: "GET",
    })
      .then((response) => response.json())
      .then((response) => {
        console.log(response);
        const resultData = response;
        // const resultData = {
        //   results: [
        //     {
        //       id: 0,
        //       name: "Roost",
        //       total_votes: 1,
        //     },
        //     {
        //       id: 1,
        //       name: "Docker",
        //       total_votes: 5,
        //     },
        //     {
        //       id: 2,
        //       name: "Minikube",
        //       total_votes: 4,
        //     },
        //     {
        //       id: 3,
        //       name: "K3d",
        //       total_votes: 0,
        //     },
        //     {
        //       id: 4,
        //       name: "kind",
        //       total_votes: 0,
        //     },
        //   ],
        // };
        resultData.results.sort((a, b) => {
          return b.total_votes - a.total_votes;
        });
        let total = 0;
        resultData.results.forEach((item) => {
          total += item.total_votes;
        });

        this.setState({ total: total });
        this.setState({ results: resultData.results });
      });
  }

  render() {
    const CustomCard = (candidate) => {
      return (
        <div className="card">
          <div className="cardBackgroundContainer">
            <div className="cardBackground"></div>
            <div className="cardBackgroundImage">
              {candidate.id == 0 ? (
                <img
                  src={roost}
                  width="150px"
                  height="150px"
                  className="image"
                />
              ) : null}
              {candidate.id == 1 ? (
                <img
                  src={docker}
                  width="150px"
                  height="150px"
                  className="image"
                />
              ) : null}
              {candidate.id == 2 ? (
                <img
                  src={minikube}
                  width="150px"
                  height="150px"
                  className="image"
                />
              ) : null}
              {candidate.id == 3 ? (
                <img src={k3d} width="150px" height="150px" className="image" />
              ) : null}
              {candidate.id == 4 ? (
                <img
                  src={kind}
                  width="150px"
                  height="150px"
                  className="image"
                />
              ) : null}
            </div>
          </div>
          <div className="cardContent">
            {candidate.name}
            <div class="progressbar_back">
              <div
                class="progressbar_front"
                style={{
                  width: `${Math.round(
                    (candidate.total_votes / this.state.total) * 100
                  )}%`,
                }}
              ></div>
              <div>
                {Math.round((candidate.total_votes / this.state.total) * 100)}%
              </div>
            </div>
          </div>
        </div>
      );
    };
    return (
      <div className="Home">
        <div className="logo">
          <img src={kubernates} width="70px" height="70px" />
        </div>
        <div className="heading_result">
          Developer prefer the below tool for building K8S cluster<br />(As of Date: {date.toLocaleString()})
        </div>
        <div className="cardContainer">
          {this.state.results.map((candidate, index) => {
            return CustomCard(candidate, index);
          })}
        </div>
      </div>
    );
  }
}

export default Result;
