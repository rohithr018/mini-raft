package main

import (
	"distributed-voting/nodes"
	"fmt"
	"math/rand"
	"time"
)


func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	numberOfNodes := 10
	numberOfCandidates := 4
	nodeList := nodes.InitializeNodes(numberOfNodes, numberOfCandidates)


	nodes.DisplayNodes(nodeList)

	// initial leader
	var leader *nodes.Node
	for _, node := range nodeList {
		if node.Type == nodes.Leader {
			leader = node
			break
		}
	}


	// simulate simulation single election round
	fmt.Println("\n=== ELECTION SIMULATION ===")
	nodes.CrashLeader(leader, 3*time.Second)

	newLeaderID := -1
	// repeat voting until a single leader is elected
	for newLeaderID == -1 {
		nodes.CastVotes(nodeList)
		newLeaderID = nodes.TallyVotes(nodeList)
		if newLeaderID == -1 {
			fmt.Println("Waiting 2 seconds before revote due to tie...")
			time.Sleep(2 * time.Second)
		}
	}

	// Update roles
	for _, node := range nodeList {
		if node.ID == newLeaderID {
			node.Type = nodes.Leader
		} else if node.Type == nodes.Leader {
			node.Type = nodes.Follower
			node.Alive = true // previous leader recovers
		}
	}

	// Update leader pointer
	for _, node := range nodeList {
		if node.Type == nodes.Leader {
			leader = node
			break
		}
	}

	nodes.DisplayNodes(nodeList)
	fmt.Println("\n== ELECTION SIMULATION ENDED ==")

}
