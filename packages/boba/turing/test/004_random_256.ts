import { BigNumber, Contract, ContractFactory, providers, Wallet } from 'ethers'
import { ethers } from 'hardhat'
import chai, { expect } from 'chai'
import { solidity } from 'ethereum-waffle'
chai.use(solidity)
const abiDecoder = require('web3-eth-abi')

const fetch = require('node-fetch')
import hre from 'hardhat'
const cfg = hre.network.config

const gasOverride =  {
  gasLimit: 3000000 //3,000,000
}

const helperPredeploy = '0x4200000000000000000000000000000000000022'
import HelloTuringJson from "../artifacts/contracts/HelloTuring.sol/HelloTuring.json"
import TuringHelper from "../artifacts/contracts/TuringHelper.sol/TuringHelper.json"

let Factory__Turing: ContractFactory
let turing: Contract
let Factory__Helper: ContractFactory
let helper: Contract

const local_provider = new providers.JsonRpcProvider(cfg['url'])

// Key for Hardhat test account #13 (0x1cbd3b2770909d4e10f157cabc84c7264073c9ec)
const testPrivateKey = '0x47c99abed3324a2707c28affff1267e45918ec8c3f20b8aa892e8b065d2942dd'
const testWallet = new Wallet(testPrivateKey, local_provider)

describe("Turing VRF", function () {

  before(async () => {
    
    Factory__Turing = new ContractFactory(
      (HelloTuringJson.abi),
      (HelloTuringJson.bytecode),
      testWallet)
    
    turing = await Factory__Turing.deploy(helperPredeploy, gasOverride)
    console.log("    Test contract deployed at", turing.address)
  })

  it("should get the number 42", async () => {
    let tr = await turing.get42()
    const res = await tr.wait()
    expect(res).to.be.ok
    const rawData = res.events[0].data
    const numberHexString = rawData.slice(-64)
    let result = parseInt(numberHexString, 16)
    console.log("    Turing 42 =",result)
  })

  it("should get a length 256 VRF", async () => {
    let tr = await turing.getRandom()
    const res = await tr.wait()
    expect(res).to.be.ok
    const rawData = res.events[0].data
    const numberHexString = '0x'+ rawData.slice(-64)
    let result = BigInt(numberHexString)
    console.log("    Turing VRF 256 =",result)
  })

  it("should get a length 256 VRF", async () => {
    let tr = await turing.getRandom()
    const res = await tr.wait()
    expect(res).to.be.ok
    const rawData = res.events[0].data
    const numberHexString = '0x'+ rawData.slice(-64)
    let result = BigInt(numberHexString)
    console.log("    Turing VRF 256 =",result)
  })

})

