/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow strict-local
 * @format
 */

import type {GroupMember} from '../utils/GroupMemberViewer';
import type {UserPermissionsGroup} from '../utils/UserManagementUtils';

import * as React from 'react';
import Button from '@fbcnms/ui/components/design-system/Button';
import MembersList from '../utils/MembersList';
import Text from '@fbcnms/ui/components/design-system/Text';
import fbt from 'fbt';
import symphony from '@fbcnms/ui/theme/symphony';
import {ASSIGNMENT_BUTTON_VIEWS} from '../utils/GroupMemberViewer';
import {makeStyles} from '@material-ui/styles';
import {useMemo} from 'react';
import {useUserSearchContext} from '../utils/userSearch/UserSearchContext';

const useStyles = makeStyles(() => ({
  root: {},
  noMembers: {
    width: '124px',
    paddingTop: '50%',
    alignSelf: 'center',
  },
  noSearchResults: {
    paddingTop: '50%',
    alignSelf: 'center',
    textAlign: 'center',
  },
  clearSearchWrapper: {
    marginTop: '16px',
  },
  clearSearch: {
    ...symphony.typography.subtitle1,
  },
}));

type Props = $ReadOnly<{|
  group: UserPermissionsGroup,
|}>;

export default function PermissionsGroupMembersList(props: Props) {
  const {group} = props;
  const classes = useStyles();
  const userSearch = useUserSearchContext();

  const groupMembers: $ReadOnlyArray<GroupMember> = useMemo(
    () =>
      group.memberUsers.map(user => ({
        user: user,
        isMember: true,
      })),
    [group.memberUsers],
  );

  const groupMembersEmptyState = useMemo(
    () => (
      <img
        className={classes.noMembers}
        src="/inventory/static/images/noMembers.png"
      />
    ),
    [classes.noMembers],
  );

  const emptyState = useMemo(() => {
    if (userSearch.isEmptySearchTerm) {
      return groupMembersEmptyState;
    }

    if (userSearch.isSearchInProgress) {
      return null;
    }

    return (
      <div className={classes.noSearchResults}>
        <Text variant="h6" color="gray">
          <fbt desc="">
            No users found for '<fbt:param name="given search term">
              {userSearch.searchTerm}
            </fbt:param>'
          </fbt>
        </Text>
        <div className={classes.clearSearchWrapper}>
          <Button variant="text" skin="gray" onClick={userSearch.clearSearch}>
            <span className={classes.clearSearch}>
              <fbt desc="">Clear Search</fbt>
            </span>
          </Button>
        </div>
      </div>
    );
  }, [
    classes.clearSearch,
    classes.clearSearchWrapper,
    classes.noSearchResults,
    groupMembersEmptyState,
    userSearch.clearSearch,
    userSearch.isEmptySearchTerm,
    userSearch.isSearchInProgress,
    userSearch.searchTerm,
  ]);

  return (
    <MembersList
      members={userSearch.isEmptySearchTerm ? groupMembers : userSearch.results}
      group={group}
      assigmentButton={
        userSearch.isEmptySearchTerm
          ? ASSIGNMENT_BUTTON_VIEWS.onHover
          : ASSIGNMENT_BUTTON_VIEWS.always
      }
      emptyState={emptyState}
    />
  );
}
