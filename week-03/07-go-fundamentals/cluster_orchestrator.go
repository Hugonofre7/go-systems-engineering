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
type StaticResource struct {
	Name string
}

func findNode(id string) *Node {
	if id == "unknown" {
		return nil
	}

	return &Node{
		ID: id,
	}
}

func getResource(id string) Resource {
	return findNode(id)
}

func describe(r Resource) {
	switch v := r.(type) {
	case *Node:
		fmt.Printf("Node resource: %s\n", v.ID)

	case *Container:
		fmt.Printf("Container resource: %s\n", v.ImageName)

	default:
		fmt.Println("Unknown resource type")
	}
}

func (s StaticResource) Start() error {
	fmt.Printf("Starting static resource %s\n", s.Name)
	return nil
}

func (s StaticResource) Stop() error {
	fmt.Printf("Stopping static resource %s\n", s.Name)
	return nil
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

	staticValue := StaticResource{
		Name: "config-map",
	}

	staticPointer := &StaticResource{
		Name: "config-map",
	}

	resources := []Resource{
		node,
		container,
		staticValue,
		staticPointer,
	}

	for _, resource := range resources {
		if err := manageLifecycle(resource); err != nil {
			fmt.Printf("Lifecycle error: %v\n", err)
		}
	}

	for _, resource := range resources {
		describe(resource)
	}

	r := getResource("unknown")

	fmt.Println("Resource is nil:", r == nil)
}
