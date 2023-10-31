package main

import (
	"context"
	"math/big"
	"time"

	"github.com/0xPolygonHermez/zkevm-node/log"
	"github.com/0xPolygonHermez/zkevm-node/test/contracts/bin/ZkevmHashTest"
	"github.com/0xPolygonHermez/zkevm-node/test/operations"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	txTimeout = 60 * time.Second
)

func main() {
	var networks = []struct {
		Name       string
		URL        string
		ChainID    uint64
		PrivateKey string
	}{
		// {Name: "Local L1", URL: operations.DefaultL1NetworkURL, ChainID: operations.DefaultL1ChainID, PrivateKey: operations.DefaultSequencerPrivateKey},
		{Name: "Local L2", URL: operations.DefaultL2NetworkURL, ChainID: operations.DefaultL2ChainID, PrivateKey: operations.DefaultSequencerPrivateKey},
	}

	for _, network := range networks {
		ctx := context.Background()

		log.Infof("connecting to %v: %v", network.Name, network.URL)
		client, err := ethclient.Dial(network.URL)
		chkErr(err)
		log.Infof("connected")

		auth := operations.MustGetAuth(network.PrivateKey, network.ChainID)
		chkErr(err)
		
		log.Debugf("Sending TX to deploy ZkevmHashTest SC")
		_, tx, zkevmHashSC, err := ZkevmHashTest.DeployZkevmHashTest(auth, client)
		chkErr(err)
		err = operations.WaitTxToBeMined(ctx, client, tx, txTimeout)
		chkErr(err)
		hashCounts := big.NewInt(10)
		log.Debugf("Sending Tx to test keccak256")
		for i := 0; i < 100; i++ {
			tx, err = zkevmHashSC.TestKeccak256Hash(auth, hashCounts)
			chkErr(err)
		}

		log.Debugf("wait for 20s to closing previous batch...")
		time.Sleep(20*time.Second)
		
		log.Debugf("Sending Tx to test sha256")
		for i := 0; i < 100; i++ {
			tx, err = zkevmHashSC.TestSha256Hash(auth, hashCounts)
			chkErr(err)
		}
	}
}


func chkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
