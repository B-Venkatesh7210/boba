// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/bobanetwork/v3-anchorage/boba-bindings/solc"
)

const L2ERC721BridgeStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L2/L2ERC721Bridge.sol:L2ERC721Bridge\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/L2/L2ERC721Bridge.sol:L2ERC721Bridge\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/L2/L2ERC721Bridge.sol:L2ERC721Bridge\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)48_storage\"}],\"types\":{\"t_array(t_uint256)48_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[48]\",\"numberOfBytes\":\"1536\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L2ERC721BridgeStorageLayout = new(solc.StorageLayout)

var L2ERC721BridgeDeployedBin = "0x608060405234801561001057600080fd5b50600436106100be5760003560e01c80637f46ddb211610076578063927ede2d1161005b578063927ede2d146101be578063aa557452146101e5578063c89701a2146101f857600080fd5b80637f46ddb21461018f5780638129fc1c146101b657600080fd5b806354fd4d50116100a757806354fd4d50146101245780635c975abb1461016d578063761f44931461017c57600080fd5b80633687011a146100c35780633cb747bf146100d8575b600080fd5b6100d66100d136600461120c565b61021e565b005b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6101606040518060400160405280600581526020017f312e362e3000000000000000000000000000000000000000000000000000000081525081565b60405161011b91906112fa565b6040516000815260200161011b565b6100d661018a36600461130d565b6102ca565b6100fa7f000000000000000000000000000000000000000000000000000000000000000081565b6100d6610831565b6100fa7f000000000000000000000000000000000000000000000000000000000000000081565b6100d66101f33660046113a5565b6109bb565b7f00000000000000000000000000000000000000000000000000000000000000006100fa565b333b156102b2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4552433732314272696467653a206163636f756e74206973206e6f742065787460448201527f65726e616c6c79206f776e65640000000000000000000000000000000000000060648201526084015b60405180910390fd5b6102c28686333388888888610a93565b505050505050565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161480156103e857507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103ac573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103d0919061141c565b73ffffffffffffffffffffffffffffffffffffffff16145b610474576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f4552433732314272696467653a2066756e6374696f6e2063616e206f6e6c792060448201527f62652063616c6c65642066726f6d20746865206f74686572206272696467650060648201526084016102a9565b3073ffffffffffffffffffffffffffffffffffffffff881603610519576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f4c324552433732314272696467653a206c6f63616c20746f6b656e2063616e6e60448201527f6f742062652073656c660000000000000000000000000000000000000000000060648201526084016102a9565b610543877f74259ebf00000000000000000000000000000000000000000000000000000000611031565b6105cf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f4c324552433732314272696467653a206c6f63616c20746f6b656e20696e746560448201527f7266616365206973206e6f7420636f6d706c69616e740000000000000000000060648201526084016102a9565b8673ffffffffffffffffffffffffffffffffffffffff1663d6c0b2c46040518163ffffffff1660e01b8152600401602060405180830381865afa15801561061a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061063e919061141c565b73ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff161461071e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604b60248201527f4c324552433732314272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433732312060648201527f6c6f63616c20746f6b656e000000000000000000000000000000000000000000608482015260a4016102a9565b6040517fa144819400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301526024820185905288169063a144819490604401600060405180830381600087803b15801561078e57600080fd5b505af11580156107a2573d6000803e3d6000fd5b505050508473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f1f39bf6707b5d608453e0ae4c067b562bcc4c85c0f562ef5d2c774d2e7f131ac878787876040516108209493929190611482565b60405180910390a450505050505050565b600054610100900460ff16158080156108515750600054600160ff909116105b8061086b5750303b15801561086b575060005460ff166001145b6108f7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016102a9565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561095557600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b80156109b857600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b73ffffffffffffffffffffffffffffffffffffffff8516610a5e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f4552433732314272696467653a206e667420726563697069656e742063616e6e60448201527f6f7420626520616464726573732830290000000000000000000000000000000060648201526084016102a9565b610a6e8787338888888888610a93565b50505050505050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b73ffffffffffffffffffffffffffffffffffffffff8716610b36576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f4c324552433732314272696467653a2072656d6f746520746f6b656e2063616e60448201527f6e6f74206265206164647265737328302900000000000000000000000000000060648201526084016102a9565b6040517f6352211e0000000000000000000000000000000000000000000000000000000081526004810185905273ffffffffffffffffffffffffffffffffffffffff891690636352211e90602401602060405180830381865afa158015610ba1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bc5919061141c565b73ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614610c7f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603e60248201527f4c324552433732314272696467653a205769746864726177616c206973206e6f60448201527f74206265696e6720696e69746961746564206279204e4654206f776e6572000060648201526084016102a9565b60008873ffffffffffffffffffffffffffffffffffffffff1663d6c0b2c46040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ccc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cf0919061141c565b90508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610dad576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f4c324552433732314272696467653a2072656d6f746520746f6b656e20646f6560448201527f73206e6f74206d6174636820676976656e2076616c756500000000000000000060648201526084016102a9565b6040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8881166004830152602482018790528a1690639dc29fac90604401600060405180830381600087803b158015610e1d57600080fd5b505af1158015610e31573d6000803e3d6000fd5b50505050600063761f449360e01b828b8a8a8a8989604051602401610e5c97969594939291906114c2565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009094169390931790925290517f3dbb202b00000000000000000000000000000000000000000000000000000000815290915073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690633dbb202b90610f71907f00000000000000000000000000000000000000000000000000000000000000009085908a9060040161151f565b600060405180830381600087803b158015610f8b57600080fd5b505af1158015610f9f573d6000803e3d6000fd5b505050508773ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff168b73ffffffffffffffffffffffffffffffffffffffff167fb7460e2a880f256ebef3406116ff3eee0cee51ebccdc2a40698f87ebb2e9c1a58a8a898960405161101d9493929190611482565b60405180910390a450505050505050505050565b600061103c83611054565b801561104d575061104d83836110b9565b9392505050565b6000611080827f01ffc9a7000000000000000000000000000000000000000000000000000000006110b9565b80156110b357506110b1827fffffffff000000000000000000000000000000000000000000000000000000006110b9565b155b92915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d91506000519050828015611171575060208210155b801561117d5750600081115b979650505050505050565b73ffffffffffffffffffffffffffffffffffffffff811681146109b857600080fd5b803563ffffffff811681146111be57600080fd5b919050565b60008083601f8401126111d557600080fd5b50813567ffffffffffffffff8111156111ed57600080fd5b60208301915083602082850101111561120557600080fd5b9250929050565b60008060008060008060a0878903121561122557600080fd5b863561123081611188565b9550602087013561124081611188565b945060408701359350611255606088016111aa565b9250608087013567ffffffffffffffff81111561127157600080fd5b61127d89828a016111c3565b979a9699509497509295939492505050565b6000815180845260005b818110156112b557602081850181015186830182015201611299565b818111156112c7576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061104d602083018461128f565b600080600080600080600060c0888a03121561132857600080fd5b873561133381611188565b9650602088013561134381611188565b9550604088013561135381611188565b9450606088013561136381611188565b93506080880135925060a088013567ffffffffffffffff81111561138657600080fd5b6113928a828b016111c3565b989b979a50959850939692959293505050565b600080600080600080600060c0888a0312156113c057600080fd5b87356113cb81611188565b965060208801356113db81611188565b955060408801356113eb81611188565b945060608801359350611400608089016111aa565b925060a088013567ffffffffffffffff81111561138657600080fd5b60006020828403121561142e57600080fd5b815161104d81611188565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff851681528360208201526060604082015260006114b8606083018486611439565b9695505050505050565b600073ffffffffffffffffffffffffffffffffffffffff808a1683528089166020840152808816604084015280871660608401525084608083015260c060a083015261151260c083018486611439565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8416815260606020820152600061154e606083018561128f565b905063ffffffff8316604083015294935050505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L2ERC721BridgeStorageLayoutJSON), L2ERC721BridgeStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2ERC721Bridge"] = L2ERC721BridgeStorageLayout
	deployedBytecodes["L2ERC721Bridge"] = L2ERC721BridgeDeployedBin
}
