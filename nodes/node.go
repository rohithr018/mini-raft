package nodes

import "fmt"

type NodeType string

const(
	Leader  	NodeType="Leader"
	Follower 	NodeType="Follower"
	Candidate 	NodeType="Candidate"
)


// Node Structure
type Node struct{
	ID      	int
	Type    	NodeType
	Alive   	bool
	Vote		int
}

func InitializeNodes(n int,c int)[]*Node{
	nodes:= make([]*Node, n)

	// Initial Leader
	nodes[0]= &Node{ID: 1, Type: Leader, Alive: true}

	// Candidates
	candidateCount:=0
	for i:=1; i<n && candidateCount<c; i++{
		nodes[i]= &Node{ID: i+1, Type: Candidate, Alive: true}
		candidateCount++
	}
	// Followers
	for i:=0;i<n;i++{
		if nodes[i]==nil{
			nodes[i]= &Node{ID: i+1, Type: Follower, Alive: true}
		}
	}
	return nodes
}


// Get CandiateID
func GetCandidateIDs(nodes []*Node)[]int{
	var candidateIDs []int
	for _, node:= range nodes{
		if node.Type==Candidate && node.Alive{
			candidateIDs= append(candidateIDs, node.ID)
		}
	}
	return candidateIDs
}

// Display Nodes
func DisplayNodes(nodes []*Node){
	fmt.Println("\n==SYSTEM NODES STATUS==")
	fmt.Println("\nLeader:")
	for _, node:= range nodes{
		if node.Type==Leader{
			fmt.Printf("Node [%d]--Alive: [%v]\n", node.ID, node.Alive)
		}
	}
	fmt.Println("\nCandidates:")
	for _, node:= range nodes{
		if node.Type==Candidate{
			fmt.Printf("Node [%d]--Alive: [%v] \n", node.ID, node.Alive)
		}
	}
	fmt.Println("\nFollowers:")
	for _, node:= range nodes{
		if node.Type==Follower{
			fmt.Printf("Node [%d]--Alive: [%v] \n", node.ID, node.Alive)
		}
	}
	fmt.Println("========================")
}

