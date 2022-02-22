package main

import (
	"testing"

	clientMSP "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperbench/hyperbench-common/base"
	fcom "github.com/hyperbench/hyperbench-common/common"
	"github.com/stretchr/testify/assert"
)

func TestClientManager(t *testing.T) {
	s := &SDK{
		ConfigFileName: "config.yaml",
		SDK:            &fabsdk.FabricSDK{},
		OrgName:        "org1",
		OrgAdmin:       "Admin",
		MspIds:         []string{"Org1MSP"},
		EndPoints:      nil,
		MSPClient:      &clientMSP.Client{},
	}
	cm, err := NewClientManager(s, false, fcom.GetLogger("client"))
	assert.NotNil(t, cm)
	assert.NoError(t, err)

	err = cm.InitAccount(1)
	assert.NoError(t, err)

	err = cm.enroll(&Client{})
	assert.Error(t, err)

	cm, err = NewClientManager(s, true, fcom.GetLogger("client"))
	assert.NotNil(t, cm)
	assert.NoError(t, err)
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	err = cm.InitAccount(1)
	assert.NoError(t, err)

}

func TestGetMspClient(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	op := make(map[string]interface{})
	op["channel"] = "mychannel"
	op["MSP"] = true
	op["instant"] = 2
	b := base.NewBlockchainBase(base.ClientConfig{
		ClientType:   "fabric",
		ConfigPath:   "./../../../benchmark/fabricExample/fabric",
		ContractPath: "github.com/hyperbench/hyperbench/benchmark/fabricExample/contract",
		Args:         []interface{}{"init", "A", "123", "B", "234"},
		Options:      op,
	})
	c, err := New(b)
	assert.NotNil(t, c)
	assert.NoError(t, err)

	//deploy
	/* err = client.DeployContract()
	assert.NoError(t, err) */

	//getContext
	client := c.(*Fabric)
	context, err := client.GetContext()
	assert.NoError(t, err)
	assert.NotNil(t, context)

	//setContext
	err = client.SetContext(context)
	assert.NoError(t, err)
}
