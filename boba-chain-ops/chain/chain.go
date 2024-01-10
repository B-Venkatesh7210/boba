package chain

import (
	"errors"
	"math/big"
)

var (
	// Boba Mainnet
	BobaMainnetChainId = big.NewInt(288)
	// Boba Mainnet genesis gas limit
	BobaMainnetGenesisGasLimit = 11000000
	// Boba Mainnet genesis block coinbase
	BobaMainnetGenesisCoinbase = "0x0000000000000000000000000000000000000000"
	// Boba Mainnet genesis block extra data
	BobaMainnetGenesisExtraData = "000000000000000000000000000000000000000000000000000000000000000000000398232e2064f896018496b4b44b3d62751f0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	// Boba Mainnet genesis root
	BobaMainnetGenesisRoot = "0x7ec54492a4504ff1ef3491825cd55e01e5c75409e4287129170e98d4693848ce"
	// Boba Mainnet genesis block hash
	BobaMainnetGenesisBlockHash = "0xdcd9e6a8f9973eaa62da2874959cb152faeb4fd6929177bd6335a1a16074ef9c"
	// Mainnet L1 BOBA Address
	BobaTokenMainnetL1Address = "0x42bBFa2e77757C645eeaAd1655E0911a7553Efbc"

	// Boba Sepolia
	BobaSepoliaChainId = big.NewInt(28882)
	// Boba Sepolia genesis gas limit
	BobaSepoliaGenesisGasLimit = 11000000
	// Boba Sepolia genesis block coinbase
	BobaSepoliaGenesisCoinbase = "0x0000000000000000000000000000000000000000"
	// Boba Sepolia genesis block extra data
	BobaSepoliaGenesisExtraData = "000000000000000000000000000000000000000000000000000000000000000000000398232e2064f896018496b4b44b3d62751f0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	// Boba Sepolia genesis root
	BobaSepoliaGenesisRoot = "0x65c5f098877579209b07264284de2fd5b40d326a843a5c79692c12e42b4867bc"
	// Boba Sepolia genesis block hash
	BobaSepoliaGenesisBlockHash = "0x8c571a8c282dcba86a5f47b823defbf108999e8375982d339122b9e7264993b1"
	// Sepolia L1 BOBA Address
	BobaTokenSepoliaL1Address = "0x33faF65b3DfcC6A1FccaD4531D9ce518F0FDc896"

	// error
	ErrInvalidChainID = errors.New("invalid chain id")
)

func IsBobaValidChainId(chainId *big.Int) bool {
	// Mainnet
	if BobaMainnetChainId.Cmp(chainId) == 0 {
		return true
	}
	// Sepolia
	if BobaSepoliaChainId.Cmp(chainId) == 0 {
		return true
	}
	return false
}

func GetBobaGenesisGasLimit(chainId *big.Int) int {
	// Mainnet
	if BobaMainnetChainId.Cmp(chainId) == 0 {
		return BobaMainnetGenesisGasLimit
	}
	// Sepolia
	if BobaSepoliaChainId.Cmp(chainId) == 0 {
		return BobaSepoliaGenesisGasLimit
	}
	return 11000000
}

func GetBobaGenesisCoinbase(chainId *big.Int) string {
	// Mainnet
	if BobaMainnetChainId.Cmp(chainId) == 0 {
		return BobaMainnetGenesisCoinbase
	}
	// Sepolia
	if BobaSepoliaChainId.Cmp(chainId) == 0 {
		return BobaSepoliaGenesisCoinbase
	}
	return "0x0000000000000000000000000000000000000000"
}

func GetBobaGenesisExtraData(chainId *big.Int) string {
	// Mainnet
	if BobaMainnetChainId.Cmp(chainId) == 0 {
		return BobaMainnetGenesisExtraData
	}
	// Sepolia
	if BobaSepoliaChainId.Cmp(chainId) == 0 {
		return BobaSepoliaGenesisExtraData
	}
	return ""
}

func GetBobaGenesisRoot(chainId *big.Int) string {
	// Mainnet
	if BobaMainnetChainId.Cmp(chainId) == 0 {
		return BobaMainnetGenesisRoot
	}
	// Sepolia
	if BobaSepoliaChainId.Cmp(chainId) == 0 {
		return BobaSepoliaGenesisRoot
	}
	return ""
}

func GetBobaGenesisHash(chainId *big.Int) string {
	// Mainnet
	if BobaMainnetChainId.Cmp(chainId) == 0 {
		return BobaMainnetGenesisBlockHash
	}
	// Sepolia
	if BobaSepoliaChainId.Cmp(chainId) == 0 {
		return BobaSepoliaGenesisBlockHash
	}
	return ""
}

func GetBobaTokenL1Address(chainId *big.Int) string {
	// Mainnet
	if BobaMainnetChainId.Cmp(chainId) == 0 {
		return BobaTokenMainnetL1Address
	}
	// Sepolia
	if BobaSepoliaChainId.Cmp(chainId) == 0 {
		return BobaTokenSepoliaL1Address
	}
	return "0x0000000000000000000000000000000000000000"
}
