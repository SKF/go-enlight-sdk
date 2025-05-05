package models

import (
	iotgrpcapi "github.com/SKF/proto/v2/iot"
)

type Device struct {
	ID                string `json:"id"`
	Manufacturer      string `json:"manufacturer"`
	Model             string `json:"model"`
	SerialNumber      string `json:"serialNumber"`
	DeviceCertificate string `json:"deviceCertificate,omitempty"`
	DevicePrivateKey  string `json:"devicePrivateKey,omitempty"`
}

type NodeDataAndPair struct {
	NodeData iotgrpcapi.NodeData `json:"nodeData"`
	NodeId   string              `json:"nodeId"`
}
