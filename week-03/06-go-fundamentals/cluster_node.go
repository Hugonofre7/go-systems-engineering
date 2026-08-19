package main

import "fmt"

type Node struct {
	ID      string
	Healthy bool
	Uptime  int
}

func (n *Node) SetHealthy(h bool) {
	n.Healthy = h
}

func (n Node) Status() string {
	return fmt.Sprintf(
		"Node %s | Healthy: %t | Uptime: %ds",
		n.ID,
		n.Healthy,
		n.Uptime,
	)
}

func (n Node) BrokenSetHealthy(h bool) {
	n.Healthy = h
}

func main() {
	node := Node{
		ID:      "node-1",
		Healthy: false,
		Uptime:  120,
	}

	node.BrokenSetHealthy(true)

	fmt.Println("After BrokenSetHealthy:", node.Status())

	node.SetHealthy(true)

	fmt.Println("After SetHealthy:", node.Status())
}
