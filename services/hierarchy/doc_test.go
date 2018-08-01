package hierarchy_test

import (
	"log"

	"github.com/SKF/go-enlight-sdk/grpc"
	"github.com/SKF/go-enlight-sdk/services/hierarchy"
	"github.com/SKF/go-enlight-sdk/services/hierarchy/grpcapi"
)

func ExampleClient() {
	host, port := "<host>", "<port>"
	clientCert, clientKey := "<clientCertPath>", "<clientKey>"
	caCert := "<caCert>"

	// Create a Hierarchy Service client
	client := hierarchy.CreateClient()

	// Dial the Hierarchy Service
	dialOption, err := grpc.WithTransportCredentials(host, clientCert, clientKey, caCert)
	if err != nil {
		log.Fatalf("Couldn't connect due to error: %+v", err)
	}

	// Dial the Hierarchy Service
	err = client.Dial(host, port, dialOption)
	if err != nil {
		log.Fatalf("Couldn't dial due to error: %+v", err)
	}
	defer client.Close()

	// Ping the Hierarchy Service
	err = client.DeepPing()
	if err != nil {
		log.Fatalf("Couldn't ping the Hierarchy Service due to error: %+v", err)
	}
}

func ExampleClient_SaveNode_create() {
	var client hierarchy.HierarchyClient
	saveNodeInput := grpcapi.SaveNodeInput{
		UserId:   "<user_id>",
		Node:     &grpcapi.Node{},
		ParentId: "<parent_id>",
	}
	nodeID, err := client.SaveNode(saveNodeInput)
	if err != nil {
		log.Printf("Couldn't save node due to error: %+v", err)
		return
	}
	log.Printf("nodeId: %q", nodeID)
}

func ExampleClient_SaveNode_update() {
	var client hierarchy.HierarchyClient
	saveNodeInput := grpcapi.SaveNodeInput{
		UserId: "<user_id>",
		Node: &grpcapi.Node{
			Id: "<node_id>",
		},
		ParentId: "<parent_id>",
	}
	nodeID, err := client.SaveNode(saveNodeInput)
	if err != nil {
		log.Printf("Couldn't save node due to error: %+v", err)
		return
	}
	log.Printf("nodeId: %q", nodeID)
}
