/*
Copyright 2019-present OmiseGO Pte Ltd

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License. */

import React from 'react'

import { Typography } from '@material-ui/core'

import * as styles from './Transaction.module.scss'
import * as S from './Transaction.styles'
import { useTheme } from '@emotion/react'
import { selectNetwork } from 'selectors/setupSelector'
import { useSelector } from 'react-redux'
import { getAllNetworks } from 'util/masterConfig'

function FastExit({
  link,
  status,
  chain,
  blockNumber,
  oriChain,
  oriHash,
  age,
  unixTime
}) {

  const theme = useTheme()

  const currentNetwork = useSelector(selectNetwork())
  const nw = getAllNetworks()

  const chainLink = ({hash}) => {
    let network = nw[currentNetwork]
    if (!!network && !!network['L2']) {
          return `${network['L2'].transaction}${hash}`
    }
    return ''
  }

  const secondsAgo = Math.round(Date.now() / 1000) - unixTime
  let timeLabel = `Fast Exit was started ${secondsAgo} seconds ago`

  return (
    <
      div style={{
        padding: '5px',
        borderRadius: '8px',
        background: theme.palette.background.secondaryLight,
      }}
    >
      <S.TableBody>

        <S.TableCell
          style={{ width: '100%' }}
        >

          <Typography variant="body3">
            {blockNumber}
          </Typography>

          <Typography variant="body3" className={styles.muted} style={{color: 'green'}}>
            {timeLabel}
          </Typography>

          <Typography variant="body3" className={styles.muted}>
            Hash:&nbsp;
            <a
              href={chainLink({hash:oriHash})}
              target={'_blank'}
              rel='noopener noreferrer'
              style={{ color: theme.palette.mode === 'light' ? 'black' : 'white' }}
            >
              {oriHash}
            </a>
          </Typography>

        </S.TableCell>

      </S.TableBody>

    </div>)

}

export default FastExit
