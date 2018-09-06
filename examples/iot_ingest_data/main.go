package main

import (
	"time"

	"github.com/SKF/go-utility/env"
	"github.com/SKF/go-utility/log"
	"github.com/SKF/go-utility/uuid"
	proto_iot "github.com/SKF/proto/iot"

	"github.com/SKF/go-enlight-sdk/grpc"
	"github.com/SKF/go-enlight-sdk/services/iot"
)

func main() {
	HOST := env.GetAsString("HOST", "grpc.sandbox.iot.enlight.skf.com")
	PORT := env.GetAsString("PORT", "50051")
	CLIENT_CRT := env.GetAsString("CLIENT_CRT", "cert/client.crt")
	CLIENT_KEY := env.GetAsString("CLIENT_KEY", "cert/client.key")
	CA_CERT := env.GetAsString("CA_CERT", "cert/ca.crt")
	CERT_NAME := env.GetAsString("CERT_NAME", "grpc.sandbox.iot.enlight.skf.com")

	var err error

	log.Info("Setup Client")
	client := iot.CreateClient()
	transportOption, err := grpc.WithTransportCredentials(
		CERT_NAME, CLIENT_CRT, CLIENT_KEY, CA_CERT,
	)
	if err != nil {
		log.
			WithError(err).
			WithField("serverName", CERT_NAME).
			WithField("clientCrt", CLIENT_CRT).
			WithField("clientKey", CLIENT_KEY).
			WithField("caCert", CA_CERT).
			Error("grpc.WithTransportCredentials")
		return
	}

	err = client.Dial(
		HOST, PORT,
		transportOption,
		grpc.WithBlock(),
		grpc.FailOnNonTempDialError(true),
	)
	if err != nil {
		log.
			WithError(err).
			WithField("host", HOST).
			WithField("port", PORT).
			Error("client.Dial")
		return
	}

	defer client.Close()

	log.Info("Deep Ping")
	if err = client.DeepPing(); err != nil {
		log.WithError(err).Error("client.DeepPing")
		return
	}

	nodeID1 := uuid.New().String()
	nodeID2 := uuid.New().String()

	log.Info("IngestNodeData")
	for _, nodeData := range createExampleData() {
		nd := *nodeData
		input := proto_iot.IngestNodeDataInput{NodeId: nodeID1, NodeData: &nd}
		err := client.IngestNodeData(input)
		if err != nil {
			log.
				WithError(err).
				WithField("nodeID", nodeID1).
				WithField("nodeData", nodeData).
				Error("client.IngestNodeData")
		}
	}
	log.Info("IngestNodeDataStream")

	doneChannel := make(chan bool)
	dataChannel := make(chan proto_iot.IngestNodeDataStreamInput)
	go func() {
		err := client.IngestNodeDataStream(dataChannel)
		if err != nil {
			log.
				WithError(err).
				Error("client.IngestNodeDataStream")
		}
		doneChannel <- true
	}()
	dataChannel <- proto_iot.IngestNodeDataStreamInput{
		NodeId:       nodeID1,
		NodeDataList: createExampleData(),
	}
	dataChannel <- proto_iot.IngestNodeDataStreamInput{
		NodeId:       nodeID2,
		NodeDataList: createExampleData(),
	}
	close(dataChannel)
	<-doneChannel

	var input proto_iot.GetNodeDataInput
	var output []proto_iot.NodeData

	log.Info("GetNodeData_All")
	input = proto_iot.GetNodeDataInput{
		NodeId: nodeID1,
	}
	output, err = client.GetNodeData(input)
	log.
		WithError(err).
		WithField("input", input).
		WithField("outputLength", len(output)).
		Info("client.GetNodeData")

	log.Info("GetNodeData_DataPoint")
	input = proto_iot.GetNodeDataInput{
		NodeId:      nodeID1,
		ContentType: proto_iot.NodeDataContentType_DATA_POINT,
	}
	output, err = client.GetNodeData(input)
	log.
		WithError(err).
		WithField("input", input).
		WithField("outputLength", len(output)).
		Info("client.GetNodeData")
}

func createExampleData() (out []*proto_iot.NodeData) {
	out = append(out, &proto_iot.NodeData{
		CreatedAt:   time.Now().UnixNano(),
		ContentType: proto_iot.NodeDataContentType_DATA_POINT,
		DataPoint: &proto_iot.DataPoint{
			Coordinate: &proto_iot.Coordinate{X: 1.0, Y: 2.0},
			XUnit:      "<dp-xunit>",
			YUnit:      "<dp-yunit>",
		},
	})

	out = append(out, &proto_iot.NodeData{
		CreatedAt:   time.Now().UnixNano(),
		ContentType: proto_iot.NodeDataContentType_SPECTRUM,
		Spectrum: &proto_iot.Spectrum{
			XUnit: "<s-xunit>",
			YUnit: "<s-yunit>",
		},
	})

	out = append(out, &proto_iot.NodeData{
		CreatedAt:   time.Now().UnixNano(),
		ContentType: proto_iot.NodeDataContentType_TIME_SERIES,
		TimeSeries: &proto_iot.TimeSeries{
			XUnit: "<ts-xunit>",
			YUnit: "<ts-yunit>",
		},
	})

	out = append(out, &proto_iot.NodeData{
		CreatedAt:   time.Now().UnixNano(),
		ContentType: proto_iot.NodeDataContentType_NOTE,
		Note:        "<note>",
	})

	out = append(out, &proto_iot.NodeData{
		CreatedAt:   time.Now().UnixNano(),
		ContentType: proto_iot.NodeDataContentType_MEDIA,
		Media:       []byte("<media>"),
	})

	out = append(out, &proto_iot.NodeData{
		CreatedAt:       time.Now().UnixNano(),
		ContentType:     proto_iot.NodeDataContentType_QUESTION_ANSWERS,
		QuestionAnswers: []string{"<answer>"},
	})

	return
}
