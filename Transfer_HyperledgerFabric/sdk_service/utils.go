package sdk_service

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"time"
)

type ServiceSetup struct {
	ChaincodeID string
	ChaincodeID2 string
	Client       *channel.Client
}

type ListResponse struct {
	Resultslist []Plain_equipment `json:"resultslist"`
	Total       int                `json:"total"`
}

type PlanListResponse struct {
	Resultslist []Cipher_equipment `json:"resultslist"`
	Total       int               `json:"total"`
}



type HistoryListResponse struct {
	HistoryList []HistoryList `json:"historylist"`
	Total       int           `json:"total"`
}

type HistoryList struct {
	ContractHash       string           `json:"ContractHash"`
	FunctionPath       string           `json:"FunctionPath"`
	Timestamp          string           `json:"Timestamp"`
	OperatingContracts string           `json:"OperatingContracts"`
	Value              Plain_equipment `json:"Value"`
}

func regitserEvent(client *channel.Client, ccID, eventID string) (fab.Registration, <-chan *fab.CCEvent) {
	reg, notifier, err := client.RegisterChaincodeEvent(ccID, eventID)
	if err != nil {
		fmt.Println("注册链码事件失败: ", err)
	}
	return reg, notifier
}

func eventResult(notifier <-chan *fab.CCEvent, eventID string) error {
	select {
	case ccEvent := <-notifier:
		fmt.Printf("接收到链码事件: %v\n", ccEvent)
	case <-time.After(time.Second * 20):
		return fmt.Errorf("不能根据指定的事件ID接收到相应的链码事件(%s)", eventID)
	}
	return nil
}
