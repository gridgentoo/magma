/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import type {ProjectTypeWorkOrderTemplatesPanel_workOrderTypes} from './__generated__/ProjectTypeWorkOrderTemplatesPanel_workOrderTypes.graphql';

import AddCircleOutlineIcon from '@material-ui/icons/AddCircleOutline';
import Button from '@fbcnms/ui/components/design-system/Button';
import ExpandingPanel from '@fbcnms/ui/components/ExpandingPanel';
import FormAction from '@fbcnms/ui/components/design-system/Form/FormAction';
import ProjectTypeSelectWorkOrdersDialog from './ProjectTypeSelectWorkOrdersDialog';
import React, {useState} from 'react';
import Text from '@fbcnms/ui/components/design-system/Text';
import nullthrows from 'nullthrows';
import {createFragmentContainer, graphql} from 'react-relay';
import {makeStyles} from '@material-ui/styles';

const useStyles = makeStyles(theme => ({
  addButton: {
    marginRight: '8px',
  },
  addIcon: {
    verticalAlign: 'middle',
  },
  item: {
    fontSize: '16px',
    fontWeight: 500,
    lineHeight: '24px',
    color: theme.palette.blueGrayDark,
    paddingBottom: '16px',
  },
}));

type Props = {
  selectedWorkOrderTypeIds: Array<string>,
  workOrderTypes: ProjectTypeWorkOrderTemplatesPanel_workOrderTypes,
  onWorkOrderTypesSelected: (ids: Array<string>) => void,
};

const ProjectTypeWorkOrderTemplatesPanel = ({
  selectedWorkOrderTypeIds,
  workOrderTypes,
  onWorkOrderTypesSelected,
}: Props) => {
  const classes = useStyles();
  const [
    selectWorkOrdersDialogShown,
    setSelectWorkOrdersDialogShown,
  ] = useState(false);

  return (
    <>
      <ExpandingPanel
        title="Work Orders"
        rightContent={
          <FormAction>
            <Button
              className={classes.addButton}
              variant="text"
              skin="primary"
              onClick={() => setSelectWorkOrdersDialogShown(true)}>
              <AddCircleOutlineIcon className={classes.addIcon} />
            </Button>
          </FormAction>
        }>
        {selectedWorkOrderTypeIds
          .map(id => nullthrows(workOrderTypes.find(type => type.id === id)))
          .map(type => (
            <Text key={type.id} className={classes.item}>
              {type.name}
            </Text>
          ))}
      </ExpandingPanel>
      <ProjectTypeSelectWorkOrdersDialog
        initialSelectedWorkOrderTypeIds={selectedWorkOrderTypeIds}
        open={selectWorkOrdersDialogShown}
        onSaveClicked={workOrderTypeIds => {
          setSelectWorkOrdersDialogShown(false);
          onWorkOrderTypesSelected(workOrderTypeIds);
        }}
        onClose={() => setSelectWorkOrdersDialogShown(false)}
        workOrderTypes={workOrderTypes}
      />
    </>
  );
};

export default createFragmentContainer(ProjectTypeWorkOrderTemplatesPanel, {
  workOrderTypes: graphql`
    fragment ProjectTypeWorkOrderTemplatesPanel_workOrderTypes on WorkOrderType
      @relay(plural: true) {
      id
      name
    }
  `,
});
