package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"math"
	"net"
	"net/http"
	"os"
	"strings"
)

var (
	port           string = "83"
	ballotEndpoint string = ""
	result         map[string]int
)

type ViewData struct {
	Results Results
	Error   string
}

type Results struct {
	TotalVotes      int
	CandidateResult []BallotResult
}

type BallotResult struct {
	CandidateID   int    `json:"candidateID"`
	CandidateName string `json:"candidateName"`
	Votes         int    `json:"votes"`
}

// CandidateVotePercentage returns of vote percentage of candidate
func (r *Results) CandidateVotePercentage(id int) int {

	for _, res := range r.CandidateResult {
		log.Println("Percentage ID: ", id)
		log.Println("Percentage CandidateID: ", res.CandidateID, "his Votes:", res.Votes, "Total Votes: ", r.TotalVotes)
		if res.CandidateID == id && r.TotalVotes != 0 {
			a := int(math.Ceil((float64(res.Votes) / float64(r.TotalVotes)) * 100))
			log.Println("Calculating percentage. Got Per: ", a)
			return a
		}
	}
	return 0
}

func Hello() string {
	return "Hey Man"
}

func serveRoot(w http.ResponseWriter, r *http.Request) {
	var result Results

	if r.Method == http.MethodGet {
		resp, err := http.Get(ballotEndpoint)
		if err != nil {
			log.Println("error getting voting result. error: ", err)
		}
		if resp == nil {
			w.Write([]byte("Ballot service is unavailable."))
			return
		}
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Println("error reading results. error: ", err)
		}

		hello := template.FuncMap{
			"hell": Hello,
			"per":  result.CandidateVotePercentage,
		}
		tmpl, err := template.New("index.html").Funcs(hello).ParseFiles("index.html")
		if err != nil {
			log.Println("error parsing template. error: ", err)
		}

		if strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
			err := json.Unmarshal(data, &result.CandidateResult)
			if err != nil {
				log.Println("error parsing result data, error: ", err)
			}
		}
		// count total votes
		for _, res := range result.CandidateResult {
			result.TotalVotes += res.Votes
		}

		log.Printf("Result: %+v", result)
		log.Printf("Percentage: %v", result.CandidateVotePercentage(1))
		err = tmpl.Execute(w, ViewData{Results: result})
		if err != nil {
			log.Printf("error: %v", err.Error())
		}
	}
}

func main() {
	result = make(map[string]int)

	var ok bool
	ballotEndpoint, ok = os.LookupEnv("BALLOT_ENDPOINT")
	ballotEndpoint = "http://" + ballotEndpoint + "/result"
	if !ok {
		log.Println("Ballot service endpoint is not set as ENV[BALLOT_ENDPOINT]")
	}

	fs := http.FileServer(http.Dir("./views"))
	http.Handle("/views/", http.StripPrefix("/views/", fs))
	http.HandleFunc("/", serveRoot)

	log.Println("listening on port: ", port)
	log.Fatal(http.ListenAndServe(net.JoinHostPort("", port), nil))
}
