// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bank

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BankMetaData contains all meta data concerning the Bank contract.
var BankMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"API\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AccountBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"}],\"name\":\"SetContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5033600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611843806100616000396000f3fe60806040526004361061007b5760003560e01c8063bb62860d1161004e578063bb62860d1461010b578063d2aadb3c14610136578063e63f341f1461015f578063ed21248c1461019c5761007b565b80630ef678871461008057806357ea89b6146100ab5780637d7b0099146100b5578063b4a99a4e146100e0575b600080fd5b34801561008c57600080fd5b506100956101a6565b6040516100a29190610b9a565b60405180910390f35b6100b36101ed565b005b3480156100c157600080fd5b506100ca61035c565b6040516100d79190610bf6565b60405180910390f35b3480156100ec57600080fd5b506100f5610380565b6040516101029190610bf6565b60405180910390f35b34801561011757600080fd5b506101206103a6565b60405161012d9190610ca1565b60405180910390f35b34801561014257600080fd5b5061015d60048036038101906101589190610d03565b610434565b005b34801561016b57600080fd5b5061018660048036038101906101819190610d03565b6106e3565b6040516101939190610b9a565b60405180910390f35b6101a4610786565b005b6000600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527f57ea89b6000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516102b79190610d77565b600060405180830381855af49150503d80600081146102f2576040519150601f19603f3d011682016040523d82523d6000602084013e6102f7565b606091505b505090507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a610325826108f5565b6040516020016103359190610e16565b6040516020818303038152906040526040516103519190610ca1565b60405180910390a150565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600180546103b390610e7a565b80601f01602080910402602001604051908101604052809291908181526020018280546103df90610e7a565b801561042c5780601f106104015761010080835404028352916020019161042c565b820191906000526020600020905b81548152906001019060200180831161040f57829003601f168201915b505050505081565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461048e57600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527fbb62860d000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516105989190610d77565b6000604051808303816000865af19150503d80600081146105d5576040519150601f19603f3d011682016040523d82523d6000602084013e6105da565b606091505b5091509150811561060d57808060200190518101906105f99190610fd1565b6001908161060791906111c6565b50610653565b6040518060400160405280600781526020017f756e6b6e6f776e000000000000000000000000000000000000000000000000008152506001908161065191906111c6565b505b7fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61069d60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610978565b6106a6846108f5565b60016040516020016106ba9392919061138d565b6040516020818303038152906040526040516106d69190610ca1565b60405180910390a1505050565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461073f57600080fd5b600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527fed21248c000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516108509190610d77565b600060405180830381855af49150503d806000811461088b576040519150601f19603f3d011682016040523d82523d6000602084013e610890565b606091505b505090507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6108be826108f5565b6040516020016108ce9190610e16565b6040516020818303038152906040526040516108ea9190610ca1565b60405180910390a150565b6060811561093a576040518060400160405280600481526020017f74727565000000000000000000000000000000000000000000000000000000008152509050610973565b6040518060400160405280600581526020017f66616c736500000000000000000000000000000000000000000000000000000081525090505b919050565b60606000602867ffffffffffffffff81111561099757610996610eb5565b5b6040519080825280601f01601f1916602001820160405280156109c95781602001600182028036833780820191505090505b50905060005b6014811015610b315760008160136109e79190611429565b60086109f3919061145d565b60026109ff91906115d2565b8573ffffffffffffffffffffffffffffffffffffffff16610a20919061164c565b60f81b9050600060108260f81c610a37919061168a565b60f81b905060008160f81c6010610a4e91906116bb565b8360f81c610a5c91906116f8565b60f81b9050610a6a82610b3b565b85856002610a78919061145d565b81518110610a8957610a8861172d565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610ac181610b3b565b856001866002610ad1919061145d565b610adb919061175c565b81518110610aec57610aeb61172d565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505050508080610b2990611790565b9150506109cf565b5080915050919050565b6000600a8260f81c60ff161015610b665760308260f81c610b5c91906117d8565b60f81b9050610b7c565b60578260f81c610b7691906117d8565b60f81b90505b919050565b6000819050919050565b610b9481610b81565b82525050565b6000602082019050610baf6000830184610b8b565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610be082610bb5565b9050919050565b610bf081610bd5565b82525050565b6000602082019050610c0b6000830184610be7565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610c4b578082015181840152602081019050610c30565b60008484015250505050565b6000601f19601f8301169050919050565b6000610c7382610c11565b610c7d8185610c1c565b9350610c8d818560208601610c2d565b610c9681610c57565b840191505092915050565b60006020820190508181036000830152610cbb8184610c68565b905092915050565b6000604051905090565b600080fd5b600080fd5b610ce081610bd5565b8114610ceb57600080fd5b50565b600081359050610cfd81610cd7565b92915050565b600060208284031215610d1957610d18610ccd565b5b6000610d2784828501610cee565b91505092915050565b600081519050919050565b600081905092915050565b6000610d5182610d30565b610d5b8185610d3b565b9350610d6b818560208601610c2d565b80840191505092915050565b6000610d838284610d46565b915081905092915050565b7f737563636573735b000000000000000000000000000000000000000000000000815250565b600081905092915050565b6000610dca82610c11565b610dd48185610db4565b9350610de4818560208601610c2d565b80840191505092915050565b7f5d00000000000000000000000000000000000000000000000000000000000000815250565b6000610e2182610d8e565b600882019150610e318284610dbf565b9150610e3c82610df0565b60018201915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610e9257607f821691505b602082108103610ea557610ea4610e4b565b5b50919050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610eed82610c57565b810181811067ffffffffffffffff82111715610f0c57610f0b610eb5565b5b80604052505050565b6000610f1f610cc3565b9050610f2b8282610ee4565b919050565b600067ffffffffffffffff821115610f4b57610f4a610eb5565b5b610f5482610c57565b9050602081019050919050565b6000610f74610f6f84610f30565b610f15565b905082815260208101848484011115610f9057610f8f610eb0565b5b610f9b848285610c2d565b509392505050565b600082601f830112610fb857610fb7610eab565b5b8151610fc8848260208601610f61565b91505092915050565b600060208284031215610fe757610fe6610ccd565b5b600082015167ffffffffffffffff81111561100557611004610cd2565b5b61101184828501610fa3565b91505092915050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b60006008830261107c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261103f565b611086868361103f565b95508019841693508086168417925050509392505050565b6000819050919050565b60006110c36110be6110b984610b81565b61109e565b610b81565b9050919050565b6000819050919050565b6110dd836110a8565b6110f16110e9826110ca565b84845461104c565b825550505050565b600090565b6111066110f9565b6111118184846110d4565b505050565b5b818110156111355761112a6000826110fe565b600181019050611117565b5050565b601f82111561117a5761114b8161101a565b6111548461102f565b81016020851015611163578190505b61117761116f8561102f565b830182611116565b50505b505050565b600082821c905092915050565b600061119d6000198460080261117f565b1980831691505092915050565b60006111b6838361118c565b9150826002028217905092915050565b6111cf82610c11565b67ffffffffffffffff8111156111e8576111e7610eb5565b5b6111f28254610e7a565b6111fd828285611139565b600060209050601f831160018114611230576000841561121e578287015190505b61122885826111aa565b865550611290565b601f19841661123e8661101a565b60005b8281101561126657848901518255600182019150602085019450602081019050611241565b86831015611283578489015161127f601f89168261118c565b8355505b6001600288020188555050505b505050505050565b7f636f6e74726163745b0000000000000000000000000000000000000000000000815250565b7f5d20737563636573735b00000000000000000000000000000000000000000000815250565b7f5d2076657273696f6e5b00000000000000000000000000000000000000000000815250565b6000815461131781610e7a565b6113218186610db4565b9450600182166000811461133c576001811461135157611384565b60ff1983168652811515820286019350611384565b61135a8561101a565b60005b8381101561137c5781548189015260018201915060208101905061135d565b838801955050505b50505092915050565b600061139882611298565b6009820191506113a88286610dbf565b91506113b3826112be565b600a820191506113c38285610dbf565b91506113ce826112e4565b600a820191506113de828461130a565b91506113e982610df0565b600182019150819050949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061143482610b81565b915061143f83610b81565b9250828203905081811115611457576114566113fa565b5b92915050565b600061146882610b81565b915061147383610b81565b925082820261148181610b81565b91508282048414831517611498576114976113fa565b5b5092915050565b60008160011c9050919050565b6000808291508390505b60018511156114f6578086048111156114d2576114d16113fa565b5b60018516156114e15780820291505b80810290506114ef8561149f565b94506114b6565b94509492505050565b60008261150f57600190506115cb565b8161151d57600090506115cb565b8160018114611533576002811461153d5761156c565b60019150506115cb565b60ff84111561154f5761154e6113fa565b5b8360020a915084821115611566576115656113fa565b5b506115cb565b5060208310610133831016604e8410600b84101617156115a15782820a90508381111561159c5761159b6113fa565b5b6115cb565b6115ae84848460016114ac565b925090508184048111156115c5576115c46113fa565b5b81810290505b9392505050565b60006115dd82610b81565b91506115e883610b81565b92506116157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846114ff565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061165782610b81565b915061166283610b81565b9250826116725761167161161d565b5b828204905092915050565b600060ff82169050919050565b60006116958261167d565b91506116a08361167d565b9250826116b0576116af61161d565b5b828204905092915050565b60006116c68261167d565b91506116d18361167d565b92508282026116df8161167d565b91508082146116f1576116f06113fa565b5b5092915050565b60006117038261167d565b915061170e8361167d565b9250828203905060ff811115611727576117266113fa565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600061176782610b81565b915061177283610b81565b925082820190508082111561178a576117896113fa565b5b92915050565b600061179b82610b81565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036117cd576117cc6113fa565b5b600182019050919050565b60006117e38261167d565b91506117ee8361167d565b9250828201905060ff811115611807576118066113fa565b5b9291505056fea2646970667358221220efff196a730424f4d5b7e73fde2fbad2a10af7e35b5619819a52e7a3d2fb203764736f6c63430008120033",
}

// BankABI is the input ABI used to generate the binding from.
// Deprecated: Use BankMetaData.ABI instead.
var BankABI = BankMetaData.ABI

// BankBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BankMetaData.Bin instead.
var BankBin = BankMetaData.Bin

// DeployBank deploys a new Ethereum contract, binding an instance of Bank to it.
func DeployBank(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bank, error) {
	parsed, err := BankMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BankBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bank{BankCaller: BankCaller{contract: contract}, BankTransactor: BankTransactor{contract: contract}, BankFilterer: BankFilterer{contract: contract}}, nil
}

// Bank is an auto generated Go binding around an Ethereum contract.
type Bank struct {
	BankCaller     // Read-only binding to the contract
	BankTransactor // Write-only binding to the contract
	BankFilterer   // Log filterer for contract events
}

// BankCaller is an auto generated read-only Go binding around an Ethereum contract.
type BankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BankSession struct {
	Contract     *Bank             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BankCallerSession struct {
	Contract *BankCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BankTransactorSession struct {
	Contract     *BankTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankRaw is an auto generated low-level Go binding around an Ethereum contract.
type BankRaw struct {
	Contract *Bank // Generic contract binding to access the raw methods on
}

// BankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BankCallerRaw struct {
	Contract *BankCaller // Generic read-only contract binding to access the raw methods on
}

// BankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BankTransactorRaw struct {
	Contract *BankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBank creates a new instance of Bank, bound to a specific deployed contract.
func NewBank(address common.Address, backend bind.ContractBackend) (*Bank, error) {
	contract, err := bindBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bank{BankCaller: BankCaller{contract: contract}, BankTransactor: BankTransactor{contract: contract}, BankFilterer: BankFilterer{contract: contract}}, nil
}

// NewBankCaller creates a new read-only instance of Bank, bound to a specific deployed contract.
func NewBankCaller(address common.Address, caller bind.ContractCaller) (*BankCaller, error) {
	contract, err := bindBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BankCaller{contract: contract}, nil
}

// NewBankTransactor creates a new write-only instance of Bank, bound to a specific deployed contract.
func NewBankTransactor(address common.Address, transactor bind.ContractTransactor) (*BankTransactor, error) {
	contract, err := bindBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BankTransactor{contract: contract}, nil
}

// NewBankFilterer creates a new log filterer instance of Bank, bound to a specific deployed contract.
func NewBankFilterer(address common.Address, filterer bind.ContractFilterer) (*BankFilterer, error) {
	contract, err := bindBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BankFilterer{contract: contract}, nil
}

// bindBank binds a generic wrapper to an already deployed contract.
func bindBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BankABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bank *BankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bank.Contract.BankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bank *BankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.Contract.BankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bank *BankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bank.Contract.BankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bank *BankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bank *BankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bank *BankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bank.Contract.contract.Transact(opts, method, params...)
}

// API is a free data retrieval call binding the contract method 0x7d7b0099.
//
// Solidity: function API() view returns(address)
func (_Bank *BankCaller) API(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "API")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// API is a free data retrieval call binding the contract method 0x7d7b0099.
//
// Solidity: function API() view returns(address)
func (_Bank *BankSession) API() (common.Address, error) {
	return _Bank.Contract.API(&_Bank.CallOpts)
}

// API is a free data retrieval call binding the contract method 0x7d7b0099.
//
// Solidity: function API() view returns(address)
func (_Bank *BankCallerSession) API() (common.Address, error) {
	return _Bank.Contract.API(&_Bank.CallOpts)
}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256)
func (_Bank *BankCaller) AccountBalance(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "AccountBalance", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256)
func (_Bank *BankSession) AccountBalance(account common.Address) (*big.Int, error) {
	return _Bank.Contract.AccountBalance(&_Bank.CallOpts, account)
}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256)
func (_Bank *BankCallerSession) AccountBalance(account common.Address) (*big.Int, error) {
	return _Bank.Contract.AccountBalance(&_Bank.CallOpts, account)
}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Bank *BankCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "Balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Bank *BankSession) Balance() (*big.Int, error) {
	return _Bank.Contract.Balance(&_Bank.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Bank *BankCallerSession) Balance() (*big.Int, error) {
	return _Bank.Contract.Balance(&_Bank.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bank *BankCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bank *BankSession) Owner() (common.Address, error) {
	return _Bank.Contract.Owner(&_Bank.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bank *BankCallerSession) Owner() (common.Address, error) {
	return _Bank.Contract.Owner(&_Bank.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0xbb62860d.
//
// Solidity: function Version() view returns(string)
func (_Bank *BankCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "Version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0xbb62860d.
//
// Solidity: function Version() view returns(string)
func (_Bank *BankSession) Version() (string, error) {
	return _Bank.Contract.Version(&_Bank.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0xbb62860d.
//
// Solidity: function Version() view returns(string)
func (_Bank *BankCallerSession) Version() (string, error) {
	return _Bank.Contract.Version(&_Bank.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bank *BankTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "Deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bank *BankSession) Deposit() (*types.Transaction, error) {
	return _Bank.Contract.Deposit(&_Bank.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bank *BankTransactorSession) Deposit() (*types.Transaction, error) {
	return _Bank.Contract.Deposit(&_Bank.TransactOpts)
}

// SetContract is a paid mutator transaction binding the contract method 0xd2aadb3c.
//
// Solidity: function SetContract(address contractAddr) returns()
func (_Bank *BankTransactor) SetContract(opts *bind.TransactOpts, contractAddr common.Address) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "SetContract", contractAddr)
}

// SetContract is a paid mutator transaction binding the contract method 0xd2aadb3c.
//
// Solidity: function SetContract(address contractAddr) returns()
func (_Bank *BankSession) SetContract(contractAddr common.Address) (*types.Transaction, error) {
	return _Bank.Contract.SetContract(&_Bank.TransactOpts, contractAddr)
}

// SetContract is a paid mutator transaction binding the contract method 0xd2aadb3c.
//
// Solidity: function SetContract(address contractAddr) returns()
func (_Bank *BankTransactorSession) SetContract(contractAddr common.Address) (*types.Transaction, error) {
	return _Bank.Contract.SetContract(&_Bank.TransactOpts, contractAddr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bank *BankTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "Withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bank *BankSession) Withdraw() (*types.Transaction, error) {
	return _Bank.Contract.Withdraw(&_Bank.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bank *BankTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Bank.Contract.Withdraw(&_Bank.TransactOpts)
}

// BankEventLogIterator is returned from FilterEventLog and is used to iterate over the raw logs and unpacked data for EventLog events raised by the Bank contract.
type BankEventLogIterator struct {
	Event *BankEventLog // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BankEventLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankEventLog)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BankEventLog)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BankEventLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankEventLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankEventLog represents a EventLog event raised by the Bank contract.
type BankEventLog struct {
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEventLog is a free log retrieval operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Bank *BankFilterer) FilterEventLog(opts *bind.FilterOpts) (*BankEventLogIterator, error) {

	logs, sub, err := _Bank.contract.FilterLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return &BankEventLogIterator{contract: _Bank.contract, event: "EventLog", logs: logs, sub: sub}, nil
}

// WatchEventLog is a free log subscription operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Bank *BankFilterer) WatchEventLog(opts *bind.WatchOpts, sink chan<- *BankEventLog) (event.Subscription, error) {

	logs, sub, err := _Bank.contract.WatchLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankEventLog)
				if err := _Bank.contract.UnpackLog(event, "EventLog", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEventLog is a log parse operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Bank *BankFilterer) ParseEventLog(log types.Log) (*BankEventLog, error) {
	event := new(BankEventLog)
	if err := _Bank.contract.UnpackLog(event, "EventLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
