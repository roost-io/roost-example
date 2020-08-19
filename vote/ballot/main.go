package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"
)

var (
	port          string = "82"
	candidateList map[int]string
	ballotStore   *ballot
)

type BallotResult struct {
	CandidateID   int    `json:"candidateID"`
	CandidateName string `json:"candidateName"`
	Votes         int    `json:"votes"`
}

type ballot struct {
	mu sync.Mutex
	// map[voter]=votedTo/candidate
	ballotPaper map[string]int
	counter     map[int]int
}

// Keep same candidateID:Name as in its Voter microservice.
// We can keep "Candidates As A Services" model via API or package based approach
func candidates() map[int]string{
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

func serveRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		log.Println("Request URI: ", r.RequestURI)

		//only GET /result would be served
		if r.RequestURI != "/result" {
			w.Write([]byte("You don't have permission to look into Ballot"))
		} else {
			//looking for result
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.Write(MarshalResult())
		}
	}
	if r.Method == http.MethodPost {
		voter := r.FormValue("voter")
		voteIn := r.FormValue("vote")
		vote, err := strconv.Atoi(voteIn)
		if err != nil {
			log.Println("error converting vote from str to int")
		}
		if voter == "" || vote == 0 {
			log.Println("forging vote detected. voter: ", voter, " vote: ", vote)
		}
		b := getBallot()
		b.add(voter, vote)
		log.Printf("ballotpaper received. Voter: %v, voted to %v", voter, vote)
		log.Printf("Result: %#v", getResult())
		w.Write([]byte("ballot paper received"))
	}
}

// Initialize candidates.
func init() {
	candidateList = candidates()
}

func main() {
	log.Println("listening on port: ", port)

	http.HandleFunc("/", serveRoot)
	log.Fatal(http.ListenAndServe(net.JoinHostPort("", port), nil))
}

func (b *ballot) add(voter string, vote int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	candidate, ok := b.ballotPaper[voter]
	b.ballotPaper[voter] = vote
	// voting for first time
	if !ok {
		b.counter[vote]++
	}
	// changing votes
	if ok && candidate != vote {
		b.counter[candidate]--
		b.counter[vote]++
	}
}

func (b *ballot) result() map[int]int {
	return b.counter
}

func getBallot() *ballot {
	if ballotStore == nil {
		ballotStore = &ballot{
			ballotPaper: make(map[string]int),
			counter:     make(map[int]int),
		}

		// Initialize candidates votes so that in Result page vote counts would be shown
		for candidateID := range candidateList {
			ballotStore.counter[candidateID] = 0
		}
	}
	return ballotStore
}

func getResult() map[int]int {
	b := getBallot()
	return b.result()
}

// MarshalResult returns json output
func MarshalResult() []byte {
	result := getResult()
	var BallotResults []BallotResult
	log.Printf("got result in marchal: %v", result)
	for id, count := range result {
		BallotResults = append(BallotResults, BallotResult{CandidateID: id, CandidateName: candidateList[id], Votes: count})
	}
	log.Printf("Marshaled result: %+v", BallotResults)
	out, err := json.Marshal(BallotResults)
	if err != nil {
		log.Println("error marchaling ballot result. error: ", err)
	}
	return out
}
