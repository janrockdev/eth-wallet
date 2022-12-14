package tests

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/janrockdev/eth-wallet/types"
)

var (
	TestNetwork     = types.RopstenNet
	TestTransaction = "0xbb3a336e3f823ec18197f1e13ee875700f08f03e2cab75f0d0b118dabb44cba0"
	TestKeyAuth     = struct {
		path string
		auth string
	}{
		path: "./key/test",
		auth: "1234",
	}
	TestWalletAuth = struct {
		path string
		auth string
	}{
		path: "./key/test",
		auth: "1234",
	}

	TestAddress = common.Address{
		0x51, 0xbf, 0x0b, 0x41, 0xBa, 0x5B, 0x03, 0x4f,
		0x15, 0x8C, 0xF1, 0x23, 0x3f, 0x16, 0xbA, 0x54,
		0x50, 0xF9, 0x35, 0x5B,
	}

	TestContractAddress = common.Address{
		0x10, 0x18, 0x48, 0xd5, 0xc5, 0xbb, 0xca, 0x18,
		0xE6, 0xb4, 0x43, 0x1e, 0xEd, 0xF6, 0xB9, 0x5E,
		0x9A, 0xDF, 0x82, 0xFA,
	}

	TestTransactionRequest = types.TransactionRequest{
		From: "0xb60e8dd61c5d32be8058bb8eb970870f07233155",
		To:   "0xd46e8dd67c5d32be8058bb8eb970870f07244567",
	}
	TestTransactionEstimateGas = []struct {
		transaction types.TransactionRequest
		gas         string
	}{
		{
			types.TransactionRequest{
				From:     "0x51bf0b41Ba5B034f158CF1233f16bA5450F9355B",
				To:       "0x101848D5C5bBca18E6b4431eEdF6B95E9ADF82FA",
				Gas:      "0x186a0",
				GasPrice: "0x3b9aca00",
				Value:    "0x0",
				Data:     "0xa9059cbb00000000000000000000000051bf0b41ba5b034f158cf1233f16ba5450f9355b0000000000000000000000000000000000000000000000000de0b6b3a7640000",
			},
			"0x7c7d",
		}, {
			types.TransactionRequest{
				From:     "0xb60e8dd61c5d32be8058bb8eb970870f07233155",
				To:       "0xe5664B93Ad268393d1F695c4180993E60c59FC3E",
				GasPrice: "0x3b9aca",
				Value:    "0x5af3107a4000",
			},
			"0x5208",
		},
	}

	TestTranasctionCall = []struct {
		transaction types.TransactionRequest
		res         string
	}{
		{
			types.TransactionRequest{
				From:     "0xe89c185211cE311DEc848BF1EFeFab437f3c7e42",
				To:       "0x101848D5C5bBca18E6b4431eEdF6B95E9ADF82FA",
				Gas:      "0x76c0",
				GasPrice: "0x3b9aca00",
				Value:    "0x0",
				Data:     "0xa9059cbb00000000000000000000000051bf0b41ba5b034f158cf1233f16ba5450f9355b0000000000000000000000000000000000000000000000000de0b6b3a7640000",
			},
			"0x",
		},
		{
			types.TransactionRequest{
				From:     "0xb60e8dd61c5d32be8058bb8eb970870f07233155",
				To:       "0xd46e8dd67c5d32be8058bb8eb970870f07244567",
				Gas:      "0x76c0",
				GasPrice: "0x9184e72a000",
				Value:    "0x9184e72a",
				Data:     "0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675",
			},
			"0x",
		},
	}

	TestErc20Call = []struct {
		addr *common.Address
		ans  types.Erc20Token
	}{
		{
			&common.Address{
				0xdd, 0x97, 0x4D, 0x5C, 0x2e, 0x29, 0x28, 0xde,
				0xA5, 0xF7, 0x1b, 0x98, 0x25, 0xb8, 0xb6, 0x46,
				0x68, 0x6B, 0xD2, 0x00,
			}, types.Erc20Token{
				Decimals: 18,
				Name:     "Kyber Network Crystal",
				Symbol:   "KNC",
				Address: &common.Address{
					0xdd, 0x97, 0x4D, 0x5C, 0x2e, 0x29, 0x28, 0xde,
					0xA5, 0xF7, 0x1b, 0x98, 0x25, 0xb8, 0xb6, 0x46,
					0x68, 0x6B, 0xD2, 0x00,
				},
			},
		},
	}

	testRlpArray = []struct {
		arr [][]byte
		res []byte
	}{
		{
			[][]byte{
				{
					0x8b,
				},
				{
					0x02, 0x54, 0x0b, 0xe4, 0x00,
				},
				{
					0x06, 0x1a, 0x80,
				},
				{
					0x06, 0x01, 0x2c, 0x8c, 0xf9, 0x7b, 0xea, 0xd5,
					0xde, 0xae, 0x23, 0x70, 0x70, 0xf9, 0x58, 0x7f,
					0x8e, 0x7a, 0x26, 0x6d,
				},
				{},
				{
					0x23, 0xb8, 0x72, 0xdd, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0xe5, 0x85, 0x9f, 0x51, 0xef, 0x87, 0xc9, 0xda,
					0x46, 0x58, 0xde, 0x3f, 0x65, 0xe7, 0x12, 0x15,
					0x33, 0xec, 0xd3, 0xb7, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x0e, 0xd7, 0x56, 0xd9, 0x9d, 0x9e, 0x88, 0xd9,
					0x7a, 0x7b, 0x1a, 0x0c, 0xc8, 0x5d, 0xfd, 0x5f,
					0x06, 0x7d, 0xf0, 0x35, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x11, 0x5c, 0x81,
				},
				{
					0x25,
				},
				{
					0x16, 0xde, 0xbe, 0xe2, 0x60, 0xb4, 0x44, 0x04,
					0x45, 0xfc, 0x73, 0x1d, 0x6e, 0x70, 0x35, 0x41,
					0xa0, 0x8c, 0xb1, 0x0b, 0x6f, 0xe0, 0xd5, 0xb7,
					0xe8, 0x78, 0xfc, 0x68, 0x78, 0xea, 0x79, 0x78,
				},
				{
					0x4d, 0x21, 0x86, 0xfc, 0xe0, 0x5a, 0xa9, 0xa4,
					0x54, 0x06, 0x3c, 0x3e, 0xef, 0x93, 0x05, 0x98,
					0xd8, 0x17, 0xa4, 0x5e, 0x74, 0x87, 0xe4, 0x3c,
					0x6b, 0xce, 0xbe, 0x17, 0x5c, 0x9b, 0xa6, 0xc9,
				},
			},
			[]byte{0xf8, 0xcb, 0x81, 0x8b, 0x85, 0x02, 0x54, 0x0b,
				0xe4, 0x00, 0x83, 0x06, 0x1a, 0x80, 0x94, 0x06,
				0x01, 0x2c, 0x8c, 0xf9, 0x7b, 0xea, 0xd5, 0xde,
				0xae, 0x23, 0x70, 0x70, 0xf9, 0x58, 0x7f, 0x8e,
				0x7a, 0x26, 0x6d, 0x80, 0xb8, 0x64, 0x23, 0xb8,
				0x72, 0xdd, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xe5, 0x85,
				0x9f, 0x51, 0xef, 0x87, 0xc9, 0xda, 0x46, 0x58,
				0xde, 0x3f, 0x65, 0xe7, 0x12, 0x15, 0x33, 0xec,
				0xd3, 0xb7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0e, 0xd7,
				0x56, 0xd9, 0x9d, 0x9e, 0x88, 0xd9, 0x7a, 0x7b,
				0x1a, 0x0c, 0xc8, 0x5d, 0xfd, 0x5f, 0x06, 0x7d,
				0xf0, 0x35, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x11,
				0x5c, 0x81, 0x25, 0xa0, 0x16, 0xde, 0xbe, 0xe2,
				0x60, 0xb4, 0x44, 0x04, 0x45, 0xfc, 0x73, 0x1d,
				0x6e, 0x70, 0x35, 0x41, 0xa0, 0x8c, 0xb1, 0x0b,
				0x6f, 0xe0, 0xd5, 0xb7, 0xe8, 0x78, 0xfc, 0x68,
				0x78, 0xea, 0x79, 0x78, 0xa0, 0x4d, 0x21, 0x86,
				0xfc, 0xe0, 0x5a, 0xa9, 0xa4, 0x54, 0x06, 0x3c,
				0x3e, 0xef, 0x93, 0x05, 0x98, 0xd8, 0x17, 0xa4,
				0x5e, 0x74, 0x87, 0xe4, 0x3c, 0x6b, 0xce, 0xbe,
				0x17, 0x5c, 0x9b, 0xa6, 0xc9},
		}, {
			[][]byte{
				{
					'd', 'o', 'g',
				},
				{
					'c', 'a', 't',
				},
			}, []byte{
				0xc8, 0x83, 'd', 'o', 'g', 0x83, 'c', 'a', 't',
			},
		},
	}
)
