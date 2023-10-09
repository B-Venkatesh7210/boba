// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/bobanetwork/v3-anchorage/boba-bindings/solc"
)

const OptimismMintableERC20FactoryStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/universal/OptimismMintableERC20Factory.sol:OptimismMintableERC20Factory\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/universal/OptimismMintableERC20Factory.sol:OptimismMintableERC20Factory\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/universal/OptimismMintableERC20Factory.sol:OptimismMintableERC20Factory\",\"label\":\"bridge\",\"offset\":2,\"slot\":\"0\",\"type\":\"t_address\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var OptimismMintableERC20FactoryStorageLayout = new(solc.StorageLayout)

var OptimismMintableERC20FactoryDeployedBin = "0x60806040523480156200001157600080fd5b5060043610620000875760003560e01c8063c4d66de81162000062578063c4d66de81462000135578063ce5ac90f146200014e578063e78cea921462000165578063ee9a31a2146200018c57600080fd5b806354fd4d50146200008c578063896f93d114620000e15780638cf0629c146200011e575b600080fd5b620000c96040518060400160405280600581526020017f312e362e3000000000000000000000000000000000000000000000000000000081525081565b604051620000d89190620005d1565b60405180910390f35b620000f8620000f2366004620006f9565b620001b1565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001620000d8565b620000f86200012f36600462000776565b620001c8565b6200014c620001463660046200080d565b620003c6565b005b620000f86200015f366004620006f9565b62000544565b600054620000f89062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff16620000f8565b6000620001c084848462000544565b949350505050565b600073ffffffffffffffffffffffffffffffffffffffff851662000273576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f4f7074696d69736d4d696e7461626c654552433230466163746f72793a206d7560448201527f73742070726f766964652072656d6f746520746f6b656e20616464726573730060648201526084015b60405180910390fd5b6000858585856040516020016200028e94939291906200082b565b604051602081830303815290604052805190602001209050600081600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1688888888604051620002de9062000555565b620002ee95949392919062000885565b8190604051809103906000f59050801580156200030f573d6000803e3d6000fd5b5090508073ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167fceeb8e7d520d7f3b65fc11a262b91066940193b05d4f93df07cfdced0eb551cf60405160405180910390a360405133815273ffffffffffffffffffffffffffffffffffffffff80891691908316907f52fe89dd5930f343d25650b62fd367bae47088bcddffd2a88350a6ecdd620cdb9060200160405180910390a39695505050505050565b600054600390610100900460ff16158015620003e9575060005460ff8083169116105b62000477576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016200026a565b6000805461010060ff84167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00009092168217177fffffffffffffffffffff000000000000000000000000000000000000000000ff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff6201000073ffffffffffffffffffffffffffffffffffffffff87160216179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b6000620001c08484846012620001c8565b611ad880620008eb83390190565b6000815180845260005b818110156200058b576020818501810151868301820152016200056d565b818111156200059e576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000620005e6602083018462000563565b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff811681146200061257600080fd5b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f8301126200065857600080fd5b813567ffffffffffffffff8082111562000676576200067662000617565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715620006bf57620006bf62000617565b81604052838152866020858801011115620006d957600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806000606084860312156200070f57600080fd5b6200071a84620005ed565b9250602084013567ffffffffffffffff808211156200073857600080fd5b620007468783880162000646565b935060408601359150808211156200075d57600080fd5b506200076c8682870162000646565b9150509250925092565b600080600080608085870312156200078d57600080fd5b6200079885620005ed565b9350602085013567ffffffffffffffff80821115620007b657600080fd5b620007c48883890162000646565b94506040870135915080821115620007db57600080fd5b50620007ea8782880162000646565b925050606085013560ff811681146200080257600080fd5b939692955090935050565b6000602082840312156200082057600080fd5b620005e682620005ed565b73ffffffffffffffffffffffffffffffffffffffff851681526080602082015260006200085c608083018662000563565b828103604084015262000870818662000563565b91505060ff8316606083015295945050505050565b600073ffffffffffffffffffffffffffffffffffffffff808816835280871660208401525060a06040830152620008c060a083018662000563565b8281036060840152620008d4818662000563565b91505060ff83166080830152969550505050505056fe6101406040523480156200001257600080fd5b5060405162001ad838038062001ad8833981016040819052620000359162000178565b600160026000858560036200004b8382620002b3565b5060046200005a8282620002b3565b50505060809290925260a05260c0526001600160a01b0393841660e0529390921661010052505060ff16610120526200037f565b80516001600160a01b0381168114620000a657600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112620000d357600080fd5b81516001600160401b0380821115620000f057620000f0620000ab565b604051601f8301601f19908116603f011681019082821181831017156200011b576200011b620000ab565b816040528381526020925086838588010111156200013857600080fd5b600091505b838210156200015c57858201830151818301840152908201906200013d565b838211156200016e5760008385830101525b9695505050505050565b600080600080600060a086880312156200019157600080fd5b6200019c866200008e565b9450620001ac602087016200008e565b60408701519094506001600160401b0380821115620001ca57600080fd5b620001d889838a01620000c1565b94506060880151915080821115620001ef57600080fd5b50620001fe88828901620000c1565b925050608086015160ff811681146200021657600080fd5b809150509295509295909350565b600181811c908216806200023957607f821691505b6020821081036200025a57634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620002ae57600081815260208120601f850160051c81016020861015620002895750805b601f850160051c820191505b81811015620002aa5782815560010162000295565b5050505b505050565b81516001600160401b03811115620002cf57620002cf620000ab565b620002e781620002e0845462000224565b8462000260565b602080601f8311600181146200031f5760008415620003065750858301515b600019600386901b1c1916600185901b178555620002aa565b600085815260208120601f198616915b8281101562000350578886015182559484019460019091019084016200032f565b50858210156200036f5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60805160a05160c05160e05161010051610120516116ed620003eb6000396000610244015260008181610317015281816103ac015281816105f101526107cb0152600081816101a9015261033d0152600061075a015260006107310152600061070801526116ed6000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c806370a08231116100d8578063ae1f6aaf1161008c578063dd62ed3e11610066578063dd62ed3e14610361578063e78cea9214610315578063ee9a31a2146103a757600080fd5b8063ae1f6aaf14610315578063c01e1bd61461033b578063d6c0b2c41461033b57600080fd5b80639dc29fac116100bd5780639dc29fac146102dc578063a457c2d7146102ef578063a9059cbb1461030257600080fd5b806370a082311461029e57806395d89b41146102d457600080fd5b806323b872dd1161012f5780633950935111610114578063395093511461026e57806340c10f191461028157806354fd4d501461029657600080fd5b806323b872dd1461022a578063313ce5671461023d57600080fd5b806306fdde031161016057806306fdde03146101f0578063095ea7b31461020557806318160ddd1461021857600080fd5b806301ffc9a71461017c578063033964be146101a4575b600080fd5b61018f61018a366004611329565b6103ce565b60405190151581526020015b60405180910390f35b6101cb7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019b565b6101f86104bf565b60405161019b919061139e565b61018f610213366004611418565b610551565b6002545b60405190815260200161019b565b61018f610238366004611442565b610569565b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016815260200161019b565b61018f61027c366004611418565b61058d565b61029461028f366004611418565b6105d9565b005b6101f8610701565b61021c6102ac36600461147e565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6101f86107a4565b6102946102ea366004611418565b6107b3565b61018f6102fd366004611418565b6108ca565b61018f610310366004611418565b61099b565b7f00000000000000000000000000000000000000000000000000000000000000006101cb565b7f00000000000000000000000000000000000000000000000000000000000000006101cb565b61021c61036f366004611499565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6101cb7f000000000000000000000000000000000000000000000000000000000000000081565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007f1d1d8b63000000000000000000000000000000000000000000000000000000007fec4fc8e3000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000851683148061048757507fffffffff00000000000000000000000000000000000000000000000000000000858116908316145b806104b657507fffffffff00000000000000000000000000000000000000000000000000000000858116908216145b95945050505050565b6060600380546104ce906114cc565b80601f01602080910402602001604051908101604052809291908181526020018280546104fa906114cc565b80156105475780601f1061051c57610100808354040283529160200191610547565b820191906000526020600020905b81548152906001019060200180831161052a57829003601f168201915b5050505050905090565b60003361055f8185856109a9565b5060019392505050565b600033610577858285610b5d565b610582858585610c34565b506001949350505050565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919061055f90829086906105d490879061154e565b6109a9565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146106a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603460248201527f4f7074696d69736d4d696e7461626c6545524332303a206f6e6c79206272696460448201527f67652063616e206d696e7420616e64206275726e00000000000000000000000060648201526084015b60405180910390fd5b6106ad8282610ee7565b8173ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885826040516106f591815260200190565b60405180910390a25050565b606061072c7f0000000000000000000000000000000000000000000000000000000000000000611007565b6107557f0000000000000000000000000000000000000000000000000000000000000000611007565b61077e7f0000000000000000000000000000000000000000000000000000000000000000611007565b60405160200161079093929190611566565b604051602081830303815290604052905090565b6060600480546104ce906114cc565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610878576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603460248201527f4f7074696d69736d4d696e7461626c6545524332303a206f6e6c79206272696460448201527f67652063616e206d696e7420616e64206275726e000000000000000000000000606482015260840161069a565b6108828282611144565b8173ffffffffffffffffffffffffffffffffffffffff167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5826040516106f591815260200190565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091908381101561098e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f000000000000000000000000000000000000000000000000000000606482015260840161069a565b61058282868684036109a9565b60003361055f818585610c34565b73ffffffffffffffffffffffffffffffffffffffff8316610a4b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f7265737300000000000000000000000000000000000000000000000000000000606482015260840161069a565b73ffffffffffffffffffffffffffffffffffffffff8216610aee576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f7373000000000000000000000000000000000000000000000000000000000000606482015260840161069a565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610c2e5781811015610c21576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000604482015260640161069a565b610c2e84848484036109a9565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610cd7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f6472657373000000000000000000000000000000000000000000000000000000606482015260840161069a565b73ffffffffffffffffffffffffffffffffffffffff8216610d7a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f6573730000000000000000000000000000000000000000000000000000000000606482015260840161069a565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610e30576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e63650000000000000000000000000000000000000000000000000000606482015260840161069a565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260208190526040808220858503905591851681529081208054849290610e7490849061154e565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610eda91815260200190565b60405180910390a3610c2e565b73ffffffffffffffffffffffffffffffffffffffff8216610f64576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640161069a565b8060026000828254610f76919061154e565b909155505073ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604081208054839290610fb090849061154e565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b60608160000361104a57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115611074578061105e816115dc565b915061106d9050600a83611643565b915061104e565b60008167ffffffffffffffff81111561108f5761108f611657565b6040519080825280601f01601f1916602001820160405280156110b9576020820181803683370190505b5090505b841561113c576110ce600183611686565b91506110db600a8661169d565b6110e690603061154e565b60f81b8183815181106110fb576110fb6116b1565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350611135600a86611643565b94506110bd565b949350505050565b73ffffffffffffffffffffffffffffffffffffffff82166111e7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f7300000000000000000000000000000000000000000000000000000000000000606482015260840161069a565b73ffffffffffffffffffffffffffffffffffffffff82166000908152602081905260409020548181101561129d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f6365000000000000000000000000000000000000000000000000000000000000606482015260840161069a565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604081208383039055600280548492906112d9908490611686565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602001610b50565b60006020828403121561133b57600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461136b57600080fd5b9392505050565b60005b8381101561138d578181015183820152602001611375565b83811115610c2e5750506000910152565b60208152600082518060208401526113bd816040850160208701611372565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461141357600080fd5b919050565b6000806040838503121561142b57600080fd5b611434836113ef565b946020939093013593505050565b60008060006060848603121561145757600080fd5b611460846113ef565b925061146e602085016113ef565b9150604084013590509250925092565b60006020828403121561149057600080fd5b61136b826113ef565b600080604083850312156114ac57600080fd5b6114b5836113ef565b91506114c3602084016113ef565b90509250929050565b600181811c908216806114e057607f821691505b602082108103611519577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082198211156115615761156161151f565b500190565b60008451611578818460208901611372565b80830190507f2e0000000000000000000000000000000000000000000000000000000000000080825285516115b4816001850160208a01611372565b600192019182015283516115cf816002840160208801611372565b0160020195945050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361160d5761160d61151f565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60008261165257611652611614565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000828210156116985761169861151f565b500390565b6000826116ac576116ac611614565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000aa164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(OptimismMintableERC20FactoryStorageLayoutJSON), OptimismMintableERC20FactoryStorageLayout); err != nil {
		panic(err)
	}

	layouts["OptimismMintableERC20Factory"] = OptimismMintableERC20FactoryStorageLayout
	deployedBytecodes["OptimismMintableERC20Factory"] = OptimismMintableERC20FactoryDeployedBin
}
