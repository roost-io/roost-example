package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
)

// Keep same sequence as in its Ballot microservice
const (
	dog = iota + 1
	cat
)

const port string = "81"

var ballotEndpoint string = ""

type Candidate struct {
	ID   int
	Name string
}

type VoterPageData struct {
	Candidates []Candidate
	Voter      string
	// To whome Voter voted
	Vote       string
	// Status
	Status string
}

type BallotPaper struct {
	Voter string `json:"Voter"`
	Vote  string `json:"Vote"`
}

func saveToballot(ballot string, ballotPaper map[string][]string) (status string){
	resp, err := http.PostForm(ballotEndpoint, ballotPaper)
	if err != nil {
		log.Printf("FAIL to send vote to ballot. BalllotEndpoint: %v\t BallotPaper: %v Error: %v\n", ballotEndpoint, ballotPaper, err)
		status = "Ballot service is unavailable."
	}
	if resp != nil {
		if resp.StatusCode == http.StatusOK {		
			log.Println("Vote saved to ballot. Voted to: ", ballotPaper["vote"][0])
			status = "Vote Saved successfully"
		} else {
			log.Printf("ballot response not received. Voted to: %s\t StatusCode: %v\t Status: %v \n", ballotPaper["vote"][0], resp.StatusCode, resp.Status)
			status = "Unable to receive acknowledgement from ballot service."
		}
	}
	return status
}

func serveRoot(w http.ResponseWriter, r *http.Request) {
	voter, err := os.Hostname()
	if err != nil {
		log.Fatalf("not a valid voter. Voter's identity is his unique hostname. Exiting...")
	}
	data := VoterPageData{
		Candidates: []Candidate{
			{ID: dog, Name: "Dog"},
			{ID: cat, Name: "Cat"},
		},
		Voter: voter,
	}
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalf("unable to parse template file. Error: %v", err)
	}

	// if already voted, display his vote
	data.Vote = r.FormValue("vote")

	if r.Method == http.MethodPost {
		ballotPaper := url.Values{"voter": {voter}, "vote": {data.Vote}}
		data.Status = saveToballot(ballotEndpoint, ballotPaper)
	}
	tmpl.Execute(w, data)
}

func main() {
	var ok bool
	ballotEndpoint, ok = os.LookupEnv("BALLOT_ENDPOINT")
	ballotEndpoint = "http://" + ballotEndpoint
	if !ok {
		log.Println("Ballot service endpoint is not set as ENV[BALLOT_ENDPOINT]")
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", serveRoot)

	fmt.Println("listening on port: ", port)
	log.Fatal(http.ListenAndServe(net.JoinHostPort("", port), nil))
}
