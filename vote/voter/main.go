package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"sort"
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
	Vote string
	// Status
	Status string
}

type BallotPaper struct {
	Voter string `json:"Voter"`
	Vote  string `json:"Vote"`
}

// Keep same candidateID:Name as in its Voter microservice.
// We can keep "Candidates As A Services" model via API or package based approach
func candidates() map[int]string {
	voters := map[int]string{
		0: "Dog",
		1: "Cat",
		2: "Parrot",
		3: "Rabbit",
		4: "Iguana",
		5: "Python",
	}
	return voters
}

func saveToballot(ballot string, ballotPaper map[string][]string) (status string) {
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

	pagedata := VoterPageData{}
	pagedata.Voter = voter

	candidates := candidates()
	for id, name := range candidates {
		candidate := Candidate{ID: id, Name: name}
		pagedata.Candidates = append(pagedata.Candidates, candidate)
	}
	// Sort candidates as map data is unordered
	sort.Slice(pagedata.Candidates, func(i, j int) bool { return pagedata.Candidates[i].ID < pagedata.Candidates[j].ID })

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalf("unable to parse template file. Error: %v", err)
	}

	// if already voted, display his vote
	pagedata.Vote = r.FormValue("vote")

	if r.Method == http.MethodPost {
		ballotPaper := url.Values{"voter": {voter}, "vote": {pagedata.Vote}}
		pagedata.Status = saveToballot(ballotEndpoint, ballotPaper)
	}
	tmpl.Execute(w, pagedata)
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
