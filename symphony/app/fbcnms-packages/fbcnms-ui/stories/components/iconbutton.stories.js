/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow strict-local
 * @format
 */

import IconButton from '../../components/design-system/IconButton';
import React from 'react';
import {AddIcon} from '../../components/design-system/Icons';
import {STORY_CATEGORIES} from '../storybookUtils';
import {storiesOf} from '@storybook/react';

const IconButtonsRoot = () => {
  const onClick = () => window.alert('Clicked!');
  return (
    <div>
      <IconButton icon={AddIcon} onClick={onClick} />
      <IconButton icon={AddIcon} skin="gray" onClick={onClick} />
      <IconButton icon={AddIcon} disabled={true} onClick={onClick} />
    </div>
  );
};

storiesOf(`${STORY_CATEGORIES.COMPONENTS}`, module).add('IconButton', () => (
  <IconButtonsRoot />
));
