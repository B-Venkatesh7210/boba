// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"math/big"
	"strings"

	ethereum "github.com/ledgerwatch/erigon"
	libcommon "github.com/ledgerwatch/erigon-lib/common"
	"github.com/ledgerwatch/erigon/accounts/abi"
	"github.com/ledgerwatch/erigon/accounts/abi/bind"
	"github.com/ledgerwatch/erigon/core/types"
	"github.com/ledgerwatch/erigon/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = libcommon.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// LegacyERC20ETHABI is the input ABI used to generate the binding from.
const LegacyERC20ETHABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BRIDGE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REMOTE_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1Token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remoteToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// LegacyERC20ETHBin is the compiled bytecode used for deploying new contracts.
var LegacyERC20ETHBin = "0x6101406040523480156200001257600080fd5b5073420000000000000000000000000000000000001060006040518060400160405280600581526020016422ba3432b960d91b8152506040518060400160405280600381526020016208aa8960eb1b8152506012600160026000858581600390816200007f919062000167565b5060046200008e828262000167565b50505060809290925260a05260c0526001600160a01b0393841660e0529390921661010052505060ff166101205262000233565b634e487b7160e01b600052604160045260246000fd5b600181811c90821680620000ed57607f821691505b6020821081036200010e57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200016257600081815260208120601f850160051c810160208610156200013d5750805b601f850160051c820191505b818110156200015e5782815560010162000149565b5050505b505050565b81516001600160401b03811115620001835762000183620000c2565b6200019b81620001948454620000d8565b8462000114565b602080601f831160018114620001d35760008415620001ba5750858301515b600019600386901b1c1916600185901b1785556200015e565b600085815260208120601f198616915b828110156200020457888601518255948401946001909101908401620001e3565b5085821015620002235787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60805160a05160c05160e0516101005161012051610e79620002916000396000610244015260008181610309015261039e0152600081816101a9015261032f015260006107a40152600061077b015260006107520152610e796000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c806370a08231116100d8578063ae1f6aaf1161008c578063dd62ed3e11610066578063dd62ed3e14610353578063e78cea9214610307578063ee9a31a21461039957600080fd5b8063ae1f6aaf14610307578063c01e1bd61461032d578063d6c0b2c41461032d57600080fd5b80639dc29fac116100bd5780639dc29fac146102ce578063a457c2d7146102e1578063a9059cbb146102f457600080fd5b806370a082311461029e57806395d89b41146102c657600080fd5b806323b872dd1161012f5780633950935111610114578063395093511461026e57806340c10f191461028157806354fd4d501461029657600080fd5b806323b872dd1461022a578063313ce5671461023d57600080fd5b806306fdde031161016057806306fdde03146101f0578063095ea7b31461020557806318160ddd1461021857600080fd5b806301ffc9a71461017c578063033964be146101a4575b600080fd5b61018f61018a366004610ab1565b6103c0565b60405190151581526020015b60405180910390f35b6101cb7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019b565b6101f86104b1565b60405161019b9190610b2a565b61018f610213366004610ba4565b610543565b6002545b60405190815260200161019b565b61018f610238366004610bce565b6105d3565b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016815260200161019b565b61018f61027c366004610ba4565b61065e565b61029461028f366004610ba4565b6106e9565b005b6101f861074b565b61021c6102ac366004610c0a565b73ffffffffffffffffffffffffffffffffffffffff163190565b6101f86107ee565b6102946102dc366004610ba4565b6107fd565b61018f6102ef366004610ba4565b61085f565b61018f610302366004610ba4565b6108ea565b7f00000000000000000000000000000000000000000000000000000000000000006101cb565b7f00000000000000000000000000000000000000000000000000000000000000006101cb565b61021c610361366004610c25565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6101cb7f000000000000000000000000000000000000000000000000000000000000000081565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007f1d1d8b63000000000000000000000000000000000000000000000000000000007fec4fc8e3000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000851683148061047957507fffffffff00000000000000000000000000000000000000000000000000000000858116908316145b806104a857507fffffffff00000000000000000000000000000000000000000000000000000000858116908216145b95945050505050565b6060600380546104c090610c58565b80601f01602080910402602001604051908101604052809291908181526020018280546104ec90610c58565b80156105395780601f1061050e57610100808354040283529160200191610539565b820191906000526020600020905b81548152906001019060200180831161051c57829003601f168201915b5050505050905090565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f4c656761637945524332304554483a20617070726f766520697320646973616260448201527f6c6564000000000000000000000000000000000000000000000000000000000060648201526000906084015b60405180910390fd5b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f4c656761637945524332304554483a207472616e7366657246726f6d2069732060448201527f64697361626c656400000000000000000000000000000000000000000000000060648201526000906084016105ca565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4c656761637945524332304554483a20696e637265617365416c6c6f77616e6360448201527f652069732064697361626c65640000000000000000000000000000000000000060648201526000906084016105ca565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4c656761637945524332304554483a206d696e742069732064697361626c656460448201526064016105ca565b60606107767f0000000000000000000000000000000000000000000000000000000000000000610974565b61079f7f0000000000000000000000000000000000000000000000000000000000000000610974565b6107c87f0000000000000000000000000000000000000000000000000000000000000000610974565b6040516020016107da93929190610cab565b604051602081830303815290604052905090565b6060600480546104c090610c58565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4c656761637945524332304554483a206275726e2069732064697361626c656460448201526064016105ca565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4c656761637945524332304554483a206465637265617365416c6c6f77616e6360448201527f652069732064697361626c65640000000000000000000000000000000000000060648201526000906084016105ca565b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f4c656761637945524332304554483a207472616e73666572206973206469736160448201527f626c65640000000000000000000000000000000000000000000000000000000060648201526000906084016105ca565b6060816000036109b757505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b81156109e157806109cb81610d50565b91506109da9050600a83610db7565b91506109bb565b60008167ffffffffffffffff8111156109fc576109fc610dcb565b6040519080825280601f01601f191660200182016040528015610a26576020820181803683370190505b5090505b8415610aa957610a3b600183610dfa565b9150610a48600a86610e11565b610a53906030610e25565b60f81b818381518110610a6857610a68610e3d565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610aa2600a86610db7565b9450610a2a565b949350505050565b600060208284031215610ac357600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610af357600080fd5b9392505050565b60005b83811015610b15578181015183820152602001610afd565b83811115610b24576000848401525b50505050565b6020815260008251806020840152610b49816040850160208701610afa565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610b9f57600080fd5b919050565b60008060408385031215610bb757600080fd5b610bc083610b7b565b946020939093013593505050565b600080600060608486031215610be357600080fd5b610bec84610b7b565b9250610bfa60208501610b7b565b9150604084013590509250925092565b600060208284031215610c1c57600080fd5b610af382610b7b565b60008060408385031215610c3857600080fd5b610c4183610b7b565b9150610c4f60208401610b7b565b90509250929050565b600181811c90821680610c6c57607f821691505b602082108103610ca5577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b60008451610cbd818460208901610afa565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551610cf9816001850160208a01610afa565b60019201918201528351610d14816002840160208801610afa565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610d8157610d81610d21565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082610dc657610dc6610d88565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082821015610e0c57610e0c610d21565b500390565b600082610e2057610e20610d88565b500690565b60008219821115610e3857610e38610d21565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000a"

// DeployLegacyERC20ETH deploys a new Ethereum contract, binding an instance of LegacyERC20ETH to it.
func DeployLegacyERC20ETH(auth *bind.TransactOpts, backend bind.ContractBackend) (libcommon.Address, types.Transaction, *LegacyERC20ETH, error) {
	parsed, err := abi.JSON(strings.NewReader(LegacyERC20ETHABI))
	if err != nil {
		return libcommon.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, libcommon.FromHex(LegacyERC20ETHBin), backend)
	if err != nil {
		return libcommon.Address{}, nil, nil, err
	}
	return address, tx, &LegacyERC20ETH{LegacyERC20ETHCaller: LegacyERC20ETHCaller{contract: contract}, LegacyERC20ETHTransactor: LegacyERC20ETHTransactor{contract: contract}, LegacyERC20ETHFilterer: LegacyERC20ETHFilterer{contract: contract}}, nil
}

// LegacyERC20ETH is an auto generated Go binding around an Ethereum contract.
type LegacyERC20ETH struct {
	LegacyERC20ETHCaller     // Read-only binding to the contract
	LegacyERC20ETHTransactor // Write-only binding to the contract
	LegacyERC20ETHFilterer   // Log filterer for contract events
}

// LegacyERC20ETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type LegacyERC20ETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyERC20ETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LegacyERC20ETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyERC20ETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LegacyERC20ETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyERC20ETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LegacyERC20ETHSession struct {
	Contract     *LegacyERC20ETH   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LegacyERC20ETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LegacyERC20ETHCallerSession struct {
	Contract *LegacyERC20ETHCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// LegacyERC20ETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LegacyERC20ETHTransactorSession struct {
	Contract     *LegacyERC20ETHTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// LegacyERC20ETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type LegacyERC20ETHRaw struct {
	Contract *LegacyERC20ETH // Generic contract binding to access the raw methods on
}

// LegacyERC20ETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LegacyERC20ETHCallerRaw struct {
	Contract *LegacyERC20ETHCaller // Generic read-only contract binding to access the raw methods on
}

// LegacyERC20ETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LegacyERC20ETHTransactorRaw struct {
	Contract *LegacyERC20ETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLegacyERC20ETH creates a new instance of LegacyERC20ETH, bound to a specific deployed contract.
func NewLegacyERC20ETH(address libcommon.Address, backend bind.ContractBackend) (*LegacyERC20ETH, error) {
	contract, err := bindLegacyERC20ETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LegacyERC20ETH{LegacyERC20ETHCaller: LegacyERC20ETHCaller{contract: contract}, LegacyERC20ETHTransactor: LegacyERC20ETHTransactor{contract: contract}, LegacyERC20ETHFilterer: LegacyERC20ETHFilterer{contract: contract}}, nil
}

// NewLegacyERC20ETHCaller creates a new read-only instance of LegacyERC20ETH, bound to a specific deployed contract.
func NewLegacyERC20ETHCaller(address libcommon.Address, caller bind.ContractCaller) (*LegacyERC20ETHCaller, error) {
	contract, err := bindLegacyERC20ETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LegacyERC20ETHCaller{contract: contract}, nil
}

// NewLegacyERC20ETHTransactor creates a new write-only instance of LegacyERC20ETH, bound to a specific deployed contract.
func NewLegacyERC20ETHTransactor(address libcommon.Address, transactor bind.ContractTransactor) (*LegacyERC20ETHTransactor, error) {
	contract, err := bindLegacyERC20ETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LegacyERC20ETHTransactor{contract: contract}, nil
}

// NewLegacyERC20ETHFilterer creates a new log filterer instance of LegacyERC20ETH, bound to a specific deployed contract.
func NewLegacyERC20ETHFilterer(address libcommon.Address, filterer bind.ContractFilterer) (*LegacyERC20ETHFilterer, error) {
	contract, err := bindLegacyERC20ETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LegacyERC20ETHFilterer{contract: contract}, nil
}

// bindLegacyERC20ETH binds a generic wrapper to an already deployed contract.
func bindLegacyERC20ETH(address libcommon.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LegacyERC20ETHABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LegacyERC20ETH *LegacyERC20ETHRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LegacyERC20ETH.Contract.LegacyERC20ETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LegacyERC20ETH *LegacyERC20ETHRaw) Transfer(opts *bind.TransactOpts) (types.Transaction, error) {
	return _LegacyERC20ETH.Contract.LegacyERC20ETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LegacyERC20ETH *LegacyERC20ETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (types.Transaction, error) {
	return _LegacyERC20ETH.Contract.LegacyERC20ETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LegacyERC20ETH *LegacyERC20ETHCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LegacyERC20ETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LegacyERC20ETH *LegacyERC20ETHTransactorRaw) Transfer(opts *bind.TransactOpts) (types.Transaction, error) {
	return _LegacyERC20ETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LegacyERC20ETH *LegacyERC20ETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (types.Transaction, error) {
	return _LegacyERC20ETH.Contract.contract.Transact(opts, method, params...)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_LegacyERC20ETH *LegacyERC20ETHCaller) BRIDGE(opts *bind.CallOpts) (libcommon.Address, error) {
	var out []interface{}
	err := _LegacyERC20ETH.contract.Call(opts, &out, "BRIDGE")

	if err != nil {
		return *new(libcommon.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(libcommon.Address)).(*libcommon.Address)

	return out0, err

}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_LegacyERC20ETH *LegacyERC20ETHSession) BRIDGE() (libcommon.Address, error) {
	return _LegacyERC20ETH.Contract.BRIDGE(&_LegacyERC20ETH.CallOpts)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_LegacyERC20ETH *LegacyERC20ETHCallerSession) BRIDGE() (libcommon.Address, error) {
	return _LegacyERC20ETH.Contract.BRIDGE(&_LegacyERC20ETH.CallOpts)
}

// REMOTETOKEN is a free data retrieval call binding the contract method 0x033964be.
//
// Solidity: function REMOTE_TOKEN() view returns(address)
func (_LegacyERC20ETH *LegacyERC20ETHCaller) REMOTETOKEN(opts *bind.CallOpts) (libcommon.Address, error) {
	var out []interface{}
	err := _LegacyERC20ETH.contract.Call(opts, &out, "REMOTE_TOKEN")

	if err != nil {
		return *new(libcommon.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(libcommon.Address)).(*libcommon.Address)

	return out0, err

}

// REMOTETOKEN is a free data retrieval call binding the contract method 0x033964be.
//
// Solidity: function REMOTE_TOKEN() view returns(address)
func (_LegacyERC20ETH *LegacyERC20ETHSession) REMOTETOKEN() (libcommon.Address, error) {
	return _LegacyERC20ETH.Contract.REMOTETOKEN(&_LegacyERC20ETH.CallOpts)
}

// REMOTETOKEN is a free data retrieval call binding the contract method 0x033964be.
//
// Solidity: function REMOTE_TOKEN() view returns(address)
func (_LegacyERC20ETH *LegacyERC20ETHCallerSession) REMOTETOKEN() (libcommon.Address, error) {
	return _LegacyERC20ETH.Contract.REMOTETOKEN(&_LegacyERC20ETH.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LegacyERC20ETH *LegacyERC20ETHCaller) Allowance(opts *bind.CallOpts, owner libcommon.Address, spender libcommon.Address) (*big.Int, error) {
	var out []interface{}
	err := _LegacyERC20ETH.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LegacyERC20ETH *LegacyERC20ETHSession) Allowance(owner libcommon.Address, spender libcommon.Address) (*big.Int, error) {
	return _LegacyERC20ETH.Contract.Allowance(&_LegacyERC20ETH.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LegacyERC20ETH *LegacyERC20ETHCallerSession) Allowance(owner libcommon.Address, spender libcommon.Address) (*big.Int, error) {
	return _LegacyERC20ETH.Contract.Allowance(&_LegacyERC20ETH.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _who) view returns(uint256)
func (_LegacyERC20ETH *LegacyERC20ETHCaller) BalanceOf(opts *bind.CallOpts, _who libcommon.Address) (*big.Int, error) {
	var out []interface{}
	err := _LegacyERC20ETH.contract.Call(opts, &out, "balanceOf", _who)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _who) view returns(uint256)
func (_LegacyERC20ETH *LegacyERC20ETHSession) BalanceOf(_who libcommon.Address) (*big.Int, error) {
	return _LegacyERC20ETH.Contract.BalanceOf(&_LegacyERC20ETH.CallOpts, _who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _who) view returns(uint256)
func (_LegacyERC20ETH *LegacyERC20ETHCallerSession) BalanceOf(_who libcommon.Address) (*big.Int, error) {
	return _LegacyERC20ETH.Contract.BalanceOf(&_LegacyERC20ETH.CallOpts, _who)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_LegacyERC20ETH *LegacyERC20ETHCaller) Bridge(opts *bind.CallOpts) (libcommon.Address, error) {
	var out []interface{}
	err := _LegacyERC20ETH.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(libcommon.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(libcommon.Address)).(*libcommon.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_LegacyERC20ETH *LegacyERC20ETHSession) Bridge() (libcommon.Address, error) {
	return _LegacyERC20ETH.Contract.Bridge(&_LegacyERC20ETH.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_LegacyERC20ETH *LegacyERC20ETHCallerSession) Bridge() (libcommon.Address, error) {
	return _LegacyERC20ETH.Contract.Bridge(&_LegacyERC20ETH.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LegacyERC20ETH *LegacyERC20ETHCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LegacyERC20ETH.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LegacyERC20ETH *LegacyERC20ETHSession) Decimals() (uint8, error) {
	return _LegacyERC20ETH.Contract.Decimals(&_LegacyERC20ETH.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LegacyERC20ETH *LegacyERC20ETHCallerSession) Decimals() (uint8, error) {
	return _LegacyERC20ETH.Contract.Decimals(&_LegacyERC20ETH.CallOpts)
}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_LegacyERC20ETH *LegacyERC20ETHCaller) L1Token(opts *bind.CallOpts) (libcommon.Address, error) {
	var out []interface{}
	err := _LegacyERC20ETH.contract.Call(opts, &out, "l1Token")

	if err != nil {
		return *new(libcommon.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(libcommon.Address)).(*libcommon.Address)

	return out0, err

}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_LegacyERC20ETH *LegacyERC20ETHSession) L1Token() (libcommon.Address, error) {
	return _LegacyERC20ETH.Contract.L1Token(&_LegacyERC20ETH.CallOpts)
}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_LegacyERC20ETH *LegacyERC20ETHCallerSession) L1Token() (libcommon.Address, error) {
	return _LegacyERC20ETH.Contract.L1Token(&_LegacyERC20ETH.CallOpts)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_LegacyERC20ETH *LegacyERC20ETHCaller) L2Bridge(opts *bind.CallOpts) (libcommon.Address, error) {
	var out []interface{}
	err := _LegacyERC20ETH.contract.Call(opts, &out, "l2Bridge")

	if err != nil {
		return *new(libcommon.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(libcommon.Address)).(*libcommon.Address)

	return out0, err

}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_LegacyERC20ETH *LegacyERC20ETHSession) L2Bridge() (libcommon.Address, error) {
	return _LegacyERC20ETH.Contract.L2Bridge(&_LegacyERC20ETH.CallOpts)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_LegacyERC20ETH *LegacyERC20ETHCallerSession) L2Bridge() (libcommon.Address, error) {
	return _LegacyERC20ETH.Contract.L2Bridge(&_LegacyERC20ETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LegacyERC20ETH *LegacyERC20ETHCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LegacyERC20ETH.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LegacyERC20ETH *LegacyERC20ETHSession) Name() (string, error) {
	return _LegacyERC20ETH.Contract.Name(&_LegacyERC20ETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LegacyERC20ETH *LegacyERC20ETHCallerSession) Name() (string, error) {
	return _LegacyERC20ETH.Contract.Name(&_LegacyERC20ETH.CallOpts)
}

// RemoteToken is a free data retrieval call binding the contract method 0xd6c0b2c4.
//
// Solidity: function remoteToken() view returns(address)
func (_LegacyERC20ETH *LegacyERC20ETHCaller) RemoteToken(opts *bind.CallOpts) (libcommon.Address, error) {
	var out []interface{}
	err := _LegacyERC20ETH.contract.Call(opts, &out, "remoteToken")

	if err != nil {
		return *new(libcommon.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(libcommon.Address)).(*libcommon.Address)

	return out0, err

}

// RemoteToken is a free data retrieval call binding the contract method 0xd6c0b2c4.
//
// Solidity: function remoteToken() view returns(address)
func (_LegacyERC20ETH *LegacyERC20ETHSession) RemoteToken() (libcommon.Address, error) {
	return _LegacyERC20ETH.Contract.RemoteToken(&_LegacyERC20ETH.CallOpts)
}

// RemoteToken is a free data retrieval call binding the contract method 0xd6c0b2c4.
//
// Solidity: function remoteToken() view returns(address)
func (_LegacyERC20ETH *LegacyERC20ETHCallerSession) RemoteToken() (libcommon.Address, error) {
	return _LegacyERC20ETH.Contract.RemoteToken(&_LegacyERC20ETH.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool)
func (_LegacyERC20ETH *LegacyERC20ETHCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LegacyERC20ETH.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool)
func (_LegacyERC20ETH *LegacyERC20ETHSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _LegacyERC20ETH.Contract.SupportsInterface(&_LegacyERC20ETH.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool)
func (_LegacyERC20ETH *LegacyERC20ETHCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _LegacyERC20ETH.Contract.SupportsInterface(&_LegacyERC20ETH.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LegacyERC20ETH *LegacyERC20ETHCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LegacyERC20ETH.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LegacyERC20ETH *LegacyERC20ETHSession) Symbol() (string, error) {
	return _LegacyERC20ETH.Contract.Symbol(&_LegacyERC20ETH.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LegacyERC20ETH *LegacyERC20ETHCallerSession) Symbol() (string, error) {
	return _LegacyERC20ETH.Contract.Symbol(&_LegacyERC20ETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LegacyERC20ETH *LegacyERC20ETHCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LegacyERC20ETH.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LegacyERC20ETH *LegacyERC20ETHSession) TotalSupply() (*big.Int, error) {
	return _LegacyERC20ETH.Contract.TotalSupply(&_LegacyERC20ETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LegacyERC20ETH *LegacyERC20ETHCallerSession) TotalSupply() (*big.Int, error) {
	return _LegacyERC20ETH.Contract.TotalSupply(&_LegacyERC20ETH.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LegacyERC20ETH *LegacyERC20ETHCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LegacyERC20ETH.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LegacyERC20ETH *LegacyERC20ETHSession) Version() (string, error) {
	return _LegacyERC20ETH.Contract.Version(&_LegacyERC20ETH.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LegacyERC20ETH *LegacyERC20ETHCallerSession) Version() (string, error) {
	return _LegacyERC20ETH.Contract.Version(&_LegacyERC20ETH.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_LegacyERC20ETH *LegacyERC20ETHTransactor) Approve(opts *bind.TransactOpts, arg0 libcommon.Address, arg1 *big.Int) (types.Transaction, error) {
	return _LegacyERC20ETH.contract.Transact(opts, "approve", arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_LegacyERC20ETH *LegacyERC20ETHSession) Approve(arg0 libcommon.Address, arg1 *big.Int) (types.Transaction, error) {
	return _LegacyERC20ETH.Contract.Approve(&_LegacyERC20ETH.TransactOpts, arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_LegacyERC20ETH *LegacyERC20ETHTransactorSession) Approve(arg0 libcommon.Address, arg1 *big.Int) (types.Transaction, error) {
	return _LegacyERC20ETH.Contract.Approve(&_LegacyERC20ETH.TransactOpts, arg0, arg1)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address , uint256 ) returns()
func (_LegacyERC20ETH *LegacyERC20ETHTransactor) Burn(opts *bind.TransactOpts, arg0 libcommon.Address, arg1 *big.Int) (types.Transaction, error) {
	return _LegacyERC20ETH.contract.Transact(opts, "burn", arg0, arg1)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address , uint256 ) returns()
func (_LegacyERC20ETH *LegacyERC20ETHSession) Burn(arg0 libcommon.Address, arg1 *big.Int) (types.Transaction, error) {
	return _LegacyERC20ETH.Contract.Burn(&_LegacyERC20ETH.TransactOpts, arg0, arg1)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address , uint256 ) returns()
func (_LegacyERC20ETH *LegacyERC20ETHTransactorSession) Burn(arg0 libcommon.Address, arg1 *big.Int) (types.Transaction, error) {
	return _LegacyERC20ETH.Contract.Burn(&_LegacyERC20ETH.TransactOpts, arg0, arg1)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address , uint256 ) returns(bool)
func (_LegacyERC20ETH *LegacyERC20ETHTransactor) DecreaseAllowance(opts *bind.TransactOpts, arg0 libcommon.Address, arg1 *big.Int) (types.Transaction, error) {
	return _LegacyERC20ETH.contract.Transact(opts, "decreaseAllowance", arg0, arg1)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address , uint256 ) returns(bool)
func (_LegacyERC20ETH *LegacyERC20ETHSession) DecreaseAllowance(arg0 libcommon.Address, arg1 *big.Int) (types.Transaction, error) {
	return _LegacyERC20ETH.Contract.DecreaseAllowance(&_LegacyERC20ETH.TransactOpts, arg0, arg1)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address , uint256 ) returns(bool)
func (_LegacyERC20ETH *LegacyERC20ETHTransactorSession) DecreaseAllowance(arg0 libcommon.Address, arg1 *big.Int) (types.Transaction, error) {
	return _LegacyERC20ETH.Contract.DecreaseAllowance(&_LegacyERC20ETH.TransactOpts, arg0, arg1)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address , uint256 ) returns(bool)
func (_LegacyERC20ETH *LegacyERC20ETHTransactor) IncreaseAllowance(opts *bind.TransactOpts, arg0 libcommon.Address, arg1 *big.Int) (types.Transaction, error) {
	return _LegacyERC20ETH.contract.Transact(opts, "increaseAllowance", arg0, arg1)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address , uint256 ) returns(bool)
func (_LegacyERC20ETH *LegacyERC20ETHSession) IncreaseAllowance(arg0 libcommon.Address, arg1 *big.Int) (types.Transaction, error) {
	return _LegacyERC20ETH.Contract.IncreaseAllowance(&_LegacyERC20ETH.TransactOpts, arg0, arg1)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address , uint256 ) returns(bool)
func (_LegacyERC20ETH *LegacyERC20ETHTransactorSession) IncreaseAllowance(arg0 libcommon.Address, arg1 *big.Int) (types.Transaction, error) {
	return _LegacyERC20ETH.Contract.IncreaseAllowance(&_LegacyERC20ETH.TransactOpts, arg0, arg1)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address , uint256 ) returns()
func (_LegacyERC20ETH *LegacyERC20ETHTransactor) Mint(opts *bind.TransactOpts, arg0 libcommon.Address, arg1 *big.Int) (types.Transaction, error) {
	return _LegacyERC20ETH.contract.Transact(opts, "mint", arg0, arg1)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address , uint256 ) returns()
func (_LegacyERC20ETH *LegacyERC20ETHSession) Mint(arg0 libcommon.Address, arg1 *big.Int) (types.Transaction, error) {
	return _LegacyERC20ETH.Contract.Mint(&_LegacyERC20ETH.TransactOpts, arg0, arg1)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address , uint256 ) returns()
func (_LegacyERC20ETH *LegacyERC20ETHTransactorSession) Mint(arg0 libcommon.Address, arg1 *big.Int) (types.Transaction, error) {
	return _LegacyERC20ETH.Contract.Mint(&_LegacyERC20ETH.TransactOpts, arg0, arg1)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) returns(bool)
func (_LegacyERC20ETH *LegacyERC20ETHTransactor) Transfer(opts *bind.TransactOpts, arg0 libcommon.Address, arg1 *big.Int) (types.Transaction, error) {
	return _LegacyERC20ETH.contract.Transact(opts, "transfer", arg0, arg1)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) returns(bool)
func (_LegacyERC20ETH *LegacyERC20ETHSession) Transfer(arg0 libcommon.Address, arg1 *big.Int) (types.Transaction, error) {
	return _LegacyERC20ETH.Contract.Transfer(&_LegacyERC20ETH.TransactOpts, arg0, arg1)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) returns(bool)
func (_LegacyERC20ETH *LegacyERC20ETHTransactorSession) Transfer(arg0 libcommon.Address, arg1 *big.Int) (types.Transaction, error) {
	return _LegacyERC20ETH.Contract.Transfer(&_LegacyERC20ETH.TransactOpts, arg0, arg1)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns(bool)
func (_LegacyERC20ETH *LegacyERC20ETHTransactor) TransferFrom(opts *bind.TransactOpts, arg0 libcommon.Address, arg1 libcommon.Address, arg2 *big.Int) (types.Transaction, error) {
	return _LegacyERC20ETH.contract.Transact(opts, "transferFrom", arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns(bool)
func (_LegacyERC20ETH *LegacyERC20ETHSession) TransferFrom(arg0 libcommon.Address, arg1 libcommon.Address, arg2 *big.Int) (types.Transaction, error) {
	return _LegacyERC20ETH.Contract.TransferFrom(&_LegacyERC20ETH.TransactOpts, arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns(bool)
func (_LegacyERC20ETH *LegacyERC20ETHTransactorSession) TransferFrom(arg0 libcommon.Address, arg1 libcommon.Address, arg2 *big.Int) (types.Transaction, error) {
	return _LegacyERC20ETH.Contract.TransferFrom(&_LegacyERC20ETH.TransactOpts, arg0, arg1, arg2)
}

// LegacyERC20ETHApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LegacyERC20ETH contract.
type LegacyERC20ETHApprovalIterator struct {
	Event *LegacyERC20ETHApproval // Event containing the contract specifics and raw log

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
func (it *LegacyERC20ETHApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyERC20ETHApproval)
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
		it.Event = new(LegacyERC20ETHApproval)
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
func (it *LegacyERC20ETHApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyERC20ETHApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyERC20ETHApproval represents a Approval event raised by the LegacyERC20ETH contract.
type LegacyERC20ETHApproval struct {
	Owner   libcommon.Address
	Spender libcommon.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LegacyERC20ETH *LegacyERC20ETHFilterer) FilterApproval(opts *bind.FilterOpts, owner []libcommon.Address, spender []libcommon.Address) (*LegacyERC20ETHApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LegacyERC20ETH.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LegacyERC20ETHApprovalIterator{contract: _LegacyERC20ETH.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LegacyERC20ETH *LegacyERC20ETHFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LegacyERC20ETHApproval, owner []libcommon.Address, spender []libcommon.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LegacyERC20ETH.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyERC20ETHApproval)
				if err := _LegacyERC20ETH.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LegacyERC20ETH *LegacyERC20ETHFilterer) ParseApproval(log types.Log) (*LegacyERC20ETHApproval, error) {
	event := new(LegacyERC20ETHApproval)
	if err := _LegacyERC20ETH.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacyERC20ETHBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the LegacyERC20ETH contract.
type LegacyERC20ETHBurnIterator struct {
	Event *LegacyERC20ETHBurn // Event containing the contract specifics and raw log

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
func (it *LegacyERC20ETHBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyERC20ETHBurn)
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
		it.Event = new(LegacyERC20ETHBurn)
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
func (it *LegacyERC20ETHBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyERC20ETHBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyERC20ETHBurn represents a Burn event raised by the LegacyERC20ETH contract.
type LegacyERC20ETHBurn struct {
	Account libcommon.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed account, uint256 amount)
func (_LegacyERC20ETH *LegacyERC20ETHFilterer) FilterBurn(opts *bind.FilterOpts, account []libcommon.Address) (*LegacyERC20ETHBurnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LegacyERC20ETH.contract.FilterLogs(opts, "Burn", accountRule)
	if err != nil {
		return nil, err
	}
	return &LegacyERC20ETHBurnIterator{contract: _LegacyERC20ETH.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed account, uint256 amount)
func (_LegacyERC20ETH *LegacyERC20ETHFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *LegacyERC20ETHBurn, account []libcommon.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LegacyERC20ETH.contract.WatchLogs(opts, "Burn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyERC20ETHBurn)
				if err := _LegacyERC20ETH.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed account, uint256 amount)
func (_LegacyERC20ETH *LegacyERC20ETHFilterer) ParseBurn(log types.Log) (*LegacyERC20ETHBurn, error) {
	event := new(LegacyERC20ETHBurn)
	if err := _LegacyERC20ETH.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacyERC20ETHMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the LegacyERC20ETH contract.
type LegacyERC20ETHMintIterator struct {
	Event *LegacyERC20ETHMint // Event containing the contract specifics and raw log

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
func (it *LegacyERC20ETHMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyERC20ETHMint)
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
		it.Event = new(LegacyERC20ETHMint)
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
func (it *LegacyERC20ETHMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyERC20ETHMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyERC20ETHMint represents a Mint event raised by the LegacyERC20ETH contract.
type LegacyERC20ETHMint struct {
	Account libcommon.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed account, uint256 amount)
func (_LegacyERC20ETH *LegacyERC20ETHFilterer) FilterMint(opts *bind.FilterOpts, account []libcommon.Address) (*LegacyERC20ETHMintIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LegacyERC20ETH.contract.FilterLogs(opts, "Mint", accountRule)
	if err != nil {
		return nil, err
	}
	return &LegacyERC20ETHMintIterator{contract: _LegacyERC20ETH.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed account, uint256 amount)
func (_LegacyERC20ETH *LegacyERC20ETHFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *LegacyERC20ETHMint, account []libcommon.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LegacyERC20ETH.contract.WatchLogs(opts, "Mint", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyERC20ETHMint)
				if err := _LegacyERC20ETH.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed account, uint256 amount)
func (_LegacyERC20ETH *LegacyERC20ETHFilterer) ParseMint(log types.Log) (*LegacyERC20ETHMint, error) {
	event := new(LegacyERC20ETHMint)
	if err := _LegacyERC20ETH.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacyERC20ETHTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LegacyERC20ETH contract.
type LegacyERC20ETHTransferIterator struct {
	Event *LegacyERC20ETHTransfer // Event containing the contract specifics and raw log

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
func (it *LegacyERC20ETHTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyERC20ETHTransfer)
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
		it.Event = new(LegacyERC20ETHTransfer)
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
func (it *LegacyERC20ETHTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyERC20ETHTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyERC20ETHTransfer represents a Transfer event raised by the LegacyERC20ETH contract.
type LegacyERC20ETHTransfer struct {
	From  libcommon.Address
	To    libcommon.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LegacyERC20ETH *LegacyERC20ETHFilterer) FilterTransfer(opts *bind.FilterOpts, from []libcommon.Address, to []libcommon.Address) (*LegacyERC20ETHTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LegacyERC20ETH.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LegacyERC20ETHTransferIterator{contract: _LegacyERC20ETH.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LegacyERC20ETH *LegacyERC20ETHFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LegacyERC20ETHTransfer, from []libcommon.Address, to []libcommon.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LegacyERC20ETH.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyERC20ETHTransfer)
				if err := _LegacyERC20ETH.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LegacyERC20ETH *LegacyERC20ETHFilterer) ParseTransfer(log types.Log) (*LegacyERC20ETHTransfer, error) {
	event := new(LegacyERC20ETHTransfer)
	if err := _LegacyERC20ETH.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
