package gobcs_test

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/initia-labs/go-bcs/bcs"
	"github.com/stretchr/testify/require"
)

// MsgExecute is the message to execute the given module function
type MsgExecute struct {
	// Type is used to specify the message type of cosmos proto interface
	Type string `protobuf:"bytes,1,opt,name=sender,proto3" json:"@type,omitempty"`
	// Sender is the that actor that signed the messages
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// ModuleAddr is the address of the module deployer
	ModuleAddress string `protobuf:"bytes,2,opt,name=module_address,json=moduleAddress,proto3" json:"module_address,omitempty"`
	// ModuleName is the name of module to execute
	ModuleName string `protobuf:"bytes,3,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty"`
	// FunctionName is the name of a function to execute
	FunctionName string `protobuf:"bytes,4,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
	// TypeArgs is the type arguments of a function to execute
	// ex) "0x1::BasicCoin::Initia", "bool", "u8", "u64"
	TypeArgs []string `protobuf:"bytes,5,rep,name=type_args,json=typeArgs,proto3" json:"type_args,omitempty"`
	// Args is the arguments of a function to execute
	// - number: little endian
	// - string: base64 bytes
	Args [][]byte `protobuf:"bytes,6,rep,name=args,proto3" json:"args,omitempty"`
}

func Test_E2E(t *testing.T) {

	pairAddrBz, err := hex.DecodeString("a2b0d3c8e53e379ede31f3a361ff02716d50ec53c6b65b8c48a81d5b06548200")
	require.NoError(t, err)

	arg0, err := bcs.Marshal(bcs.NewAddressFromBytes(pairAddrBz))
	require.NoError(t, err)

	offerAddrBz, err := hex.DecodeString("29824d952e035490fae7567deea5f15b504a68fa73610063c160ab1fa87dd609")
	require.NoError(t, err)

	arg1, err := bcs.Marshal(bcs.NewAddressFromBytes(offerAddrBz))
	require.NoError(t, err)

	arg2, err := bcs.Marshal(uint64(100))
	require.NoError(t, err)

	arg3, err := bcs.Marshal(bcs.Some(uint64(10)))
	require.NoError(t, err)

	msg := MsgExecute{
		Type:          "/initia.move.v1.MsgExecute",
		Sender:        "init1h8lpl5qcs5k5nngxvdum5v20jnww2lckg3n2ta",
		ModuleAddress: "0x1",
		ModuleName:    "dex",
		FunctionName:  "swap_script",
		TypeArgs:      []string{},
		Args:          [][]byte{arg0, arg1, arg2, arg3},
	}

	bz, err := json.MarshalIndent(msg, "", "\t")
	require.NoError(t, err)
	require.Equal(t, `{
	"@type": "/initia.move.v1.MsgExecute",
	"sender": "init1h8lpl5qcs5k5nngxvdum5v20jnww2lckg3n2ta",
	"module_address": "0x1",
	"module_name": "dex",
	"function_name": "swap_script",
	"args": [
		"orDTyOU+N57eMfOjYf8CcW1Q7FPGtluMSKgdWwZUggA=",
		"KYJNlS4DVJD651Z97qXxW1BKaPpzYQBjwWCrH6h91gk=",
		"ZAAAAAAAAAA=",
		"AQoAAAAAAAAA"
	]
}`, fmt.Sprint(string(bz)))
}
