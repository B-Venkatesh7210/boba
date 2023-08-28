import { DeployFunction } from 'hardhat-deploy/dist/types'

import {
  assertContractVariable,
  deploy,
  getContractFromArtifact,
} from '../scripts/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const OptimismPortalProxy = await getContractFromArtifact(
    hre,
    'OptimismPortalProxy'
  )

  await deploy({
    hre,
    name: 'L1CrossDomainMessenger',
    args: [],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'PORTAL',
        OptimismPortalProxy.address
      )
    },
  })
}

deployFn.tags = ['L1CrossDomainMessengerImpl', 'setup', 'l1']

export default deployFn
