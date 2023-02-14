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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AccountBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040518060400160405280600581526020017f302e312e300000000000000000000000000000000000000000000000000000008152506001908162000098919062000319565b5062000400565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200012157607f821691505b602082108103620001375762000136620000d9565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620001a17fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000162565b620001ad868362000162565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620001fa620001f4620001ee84620001c5565b620001cf565b620001c5565b9050919050565b6000819050919050565b6200021683620001d9565b6200022e620002258262000201565b8484546200016f565b825550505050565b600090565b6200024562000236565b620002528184846200020b565b505050565b5b818110156200027a576200026e6000826200023b565b60018101905062000258565b5050565b601f821115620002c95762000293816200013d565b6200029e8462000152565b81016020851015620002ae578190505b620002c6620002bd8562000152565b83018262000257565b50505b505050565b600082821c905092915050565b6000620002ee60001984600802620002ce565b1980831691505092915050565b6000620003098383620002db565b9150826002028217905092915050565b62000324826200009f565b67ffffffffffffffff81111562000340576200033f620000aa565b5b6200034c825462000108565b620003598282856200027e565b600060209050601f8311600181146200039157600084156200037c578287015190505b620003888582620002fb565b865550620003f8565b601f198416620003a1866200013d565b60005b82811015620003cb57848901518255600182019150602085019450602081019050620003a4565b86831015620003eb5784890151620003e7601f891682620002db565b8355505b6001600288020188555050505b505050505050565b61119a80620004106000396000f3fe6080604052600436106100555760003560e01c80630ef678871461005a57806357ea89b614610085578063b4a99a4e1461008f578063bb62860d146100ba578063e63f341f146100e5578063ed21248c14610122575b600080fd5b34801561006657600080fd5b5061006f61012c565b60405161007c9190610933565b60405180910390f35b61008d610173565b005b34801561009b57600080fd5b506100a4610336565b6040516100b1919061098f565b60405180910390f35b3480156100c657600080fd5b506100cf61035a565b6040516100dc9190610a3a565b60405180910390f35b3480156100f157600080fd5b5061010c60048036038101906101079190610a8d565b6103e8565b6040516101199190610933565b60405180910390f35b61012a61048a565b005b6000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905090565b60003390506000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054036101fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101f190610b06565b60405180910390fd5b6000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015610284573d6000803e3d6000fd5b506000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6102f433610589565b6102fd8361074c565b60405160200161030e929190610bd4565b60405160208183030381529060405260405161032a9190610a3a565b60405180910390a15050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6001805461036790610c54565b80601f016020809104026020016040519081016040528092919081815260200182805461039390610c54565b80156103e05780601f106103b5576101008083540402835291602001916103e0565b820191906000526020600020905b8154815290600101906020018083116103c357829003601f168201915b505050505081565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461044357600080fd5b600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b34600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546104d99190610cb4565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61050a33610589565b610552600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461074c565b604051602001610563929190610d34565b60405160208183030381529060405260405161057f9190610a3a565b60405180910390a1565b60606000602867ffffffffffffffff8111156105a8576105a7610d85565b5b6040519080825280601f01601f1916602001820160405280156105da5781602001600182028036833780820191505090505b50905060005b60148110156107425760008160136105f89190610db4565b60086106049190610de8565b60026106109190610f5d565b8573ffffffffffffffffffffffffffffffffffffffff166106319190610fd7565b60f81b9050600060108260f81c6106489190611015565b60f81b905060008160f81c601061065f9190611046565b8360f81c61066d9190611083565b60f81b905061067b826108d4565b858560026106899190610de8565b8151811061069a576106996110b8565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506106d2816108d4565b8560018660026106e29190610de8565b6106ec9190610cb4565b815181106106fd576106fc6110b8565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350505050808061073a906110e7565b9150506105e0565b5080915050919050565b606060008203610793576040518060400160405280600181526020017f300000000000000000000000000000000000000000000000000000000000000081525090506108cf565b600082905060005b600082146107c55780806107ae906110e7565b915050600a826107be9190610fd7565b915061079b565b60008167ffffffffffffffff8111156107e1576107e0610d85565b5b6040519080825280601f01601f1916602001820160405280156108135781602001600182028036833780820191505090505b50905060008290505b600086146108c7576001816108319190610db4565b90506000600a80886108439190610fd7565b61084d9190610de8565b876108589190610db4565b6030610864919061112f565b905060008160f81b905080848481518110610882576108816110b8565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a886108be9190610fd7565b9750505061081c565b819450505050505b919050565b6000600a8260f81c60ff1610156108ff5760308260f81c6108f5919061112f565b60f81b9050610915565b60578260f81c61090f919061112f565b60f81b90505b919050565b6000819050919050565b61092d8161091a565b82525050565b60006020820190506109486000830184610924565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006109798261094e565b9050919050565b6109898161096e565b82525050565b60006020820190506109a46000830184610980565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156109e45780820151818401526020810190506109c9565b60008484015250505050565b6000601f19601f8301169050919050565b6000610a0c826109aa565b610a1681856109b5565b9350610a268185602086016109c6565b610a2f816109f0565b840191505092915050565b60006020820190508181036000830152610a548184610a01565b905092915050565b600080fd5b610a6a8161096e565b8114610a7557600080fd5b50565b600081359050610a8781610a61565b92915050565b600060208284031215610aa357610aa2610a5c565b5b6000610ab184828501610a78565b91505092915050565b7f6e6f7420656e6f7567682062616c616e63650000000000000000000000000000600082015250565b6000610af06012836109b5565b9150610afb82610aba565b602082019050919050565b60006020820190508181036000830152610b1f81610ae3565b9050919050565b7f77697468647261775b0000000000000000000000000000000000000000000000815250565b600081905092915050565b6000610b62826109aa565b610b6c8185610b4c565b9350610b7c8185602086016109c6565b80840191505092915050565b7f5d20616d6f756e745b0000000000000000000000000000000000000000000000815250565b7f5d00000000000000000000000000000000000000000000000000000000000000815250565b6000610bdf82610b26565b600982019150610bef8285610b57565b9150610bfa82610b88565b600982019150610c0a8284610b57565b9150610c1582610bae565b6001820191508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610c6c57607f821691505b602082108103610c7f57610c7e610c25565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610cbf8261091a565b9150610cca8361091a565b9250828201905080821115610ce257610ce1610c85565b5b92915050565b7f6465706f7369745b000000000000000000000000000000000000000000000000815250565b7f5d2062616c616e63655b00000000000000000000000000000000000000000000815250565b6000610d3f82610ce8565b600882019150610d4f8285610b57565b9150610d5a82610d0e565b600a82019150610d6a8284610b57565b9150610d7582610bae565b6001820191508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000610dbf8261091a565b9150610dca8361091a565b9250828203905081811115610de257610de1610c85565b5b92915050565b6000610df38261091a565b9150610dfe8361091a565b9250828202610e0c8161091a565b91508282048414831517610e2357610e22610c85565b5b5092915050565b60008160011c9050919050565b6000808291508390505b6001851115610e8157808604811115610e5d57610e5c610c85565b5b6001851615610e6c5780820291505b8081029050610e7a85610e2a565b9450610e41565b94509492505050565b600082610e9a5760019050610f56565b81610ea85760009050610f56565b8160018114610ebe5760028114610ec857610ef7565b6001915050610f56565b60ff841115610eda57610ed9610c85565b5b8360020a915084821115610ef157610ef0610c85565b5b50610f56565b5060208310610133831016604e8410600b8410161715610f2c5782820a905083811115610f2757610f26610c85565b5b610f56565b610f398484846001610e37565b92509050818404811115610f5057610f4f610c85565b5b81810290505b9392505050565b6000610f688261091a565b9150610f738361091a565b9250610fa07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484610e8a565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000610fe28261091a565b9150610fed8361091a565b925082610ffd57610ffc610fa8565b5b828204905092915050565b600060ff82169050919050565b600061102082611008565b915061102b83611008565b92508261103b5761103a610fa8565b5b828204905092915050565b600061105182611008565b915061105c83611008565b925082820261106a81611008565b915080821461107c5761107b610c85565b5b5092915050565b600061108e82611008565b915061109983611008565b9250828203905060ff8111156110b2576110b1610c85565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006110f28261091a565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361112457611123610c85565b5b600182019050919050565b600061113a82611008565b915061114583611008565b9250828201905060ff81111561115e5761115d610c85565b5b9291505056fea26469706673582212201b2b77206cba7848335602cb78cb1098cab065f0e847e0ad0024b75b3e2f16a664736f6c63430008120033",
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
