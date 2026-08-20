package main

import "fmt"

type Resource interface {
	Start() error
	Stop() error
}

type Node struct {
	ID string
}

func (n *Node) Start() error {
	fmt.Printf("Starting node %s\n", n.ID)
	return nil
}

func (n *Node) Stop() error {
	fmt.Printf("Stopping node %s\n", n.ID)
	return nil
}

type Container struct {
	ImageName string
}

func (c *Container) Start() error {
	fmt.Printf("Starting container %s\n", c.ImageName)
	return nil
}

func (c *Container) Stop() error {
	fmt.Printf("Stopping container %s\n", c.ImageName)
	return nil
}

func manageLifecycle(r Resource) error {
	if err := r.Start(); err != nil {
		return err
	}

	return r.Stop()
}

func main() {
	node := &Node{
		ID: "node-1",
	}

	container := &Container{
		ImageName: "nginx:latest",
	}

	resources := []Resource{
		node,
		container,
	}

	for _, resource := range resources {
		if err := manageLifecycle(resource); err != nil {
			fmt.Printf("Lifecycle error: %v\n", err)
		}
	}
}
