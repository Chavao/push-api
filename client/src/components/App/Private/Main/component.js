import React from 'react'
import { Redirect, Route, Switch } from 'react-router-dom'

import Container from '@material-ui/core/Container'

import { privateStatsPath, privateChannelsPathWithOptionalParam } from 'navigation'
import { useStyles } from 'components/App/Private/styles'

import Stats from './views/Stats'
import Channels from './views/Channels'

const Main = (props) => {
  const classes = useStyles()
  return (
    <main className={classes.content}>
      <div className={classes.appBarSpacer} />
      <Container maxWidth="lg" className={classes.container}>
        <Switch>
          <Route path={privateStatsPath} exact component={Stats} />
          <Route path={privateChannelsPathWithOptionalParam} component={Channels} />
          <Redirect to={privateStatsPath} />
        </Switch>
      </Container>
    </main>
  )
}

export default Main
