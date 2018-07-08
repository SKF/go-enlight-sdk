package main

import (
	"time"

	"github.com/SKF/go-utility/env"
	"github.com/SKF/go-utility/log"
	"github.com/SKF/go-utility/uuid"

	"github.com/SKF/go-enlight-sdk/grpc"
	"github.com/SKF/go-enlight-sdk/services/iot"
	api "github.com/SKF/go-enlight-sdk/services/iot/iot_grpc_api"
)

func main() {
	HOST := env.GetAsString("HOST", "grpc.sandbox.iot.enlight.skf.com")
	PORT := env.GetAsString("PORT", "50051")
	CLIENT_CRT := env.GetAsString("CLIENT_CRT", "cert/client.crt")
	CLIENT_KEY := env.GetAsString("CLIENT_KEY", "cert/client.key")
	CA_CERT := env.GetAsString("CA_CERT", "cert/ca.crt")
	CERT_NAME := env.GetAsString("CERT_NAME", "grpc.sandbox.iot.enlight.skf.com")

	userID := uuid.New().String()
	nodeID := uuid.New().String()
	var taskID string
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

	log.Info("Get All Tasks")
	if _, err = client.GetAllTasks(userID); err != nil {
		log.
			WithError(err).
			WithField("userId", userID).
			Error("client.GetAllTasks")
		return
	}

	log.Info("Get Uncompleted Tasks")
	if _, err = client.GetUncompletedTasks(userID); err != nil {
		log.
			WithError(err).
			WithField("userId", userID).
			Error("client.GetUncompletedTasks")
		return
	}

	log.Info("Get Uncompleted Tasks By Hierarachy")
	if _, err = client.GetUncompletedTasksByHierarchy(nodeID); err != nil {
		log.
			WithError(err).
			WithField("nodeID", nodeID).
			Error("client.GetUncompletedTasksByHierarchy")
		return
	}

	log.Info("Create Task")
	createTaskInput := api.InitialTaskDescription{
		UserId:           userID,
		TaskName:         "MyTaskName",
		HierarchyId:      uuid.New().String(),
		DueDateTimestamp: time.Now().Unix() * 1000,
		FunctionalLocationIds: &api.FunctionalLocationIds{
			IdArr: []string{
				uuid.New().String(),
				uuid.New().String(),
			},
		},
		ExternalTaskId: uuid.New().String(),
	}

	if taskID, err = client.CreateTask(createTaskInput); err != nil {
		log.WithError(err).Error("client.CreateTask")
		return
	}

	log.Info("Set Task Status")
	if err = client.SetTaskStatus(userID, taskID, api.TaskStatus_RECEIVED); err != nil {
		log.
			WithError(err).
			WithField("userId", userID).
			WithField("taskId", taskID).
			Error("client.SetTaskCompleted")
		return
	}

	log.Info("Set Task Completed")
	if err = client.SetTaskCompleted(userID, taskID); err != nil {
		log.
			WithError(err).
			WithField("userId", userID).
			WithField("taskId", taskID).
			Error("client.SetTaskCompleted")
		return
	}

	log.Info("Delete Task")
	if err = client.DeleteTask(userID, taskID); err != nil {
		log.
			WithError(err).
			WithField("userId", userID).
			WithField("taskId", taskID).
			Error("client.DeleteTask")
		return
	}
}
