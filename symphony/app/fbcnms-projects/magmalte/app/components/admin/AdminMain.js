/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow strict-local
 * @format
 */

import * as React from 'react';
import AppContent from '@fbcnms/ui/components/layout/AppContent';
import AppContext from '@fbcnms/ui/context/AppContext';
import AppSideBar from '@fbcnms/ui/components/layout/AppSideBar';

import nullthrows from '@fbcnms/util/nullthrows';
import {getProjectLinks} from '@fbcnms/projects/projects';
import {makeStyles} from '@material-ui/styles';
import {useContext} from 'react';

const useStyles = makeStyles(_theme => ({
  root: {
    display: 'flex',
  },
}));

type Props = {
  navItems: () => React.Node,
  navRoutes: () => React.Node,
};

export default function AdminMain(props: Props) {
  const classes = useStyles();
  const {tabs, user, ssoEnabled} = useContext(AppContext);

  return (
    <div className={classes.root}>
      <AppSideBar
        mainItems={props.navItems()}
        projects={getProjectLinks(tabs, user)}
        user={nullthrows(user)}
        showSettings={!ssoEnabled}
      />
      <AppContent>{props.navRoutes()}</AppContent>
    </div>
  );
}
