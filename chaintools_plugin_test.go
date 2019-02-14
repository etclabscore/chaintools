package main

import (
	"plugin"
	"testing"
)

func TestPluginProduceTestChainFromGenesisFile(t *testing.T) {
	p, err := plugin.Open("./chaintools.so")
	if err != nil {
		t.Fatal(err)
	}
	sym, err := p.Lookup("ProduceTestChainFromGenesisFile")
	if err != nil {
		t.Fatal(err)
	}
	err = sym.(func(sourceGenesis, outputPath string, blockCount, blockTimeInSeconds uint) error)("./examples/input/genesis.json", "./examples/output/chain.rlp", 3, 30)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPluginProduceSimpleTestChain(t *testing.T) {
	p, err := plugin.Open("./chaintools.so")
	if err != nil {
		t.Fatal(err)
	}
	sym, err := p.Lookup("ProduceSimpleTestChain")
	if err != nil {
		t.Fatal(err)
	}
	err = sym.(func(path string, blockCount uint) error)("./examples/output/", 3)
	if err != nil {
		t.Fatal(err)
	}
}
