package main

import (
	"testing"
)

func TestProduceTestChainFromGenesisFile(t *testing.T) {
	err := ProduceTestChainFromGenesisFile("./examples/input/genesis.json", "./examples/output/chain.rlp", 3, 30)
	if err != nil {
		t.Fatal(err)
	}
}

func TestProduceSimpleTestChain(t *testing.T) {
	err := ProduceSimpleTestChain("./examples/output/", 3)
	if err != nil {
		t.Fatal(err)
	}
}
