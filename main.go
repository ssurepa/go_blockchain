package main

import (
	"github.com/ssurepa/go_blockchain/cli"
	"github.com/ssurepa/go_blockchain/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
