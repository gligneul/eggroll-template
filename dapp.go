// Copyright (c) Gabriel de Quadros Ligneul
// SPDX-License-Identifier: MIT (see LICENSE)

package main

//go:generate go run github.com/gligneul/eggroll/cmd/eggroll schema gen

import (
	"github.com/gligneul/eggroll/pkg/eggroll"
)

type Contract struct {
}

func (c *Contract) AdvanceEcho(env eggroll.Env, value string) error {
	env.Report(EncodeEchoResponse(value))

	for kbSize := 100; kbSize < 4000; kbSize += 100 {
		bSize := kbSize << 10
		bytes := make([]byte, bSize)
		for i := range bytes {
			bytes[i] = 0xfa
		}
		env.Logf("sending voucher with %v kb", kbSize)
		env.Voucher(env.Sender(), bytes)
	}

	return nil
}

func (c *Contract) InspectEcho(env eggroll.EnvReader, value string) error {
	env.Report(EncodeEchoResponse(value))
	return nil
}

func main() {
	Roll(&Contract{})
}
