// Code generated by EggRoll - DO NOT EDIT.

package main

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gligneul/eggroll/pkg/eggroll"
	"github.com/gligneul/eggroll/pkg/eggtypes"
)

var (
	_ = big.NewInt
	_ = common.Big1
	_ = eggtypes.MustAddSchema
)

// Messages encoded as JSON ABI.
const _JSON_ABI = `[
  {
    "name": "echoResponse",
    "type": "function",
    "stateMutability": "nonpayable",
    "inputs": [
      {
        "name": "value",
        "type": "string",
        "internalType": "string",
        "components": null
      }
    ],
    "outputs": null
  },
  {
    "name": "advanceEcho",
    "type": "function",
    "stateMutability": "nonpayable",
    "inputs": [
      {
        "name": "value",
        "type": "string",
        "internalType": "string",
        "components": null
      }
    ],
    "outputs": null
  },
  {
    "name": "inspectEcho",
    "type": "function",
    "stateMutability": "nonpayable",
    "inputs": [
      {
        "name": "value",
        "type": "string",
        "internalType": "string",
        "components": null
      }
    ],
    "outputs": null
  }
]
`

// Solidity ABI.
var _abi abi.ABI

//
// Struct Types
//

type EchoResponse struct {
	Value string
}

type AdvanceEcho struct {
	Value string
}

type InspectEcho struct {
	Value string
}

//
// ID for each schema
//

// 4-byte function selector of echoResponse
var EchoResponseID eggtypes.ID

// 4-byte function selector of advanceEcho
var AdvanceEchoID eggtypes.ID

// 4-byte function selector of inspectEcho
var InspectEchoID eggtypes.ID

//
// Encode functions for each message schema
//

// Encode echoResponse into binary data.
func EncodeEchoResponse(
	Value string,
) []byte {
	values := make([]any, 1)
	values[0] = Value
	data, err := _abi.Methods["echoResponse"].Inputs.PackValues(values)
	if err != nil {
		panic(fmt.Sprintf("failed to encode echoResponse: %v", err))
	}
	return append(EchoResponseID[:], data...)
}

// Encode echoResponse into binary data.
func (v EchoResponse) Encode() []byte {
	return EncodeEchoResponse(
		v.Value,
	)
}

// Encode advanceEcho into binary data.
func EncodeAdvanceEcho(
	Value string,
) []byte {
	values := make([]any, 1)
	values[0] = Value
	data, err := _abi.Methods["advanceEcho"].Inputs.PackValues(values)
	if err != nil {
		panic(fmt.Sprintf("failed to encode advanceEcho: %v", err))
	}
	return append(AdvanceEchoID[:], data...)
}

// Encode advanceEcho into binary data.
func (v AdvanceEcho) Encode() []byte {
	return EncodeAdvanceEcho(
		v.Value,
	)
}

// Encode inspectEcho into binary data.
func EncodeInspectEcho(
	Value string,
) []byte {
	values := make([]any, 1)
	values[0] = Value
	data, err := _abi.Methods["inspectEcho"].Inputs.PackValues(values)
	if err != nil {
		panic(fmt.Sprintf("failed to encode inspectEcho: %v", err))
	}
	return append(InspectEchoID[:], data...)
}

// Encode inspectEcho into binary data.
func (v InspectEcho) Encode() []byte {
	return EncodeInspectEcho(
		v.Value,
	)
}

//
// Decode functions for each message schema
//

func _decode_EchoResponse(values []any) (any, error) {
	if len(values) != 1 {
		return nil, fmt.Errorf("wrong number of values")
	}
	var ok bool
	var v EchoResponse
	v.Value, ok = values[0].(string)
	if !ok {
		return nil, fmt.Errorf("failed to decode echoResponse.value")
	}
	return v, nil
}

func _decode_AdvanceEcho(values []any) (any, error) {
	if len(values) != 1 {
		return nil, fmt.Errorf("wrong number of values")
	}
	var ok bool
	var v AdvanceEcho
	v.Value, ok = values[0].(string)
	if !ok {
		return nil, fmt.Errorf("failed to decode advanceEcho.value")
	}
	return v, nil
}

func _decode_InspectEcho(values []any) (any, error) {
	if len(values) != 1 {
		return nil, fmt.Errorf("wrong number of values")
	}
	var ok bool
	var v InspectEcho
	v.Value, ok = values[0].(string)
	if !ok {
		return nil, fmt.Errorf("failed to decode inspectEcho.value")
	}
	return v, nil
}

//
// Init function
//

func init() {
	var err error
	_abi, err = abi.JSON(strings.NewReader(_JSON_ABI))
	if err != nil {
		// This should not happen
		panic(fmt.Sprintf("failed to decode ABI: %v", err))
	}
	EchoResponseID = eggtypes.ID(_abi.Methods["echoResponse"].ID)
	eggtypes.MustAddSchema(eggtypes.MessageSchema{
		ID:        EchoResponseID,
		Kind:      "echoResponse",
		Arguments: _abi.Methods["echoResponse"].Inputs,
		Decoder:   _decode_EchoResponse,
	})
	AdvanceEchoID = eggtypes.ID(_abi.Methods["advanceEcho"].ID)
	eggtypes.MustAddSchema(eggtypes.MessageSchema{
		ID:        AdvanceEchoID,
		Kind:      "advanceEcho",
		Arguments: _abi.Methods["advanceEcho"].Inputs,
		Decoder:   _decode_AdvanceEcho,
	})
	InspectEchoID = eggtypes.ID(_abi.Methods["inspectEcho"].ID)
	eggtypes.MustAddSchema(eggtypes.MessageSchema{
		ID:        InspectEchoID,
		Kind:      "inspectEcho",
		Arguments: _abi.Methods["inspectEcho"].Inputs,
		Decoder:   _decode_InspectEcho,
	})
}

//
// Middleware
//

// High-level contract
type iContract interface {
	AdvanceEcho(
		eggroll.Env,
		string,
	) error

	InspectEcho(
		eggroll.EnvReader,
		string,
	) error
}

// Middleware that implements the EggRoll Middleware interface.
// The middleware requires a high-level contract to work.
type Middleware struct {
	contract iContract
}

func (m Middleware) Advance(env eggroll.Env, input []byte) error {
	unpacked, err := eggtypes.Decode(input)
	if err != nil {
		return err
	}
	env.Logf("middleware: received %#v", unpacked)
	switch input := unpacked.(type) {
	case AdvanceEcho:
		return m.contract.AdvanceEcho(
			env,
			input.Value,
		)
	default:
		return fmt.Errorf("middleware: input isn't an advance")
	}
}

func (m Middleware) Inspect(env eggroll.EnvReader, input []byte) error {
	unpacked, err := eggtypes.Decode(input)
	if err != nil {
		return err
	}
	env.Logf("middleware: received %#v", unpacked)
	switch input := unpacked.(type) {
	case InspectEcho:
		return m.contract.InspectEcho(
			env,
			input.Value,
		)
	default:
		return fmt.Errorf("middleware: input isn't an inspect")
	}
}

// Call eggroll.Roll for the contract using the middleware wrapper.
func Roll(contract iContract) {
	eggroll.Roll(Middleware{contract})
}
