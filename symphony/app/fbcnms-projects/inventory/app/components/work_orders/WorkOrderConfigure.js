/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow strict-local
 * @format
 */

import NavigatableViews from '@fbcnms/ui/components/design-system/View/NavigatableViews';
import React, {useMemo} from 'react';
import WorkOrderProjectTypes from '../configure/WorkOrderProjectTypes';
import WorkOrderTypes from '../configure/WorkOrderTypes';
import fbt from 'fbt';
import {makeStyles} from '@material-ui/styles';

const useStyles = makeStyles(() => ({
  root: {
    display: 'flex',
    height: '100vh',
    transform: 'translateZ(0)',
  },
}));

export default function WorkOrderConfigure() {
  const menuItems = useMemo(
    () => [
      {
        menuItem: {
          label: `${fbt('Work Orders', '')}`,
          tooltip: '',
        },
        component: {
          children: <WorkOrderTypes />,
        },
        routingPath: 'work_order_types',
      },
      {
        menuItem: {
          label: `${fbt('Projects', '')}`,
          tooltip: '',
        },
        component: {
          children: <WorkOrderProjectTypes />,
        },
        routingPath: 'project_types',
      },
    ],
    [],
  );

  const classes = useStyles();
  return (
    <div className={classes.root}>
      <NavigatableViews
        header={<fbt desc="">Templates</fbt>}
        views={menuItems}
        routingBasePath="/workorders/configure"
      />
    </div>
  );
}
