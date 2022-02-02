package contracts

import (
	"testing"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/core"
	"math/big"
)

// Test greeter contract gets deployed correctly
func TestDeployGreeter(t *testing.T) {
	// Setup simulated block chain
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	blockchain := backends.NewSimulatedBackend(alloc, 10000000)

	// Deploy contract
	address, _, _, err := DeployGreeter(
		auth,
		blockchain,
		"Hello World",
	)

	// Commit all pending transactions
	blockchain.Commit()
	
	if err != nil {
		t.Fatalf("Failed to deploy the Greeter contract: %v", err)
	}

	if len(address.Bytes()) == 0 {
		t.Error("Expected a valid deployment address. Received empty address byte array instead")
	}
}