package main

import "fmt"

type Logger struct {
	Prefix string
}

func (l *Logger) Log(msg string) {
	fmt.Printf("[%s] %s\n", l.Prefix, msg)
}

type Metrics struct {
	RequestCount int
}

func (m *Metrics) Log(msg string) {
	fmt.Printf("[metrics] %s (requests: %d)\n", msg, m.RequestCount)
}

type Node struct {
	Logger
	Metrics
	ID      string
	Healthy bool
	Uptime  int
}

func (n *Node) SetHealthy(h bool) {
	n.Healthy = h
}

func (n *Node) Status() string {
	return fmt.Sprintf(
		"Node %s | Healthy: %t | Uptime: %ds",
		n.ID,
		n.Healthy,
		n.Uptime,
	)
}

func (n *Node) IncrementUptime(seconds int) {
	n.Uptime += seconds
}

//func (n Node) BrokenSetHealthy(h bool) {
//	n.Healthy = h
//}

func main() {
	node := Node{
		Logger: Logger{
			Prefix: "node-1",
		},
		Metrics: Metrics{
			RequestCount: 42,
		},
		ID:      "node-1",
		Healthy: false,
		Uptime:  120,
	}

	//node.Log("test")
	node.Logger.Log("test")
	node.Metrics.Log("test")

	node.SetHealthy(true)
	node.IncrementUptime(30)
	node.IncrementUptime(45)

	fmt.Println(node.Status())
}
