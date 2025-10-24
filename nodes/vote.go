package nodes

import (
	"fmt"
	"math/rand"
	"time"
)

// simulate leader crash
func CrashLeader(leader *Node,delay time.Duration){
	time.Sleep(delay)
	leader.Alive= false
	fmt.Printf("Leader Node [%d] has crashed!\n", leader.ID)
}


// simulate vote casting
func CastVotes(nodes []*Node){
	candidates:= GetCandidateIDs(nodes)
	if len(candidates)==0{
		fmt.Println("No candidates available to vote for.")
		return
	}

	fmt.Println("\n==VOTING PHASE==")

	for _, node:= range nodes{
		if node.Type==Follower && node.Alive{
			node.Vote=candidates[rand.Intn(len(candidates))]
		}
		fmt.Printf("Node [%d] [%v]", node.ID, node.Type)
		if node.Alive{
			if node.Type==Follower{
				fmt.Printf(" voted for Candidate Node [%d]\n", node.Vote)
			}else{
				fmt.Println(" did not vote (not a Follower)")
			}
		}else{
			fmt.Println(" is down and could not vote")
		}
	}
}


// Tally Votes
func TallyVotes(nodes []*Node) int {
	voteCount := make(map[int]int)
	for _, node := range nodes {
		if node.Type == Follower && node.Alive {
			voteCount[node.Vote]++
		}
	}

	fmt.Println("\n== VOTE TALLY ==")
	maxVotes := 0
	var topCandidates []int

	for candidateID, count := range voteCount {
		fmt.Printf("Candidate Node [%d] received [%d] votes\n", candidateID, count)
		if count > maxVotes {
			maxVotes = count
			topCandidates = []int{candidateID}
		} else if count == maxVotes {
			topCandidates = append(topCandidates, candidateID)
		}
	}

	if len(topCandidates) == 1 {
		newLeaderID := topCandidates[0]
		fmt.Printf("New Leader elected: Candidate Node [%d] with %d votes\n", newLeaderID, maxVotes)
		return newLeaderID
	} else if len(topCandidates) > 1 {
		fmt.Printf("Tie detected among nodes: %v. Revoting required...\n", topCandidates)
		return -1
	} else {
		fmt.Println("No votes cast. No leader elected.")
		return -1
	}
}