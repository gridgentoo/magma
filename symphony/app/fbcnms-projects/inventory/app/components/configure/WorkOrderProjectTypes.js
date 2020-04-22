/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import type {WorkOrderProjectTypesQueryResponse} from './__generated__/WorkOrderProjectTypesQuery.graphql';

import AddEditProjectTypeCard from './AddEditProjectTypeCard';
import InventoryQueryRenderer from '../InventoryQueryRenderer';
import InventoryView from '../InventoryViewContainer';
import ProjectTypeCard from './ProjectTypeCard';
import React, {useState} from 'react';
import fbt from 'fbt';
import {LogEvents, ServerLogger} from '../../common/LoggingUtils';
import {graphql} from 'relay-runtime';
import {makeStyles} from '@material-ui/styles';

const useStyles = makeStyles(() => ({
  typeCards: {
    display: 'flex',
    flexWrap: 'wrap',
    flexDirection: 'row',
  },
  typeCard: {
    padding: '8px',
    flexBasis: '16.66%', // 6 cards
  },
  '@media (max-width: 1950px)': {
    typeCard: {
      flexBasis: '20%', // 5 cards
    },
  },
  '@media (max-width: 1600px)': {
    typeCard: {
      flexBasis: '25%', // 4 cards
    },
  },
  '@media (max-width: 1024px)': {
    typeCard: {
      flexBasis: '33.33%', // 3 cards
    },
  },
  '@media (max-width: 650px)': {
    typeCard: {
      flexBasis: '100%', // 1 card
    },
  },
}));

const projectTypesQuery = graphql`
  query WorkOrderProjectTypesQuery {
    projectTypes(first: 500)
      @connection(key: "WorkOrderProjectTypesQuery_projectTypes") {
      edges {
        node {
          id
          ...ProjectTypeCard_projectType
          ...AddEditProjectTypeCard_editingProjectType
        }
      }
    }
    workOrderTypes {
      edges {
        node {
          ...ProjectTypeWorkOrderTemplatesPanel_workOrderTypes
        }
      }
    }
  }
`;

const WorkOrderProjectTypes = () => {
  const classes = useStyles();
  const [editingProjectType, setEditingProjectType] = useState(null);
  const [showAddEditCard, setShowAddEditCard] = useState(false);
  const hideAddEditCard = () => {
    setEditingProjectType(null);
    setShowAddEditCard(false);
  };
  return (
    <InventoryQueryRenderer
      query={projectTypesQuery}
      variables={{}}
      render={(props: WorkOrderProjectTypesQueryResponse) => {
        if (showAddEditCard || editingProjectType) {
          const workOrderTypes = props.workOrderTypes?.edges ?? [];
          return (
            <AddEditProjectTypeCard
              workOrderTypes={workOrderTypes.map(e => e?.node).filter(Boolean)}
              editingProjectType={editingProjectType}
              onCancelClicked={hideAddEditCard}
              onProjectTypeSaved={hideAddEditCard}
            />
          );
        }

        return (
          <InventoryView
            header={{
              title: <fbt desc="">Project Templates</fbt>,
              subtitle: (
                <fbt desc="">Create and manage reusable project workflows.</fbt>
              ),
              actionButtons: [
                {
                  title: `${fbt('Create Project Template', '')}`,
                  action: () => {
                    ServerLogger.info(
                      LogEvents.ADD_PROJECT_TEMPLATE_BUTTON_CLICKED,
                    );
                    setShowAddEditCard(true);
                  },
                },
              ],
            }}>
            <div className={classes.typeCards}>
              {(props.projectTypes?.edges ?? [])
                .map(edge => edge.node)
                .filter(Boolean)
                .map(projectType => (
                  <div key={projectType.id} className={classes.typeCard}>
                    <ProjectTypeCard
                      projectType={projectType}
                      onEditClicked={() => setEditingProjectType(projectType)}
                    />
                  </div>
                ))}
            </div>
          </InventoryView>
        );
      }}
    />
  );
};

export default WorkOrderProjectTypes;
