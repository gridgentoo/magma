/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow strict-local
 * @format
 *
 * Base component for rule editors to render. Handles rendering common elements
 * such as receiver config and label editor.
 */

import * as React from 'react';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import CardHeader from '@material-ui/core/CardHeader';
import Editor from '../common/Editor';
import Grid from '@material-ui/core/Grid';
import LabelsEditor from './LabelsEditor';
import RuleContext from './RuleContext';
import SelectReceiver from '../alertmanager/Receivers/SelectReceiver';
import SelectRuleType from './SelectRuleType';
import TextField from '@material-ui/core/TextField';
import Typography from '@material-ui/core/Typography';
import {useAlarmContext} from '../AlarmContext';
import {useAlertRuleReceiver} from '../hooks';
import {useForm} from '@fbcnms/ui/hooks';

import type {Props as EditorProps} from '../common/Editor';
import type {Labels} from '../AlarmAPIType';

type Props = EditorProps & {
  onChange: (form: RuleEditorBaseFields) => void,
  initialState: ?RuleEditorBaseFields,
};

// Fields for inputs which are standard between different rule editors
export type RuleEditorBaseFields = {
  name: string,
  description: string,
  labels: Labels,
};

export default function RuleEditorBase({
  isNew,
  children,
  initialState,
  onChange,
  onSave,
  ...props
}: Props) {
  const {apiUtil} = useAlarmContext();
  const ruleContext = React.useContext(RuleContext);
  const {formState, handleInputChange, updateFormState} = useForm({
    initialState: initialState || defaultState(),
    onFormUpdated: onChange,
  });
  const {receiver, setReceiver, saveReceiver} = useAlertRuleReceiver({
    ruleName: formState?.name || '',
    apiUtil,
  });

  const handleSave = React.useCallback(async () => {
    await onSave();
    await saveReceiver();
  }, [saveReceiver, onSave]);

  const handleLabelsChange = React.useCallback(
    (labels: Labels) => {
      updateFormState({
        labels,
      });
    },
    [updateFormState],
  );

  return (
    <Editor
      {...props}
      title={
        formState?.name == null || formState?.name.trim() === ''
          ? 'New Alert Rule'
          : formState.name
      }
      description="Configure rules to fire alerts"
      isNew={isNew}
      onSave={handleSave}>
      <Grid container item spacing={4}>
        <Grid container direction="column" item xs={7} spacing={4}>
          <Grid item>
            <Card>
              <CardHeader title="Details" />
              <CardContent>
                <Grid container direction="column" spacing={2}>
                  {isNew && (
                    <Grid item xs={6}>
                      <SelectRuleType
                        ruleMap={ruleContext.ruleMap}
                        value={ruleContext.ruleType}
                        onChange={ruleContext.selectRuleType}
                      />
                    </Grid>
                  )}
                  {children}
                </Grid>
              </CardContent>
            </Card>
          </Grid>
          <Grid item>
            <Card>
              <CardHeader title="Summary" />
              <CardContent>
                <Grid item>
                  <TextField
                    disabled={!isNew}
                    required
                    label="Rule Name"
                    placeholder="Ex: Link down"
                    fullWidth
                    value={formState.name}
                    onChange={handleInputChange(val => ({name: val}))}
                  />
                </Grid>
                <Grid item>
                  <TextField
                    disabled={!isNew}
                    required
                    label="Description"
                    placeholder="Ex: The link is down"
                    fullWidth
                    value={formState.description}
                    onChange={handleInputChange(val => ({description: val}))}
                  />
                </Grid>
              </CardContent>
            </Card>
          </Grid>
        </Grid>
        <Grid container direction="column" item spacing={4} xs={5}>
          <Grid item>
            <Card>
              <CardHeader
                title={
                  <>
                    <Typography variant="h5" gutterBottom>
                      Notify
                    </Typography>
                    <Typography
                      color="textSecondary"
                      gutterBottom
                      variant="body2">
                      Team or service to notify when this alert fires
                    </Typography>
                  </>
                }
              />
              <CardContent>
                <SelectReceiver
                  label="Send Notification To:"
                  fullWidth
                  receiver={receiver}
                  onChange={setReceiver}
                />
              </CardContent>
            </Card>
          </Grid>
          <Grid item>
            <LabelsEditor
              labels={formState.labels}
              onChange={handleLabelsChange}
            />
          </Grid>
        </Grid>
      </Grid>
    </Editor>
  );
}

function defaultState(): RuleEditorBaseFields {
  return {
    name: '',
    description: '',
    labels: {},
  };
}
