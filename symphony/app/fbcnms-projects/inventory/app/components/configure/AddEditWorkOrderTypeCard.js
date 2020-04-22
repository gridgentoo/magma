/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import type {AddEditWorkOrderTypeCard_editingWorkOrderType} from './__generated__/AddEditWorkOrderTypeCard_editingWorkOrderType.graphql';
import type {
  AddWorkOrderTypeMutationResponse,
  AddWorkOrderTypeMutationVariables,
} from '../../mutations/__generated__/AddWorkOrderTypeMutation.graphql';
import type {
  EditWorkOrderTypeMutationResponse,
  EditWorkOrderTypeMutationVariables,
} from '../../mutations/__generated__/EditWorkOrderTypeMutation.graphql';
import type {MutationCallbacks} from '../../mutations/MutationCallbacks.js';
import type {WithAlert} from '@fbcnms/ui/components/Alert/withAlert';
import type {WithSnackbarProps} from 'notistack';
import type {WithStyles} from '@material-ui/core';
import type {WorkOrderType} from '../../common/WorkOrder';

import AddWorkOrderTypeMutation from '../../mutations/AddWorkOrderTypeMutation';
import Breadcrumbs from '@fbcnms/ui/components/Breadcrumbs';
import Button from '@fbcnms/ui/components/design-system/Button';
import DeleteOutlineIcon from '@material-ui/icons/DeleteOutline';
import EditWorkOrderTypeMutation from '../../mutations/EditWorkOrderTypeMutation';
import ExpandingPanel from '@fbcnms/ui/components/ExpandingPanel';
import FormAction from '@fbcnms/ui/components/design-system/Form/FormAction';
import NameDescriptionSection from '@fbcnms/ui/components/NameDescriptionSection';
import PropertyTypeTable from '../form/PropertyTypeTable';
import React from 'react';
import RemoveWorkOrderTypeMutation from '../../mutations/RemoveWorkOrderTypeMutation';
import SnackbarItem from '@fbcnms/ui/components/SnackbarItem';
import fbt from 'fbt';
import symphony from '@fbcnms/ui/theme/symphony';
import update from 'immutability-helper';
import withAlert from '@fbcnms/ui/components/Alert/withAlert';
import {ConnectionHandler} from 'relay-runtime';
import {FormContextProvider} from '../../common/FormContext';
import {createFragmentContainer, graphql} from 'react-relay';
import {getGraphError} from '../../common/EntUtils';
import {getPropertyDefaultValue} from '../../common/PropertyType';
import {sortByIndex} from '../draggable/DraggableUtils';
import {withSnackbar} from 'notistack';
import {withStyles} from '@material-ui/core/styles';

const styles = {
  root: {
    padding: '24px 16px',
    maxHeight: '100%',
    overflow: 'hidden',
    display: 'flex',
    flexDirection: 'column',
  },
  header: {
    display: 'flex',
    paddingBottom: '24px',
  },
  body: {
    overflowY: 'auto',
  },
  buttons: {
    display: 'flex',
  },
  cancelButton: {
    marginRight: '8px',
  },
  deleteButton: {
    cursor: 'pointer',
    color: symphony.palette.D400,
    width: '32px',
    height: '32px',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    marginRight: '8px',
  },
};

type Props = WithSnackbarProps &
  WithStyles<typeof styles> &
  WithAlert & {
    open: boolean,
    onClose: () => void,
    onSave: (workOrderType: any) => void,
    editingWorkOrderType?: AddEditWorkOrderTypeCard_editingWorkOrderType,
  };

type State = {
  editingWorkOrderType: WorkOrderType,
  isSaving: boolean,
};

class AddEditWorkOrderTypeCard extends React.Component<Props, State> {
  state = {
    editingWorkOrderType: this.getEditingWorkOrderType(),
    isSaving: false,
  };

  render() {
    const {classes, onClose} = this.props;
    const {editingWorkOrderType} = this.state;
    const propertyTypes = editingWorkOrderType.propertyTypes
      .slice()
      .sort(sortByIndex);
    return (
      <FormContextProvider>
        <div className={classes.root}>
          <div className={classes.header}>
            <Breadcrumbs
              breadcrumbs={[
                {
                  id: 'wo_templates',
                  name: 'Work Order Templates',
                  onClick: onClose,
                },
                this.props.editingWorkOrderType
                  ? {
                      id: this.props.editingWorkOrderType.id,
                      name: this.props.editingWorkOrderType.name,
                    }
                  : {
                      id: 'new_wo_type',
                      name: `${fbt('New work order template', '')}`,
                    },
              ]}
              size="large"
            />
            <div className={classes.buttons}>
              {!!this.props.editingWorkOrderType && (
                <FormAction>
                  <Button
                    className={classes.deleteButton}
                    variant="text"
                    skin="gray"
                    onClick={this.onDelete}>
                    <DeleteOutlineIcon />
                  </Button>
                </FormAction>
              )}
              <Button
                className={classes.cancelButton}
                skin="regular"
                onClick={onClose}>
                Cancel
              </Button>
              <FormAction>
                <Button disabled={this.isSaveDisabled()} onClick={this.onSave}>
                  Save
                </Button>
              </FormAction>
            </div>
          </div>
          <div className={classes.body}>
            <ExpandingPanel title="Details">
              <NameDescriptionSection
                title="Title"
                name={editingWorkOrderType.name ?? ''}
                namePlaceholder={`${fbt('New work order template', '')}`}
                description={editingWorkOrderType.description ?? ''}
                descriptionPlaceholder={`${fbt(
                  'Write a description if you want it to appear whenever this template of work order is created',
                  '',
                )}`}
                onNameChange={this.nameChanged}
                onDescriptionChange={this.descriptionChanged}
              />
            </ExpandingPanel>
            <ExpandingPanel title="Properties">
              <PropertyTypeTable
                supportDelete={true}
                propertyTypes={propertyTypes}
                onPropertiesChanged={this._propertyChangedHandler}
              />
            </ExpandingPanel>
          </div>
        </div>
      </FormContextProvider>
    );
  }

  isSaveDisabled() {
    return (
      !this.state.editingWorkOrderType.name ||
      this.state.isSaving ||
      !this.state.editingWorkOrderType.propertyTypes.every(propType => {
        return (
          propType.isInstanceProperty || !!getPropertyDefaultValue(propType)
        );
      })
    );
  }

  deleteTempId = (definition: Object) => {
    const newDef = {...definition};
    if (definition.id && definition.id.includes('@tmp')) {
      newDef['id'] = undefined;
    }
    return newDef;
  };

  onSave = () => {
    const {name} = this.state.editingWorkOrderType;
    if (!name) {
      const msg = 'Name cannot be empty';
      this.props.enqueueSnackbar(msg, {
        children: key => (
          <SnackbarItem id={key} message={msg} variant="error" />
        ),
      });
      return;
    }
    this.setState({isSaving: true});
    if (this.props.editingWorkOrderType) {
      this.editWorkOrderType();
    } else {
      this.addNewWorkOrderType();
    }
  };

  _onError = (error: Error) => {
    this._showError(getGraphError(error));
    this.setState({isSaving: false});
  };

  editWorkOrderType = () => {
    const {
      id,
      name,
      description,
      propertyTypes,
    } = this.state.editingWorkOrderType;
    const variables: EditWorkOrderTypeMutationVariables = {
      input: {
        id,
        name,
        description,
        properties: propertyTypes
          .filter(propType => !!propType.name)
          .map(this.deleteTempId),
      },
    };
    const callbacks: MutationCallbacks<EditWorkOrderTypeMutationResponse> = {
      onCompleted: (response, errors) => {
        this.setState({isSaving: false});
        if (errors && errors[0]) {
          this._showError(errors[0].message);
        } else {
          this.props.onSave && this.props.onSave(response.editWorkOrderType);
        }
      },
      onError: this._onError,
    };
    EditWorkOrderTypeMutation(variables, callbacks);
  };

  _showError = (msg: string) => {
    this.props.enqueueSnackbar(msg, {
      children: key => <SnackbarItem id={key} message={msg} variant="error" />,
    });
  };

  addNewWorkOrderType = () => {
    const {name, description, propertyTypes} = this.state.editingWorkOrderType;
    const variables: AddWorkOrderTypeMutationVariables = {
      input: {
        name,
        description,
        properties: propertyTypes
          .filter(propType => !!propType.name)
          .map(this.deleteTempId),
      },
    };

    const callbacks: MutationCallbacks<AddWorkOrderTypeMutationResponse> = {
      onCompleted: (response, errors) => {
        this.setState({isSaving: false});
        if (errors && errors[0]) {
          this._showError(errors[0].message);
        } else {
          this.props.onSave && this.props.onSave(response.addWorkOrderType);
        }
      },
      onError: this._onError,
    };
    const updater = store => {
      // $FlowFixMe (T62907961) Relay flow types
      const rootQuery = store.getRoot();
      // $FlowFixMe (T62907961) Relay flow types
      const newNode = store.getRootField('addWorkOrderType');
      if (!newNode) {
        return;
      }
      const types = ConnectionHandler.getConnection(
        rootQuery,
        'Configure_workOrderTypes',
      );
      const edge = ConnectionHandler.createEdge(
        // $FlowFixMe (T62907961) Relay flow types
        store,
        // $FlowFixMe (T62907961) Relay flow types
        types,
        newNode,
        'WorkOrderTypesEdge',
      );
      // $FlowFixMe (T62907961) Relay flow types
      ConnectionHandler.insertEdgeBefore(types, edge);
    };
    AddWorkOrderTypeMutation(variables, callbacks, updater);
  };

  onDelete = () => {
    const {editingWorkOrderType, confirm, onClose} = this.props;
    if (!editingWorkOrderType) {
      return;
    }

    confirm(
      `Are you sure you want to delete "${editingWorkOrderType.name}"?`,
    ).then(confirm => {
      if (!confirm) {
        return;
      }
      RemoveWorkOrderTypeMutation(
        {id: editingWorkOrderType.id},
        {
          onCompleted: onClose,
          onError: (error: Error) => {
            this.props.alert(`Error: ${error.message}`);
          },
        },
        store => {
          // $FlowFixMe (T62907961) Relay flow types
          const rootQuery = store.getRoot();
          const workOrderTypes = ConnectionHandler.getConnection(
            rootQuery,
            'Configure_workOrderTypes',
          );
          // $FlowFixMe (T62907961) Relay flow types
          ConnectionHandler.deleteNode(workOrderTypes, editingWorkOrderType.id);
          // $FlowFixMe (T62907961) Relay flow types
          store.delete(editingWorkOrderType.id);
        },
      );
    });
  };

  fieldChangedHandler = (field: 'name' | 'description') => value =>
    this.setState({
      editingWorkOrderType: {
        ...this.state.editingWorkOrderType,
        // $FlowFixMe Set state for each field
        [field]: value,
      },
    });

  nameChanged = this.fieldChangedHandler('name');
  descriptionChanged = this.fieldChangedHandler('description');

  _propertyChangedHandler = properties => {
    this.setState(prevState => {
      return {
        editingWorkOrderType: update(prevState.editingWorkOrderType, {
          propertyTypes: {$set: properties},
        }),
      };
    });
  };

  getEditingWorkOrderType(): WorkOrderType {
    const editingWorkOrderType = this.props.editingWorkOrderType;
    const propertyTypes = (editingWorkOrderType?.propertyTypes ?? [])
      .filter(Boolean)
      .filter(propertyType => !propertyType.isDeleted)
      .map(p => ({
        id: p.id,
        name: p.name,
        index: p.index || 0,
        type: p.type,
        nodeType: p.nodeType,
        booleanValue: p.booleanValue,
        stringValue: p.stringValue,
        intValue: p.intValue,
        floatValue: p.floatValue,
        latitudeValue: p.latitudeValue,
        longitudeValue: p.longitudeValue,
        isEditable: p.isEditable,
        isMandatory: p.isMandatory,
        isInstanceProperty: p.isInstanceProperty,
        isDeleted: p.isDeleted,
      }));

    return {
      id: editingWorkOrderType?.id ?? 'WorkOrderType@tmp-0',
      name: editingWorkOrderType?.name ?? '',
      description: editingWorkOrderType?.description,
      numberOfWorkOrders: editingWorkOrderType?.numberOfWorkOrders ?? 0,
      propertyTypes:
        propertyTypes.length > 0
          ? propertyTypes
          : [
              {
                id: 'PropertyType@tmp',
                name: '',
                index: editingWorkOrderType?.propertyTypes?.length ?? 0,
                type: 'string',
                nodeType: null,
                booleanValue: false,
                stringValue: null,
                intValue: null,
                floatValue: null,
                latitudeValue: null,
                longitudeValue: null,
                isEditable: true,
                isMandatory: false,
                isInstanceProperty: true,
                isDeleted: false,
              },
            ],
    };
  }
}

export default withStyles(styles)(
  withAlert(
    withSnackbar(
      createFragmentContainer(AddEditWorkOrderTypeCard, {
        editingWorkOrderType: graphql`
          fragment AddEditWorkOrderTypeCard_editingWorkOrderType on WorkOrderType {
            id
            name
            description
            numberOfWorkOrders
            propertyTypes {
              id
              name
              type
              nodeType
              index
              stringValue
              intValue
              booleanValue
              floatValue
              latitudeValue
              longitudeValue
              rangeFromValue
              rangeToValue
              isEditable
              isMandatory
              isInstanceProperty
              isDeleted
            }
          }
        `,
      }),
    ),
  ),
);
