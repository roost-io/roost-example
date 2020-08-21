package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
)

var port string = "80"
var once sync.Once

// global candidate list for fast lookup, if candidate exists; map[CandidateID] = TotalVotes
// As Ballot is black boxed and don't know who care the elective candidates, ballot service has to iterate
// over all incoming caldidates data and keep updating it's store with new candidates
// use getCandidates() function which implements singleton instance
var electiveCandidates map[string]int

// global ballotStore is a in-memory store to keep voting data.
// use getBallot() function which implements singleton instance
var ballotStore *Ballot

// Candidate participating in election
type Candidate struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// BallotPaper contains every candidates information, voter and candidate ID to whome vote is given
type BallotPaper struct {
	Candidates []Candidate `json:"candidates"`
	// vote contains cadidate's ID to whome vote is given
	Vote string `json:"vote"`
	// voter identity (any random names received from frontend)
	Voter string `json:"voter"`
}

// CandidateVotes contains candidates and their vote counts
type CandidateVotes struct {
	Candidate
	TotalVotes int `json:"total_votes"`
}

// Ballot contains total vote count and information of the elective candidates
type Ballot struct {
	data []*CandidateVotes
}

// Response to send when result requested
type Response struct {
	Results []CandidateVotes `json:"results"`
	Status  Status           `json:"status"`
}

// Status in reponse
type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// getCandidates output map if map[CandidateID] = CandidatesTotalVotes
func getCandidates() map[string]int {
	once.Do(func() {
		electiveCandidates = make(map[string]int)
	})
	return electiveCandidates
}

func getBallot() *Ballot {
	if ballotStore == nil {
		ballotStore = &Ballot{}
	}
	return ballotStore
}

func saveBallotPaper(paper BallotPaper) error {
	existingCandidates := getCandidates()
	ballot := getBallot()

	// If candidate exists, increase his vote count
	_, candidateFound := existingCandidates[paper.Vote]
	if candidateFound {
		existingCandidates[paper.Vote]++
		for _, candidate := range ballot.data {
			if candidate.ID == paper.Vote {
				candidate.TotalVotes = candidate.TotalVotes + 1
			}
		}
		return nil
	}

	// vote must have been given to valid candidate ID coming from frontend
	// if sent candidate info are new, add them into the store;
	isValidVote := false
	for _, candidate := range paper.Candidates {
		// Add new candidates to the store
		if _, found := existingCandidates[candidate.ID]; !found {
			newCandidate := CandidateVotes{Candidate: candidate}
			existingCandidates[candidate.ID] = 0
			// vote MUST be given to one of new Candidates
			if paper.Vote == candidate.ID {
				isValidVote = true
				newCandidate.TotalVotes++
				existingCandidates[candidate.ID]++
			}
			ballot.data = append(ballot.data, &newCandidate)
		}
	}
	if !isValidVote {
		return fmt.Errorf("Vote does not match to valid elective candidate. voting to candidateID: %v", paper.Vote)
	}
	return nil
}

func writeVoterResponse(w http.ResponseWriter, status Status) {
	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(status)
	if err != nil {
		log.Println("error marshaling response to vote request. error: ", err)
	}
	w.Write(resp)
}

func serveRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	switch r.Method {
	case http.MethodGet:
		defer r.Body.Close()
		log.Println("result request received")
		res := Response{}
		ballot := getBallot()
		for _, b := range ballot.data {
			res.Results = append(res.Results, CandidateVotes{Candidate: b.Candidate, TotalVotes: b.TotalVotes})
		}

		log.Printf("result data: %+v", res.Results)
		if len(res.Results) > 0 {
			res.Status.Code = http.StatusOK
			res.Status.Message = "Publish results"
		} else {
			res.Status.Code = http.StatusNoContent
			res.Status.Message = "Voting is not done yet. Result is not available"
		}

		out, err := json.Marshal(res)
		if err != nil {
			log.Println("error marshaling response to result request. error: ", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)

	case http.MethodPost:
		log.Println("post request received. expect ballot paper")
		ballotPaper := BallotPaper{}
		status := Status{}
		defer r.Body.Close()

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&ballotPaper)
		if err != nil {
			log.Printf("error parsing ballot paper. error: %v\n", err)
			status.Code = http.StatusBadRequest
			status.Message = "Ballot paper is not valid. Vote can not be saved"
			writeVoterResponse(w, status)
			return
		}

		log.Printf("ballot paper data: %+v", ballotPaper)
		err = saveBallotPaper(ballotPaper)
		if err != nil {
			log.Println(err)
			status.Code = http.StatusBadRequest
			status.Message = "Ballot paper is not valid. Vote can not be saved"
			writeVoterResponse(w, status)
			return
		}

		bellotData := getBallot()
		log.Printf("existingCandidates: %+v", getCandidates())
		for i, v := range bellotData.data {
			log.Printf("ballot result [%d]: %+v", i, *v)
		}
		status.Code = http.StatusCreated
		status.Message = "Vote saved suessfully"
		writeVoterResponse(w, status)
		return

	default:
		status := Status{}
		status.Code = http.StatusMethodNotAllowed
		status.Message = "Bad Request. Vote can not be saved"
		writeVoterResponse(w, status)
		return
	}

}

func main() {
	http.HandleFunc("/", serveRoot)
	log.Println(http.ListenAndServe(net.JoinHostPort("", port), nil))
}
