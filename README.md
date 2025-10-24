# Load Balancer - Setup & Usage Guide

Mini RaFT-Consensus (Reliable and Fault-Tolerant)

---

## Implemented Features

- Simulates leader, candidate, and follower nodes in a distributed cluster.
- Supports voting and leader election among candidate nodes.
- Implements leader crash and recovery, demonstrating fault tolerance.
- Tie resolution: automatic revoting if multiple candidates receive the same maximum votes.
- Configurable number of nodes, number of candidates, and crash timing.

## Local setup

### 1. Clone the Repository

- clone

  ```bash
  git clone https://github.com/rohithr018/mini-raft.git
  ```

- change directory
  ```bash
  cd mini-raft
  ```

---

### 2. Configure Nodes

- Adjust Number of Nodes at **main.go** file

  ```go
    numberOfNodes := X          //Total Number of Nodes in the System
    numberOfCandidates := Y     //Number or Candidate Nodes(must be lessthan numberOfNodes)
  ```

---

### 3. Run simulation

- Run the main file:

  ```bash
  go run main.go
  ```

---

### 4. Reference

- Terminal Images

  - [Intial Node Status](./ref-images/Inital-node-status.png)
  - [Normal Election Scenario](./ref-images/normal-election-scenario.png)
  - [Tie Election Scenario](./ref-images/tie-election-scenario.png)
