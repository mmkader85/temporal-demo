package main

import (
	"context"
	"fmt"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/workflow"
)

func main() {
	// Create the client object just once per process
	c, err := client.NewLazyClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "001-hello-world-workflow",
		TaskQueue: "001-hello-world",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, HelloWorldWorkflow)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}

	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("unable to get Workflow result", err)
	}

	fmt.Printf("Result: %s\n", result)
}

// HelloWorldWorkflow is a dummy method to satisfy the type requirement for the c.ExecuteWorkflow call
func HelloWorldWorkflow(ctx workflow.Context) (string, error) {
	return "", nil
}
