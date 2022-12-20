package config

import (
	"math/big"

	"github.com/scroll-tech/go-ethereum/common"
)

// L1Config loads l1eth configuration items.
type L1Config struct {
	// Confirmations block height confirmations number.
	Confirmations uint64 `json:"confirmations"`
	// l1 eth node url.
	Endpoint string `json:"endpoint"`
	// The start height to sync event from layer 1
	StartHeight *big.Int `json:"start_height"`
	// The messenger contract address deployed on layer 1 chain.
	L1MessengerAddress common.Address `json:"l1_messenger_address,omitempty"`
	// The relayer config
	RelayerConfig *RelayerConfig `json:"relayer_config"`
}