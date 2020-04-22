/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import type {
  LocationDocumentsCard_location$data,
  LocationDocumentsCard_location$key,
} from './__generated__/LocationDocumentsCard_location.graphql';

import AddHyperlinkButton from '../AddHyperlinkButton';
import Card from '@fbcnms/ui/components/design-system/Card/Card';
import CardHeader from '@fbcnms/ui/components/design-system/Card/CardHeader';
import DocumentsAddButton from '../DocumentsAddButton';
import EntityDocumentsTable from '../EntityDocumentsTable';
import InventoryStrings from '../../common/InventoryStrings';
import React, {useMemo} from 'react';
import classNames from 'classnames';
import {graphql, useFragment} from 'react-relay/hooks';
import {makeStyles} from '@material-ui/styles';

const useStyles = makeStyles(_theme => ({
  cardHasNoContent: {
    marginBottom: '0px',
  },
  actionButtonsContainer: {
    '&>*': {
      marginLeft: '8px',
    },
  },
}));

type Props = {
  className?: string,
  location: LocationDocumentsCard_location$key,
};

const LocationDocumentsCard = (props: Props) => {
  const {className, location} = props;
  const classes = useStyles();

  const data: LocationDocumentsCard_location$data = useFragment(
    graphql`
      fragment LocationDocumentsCard_location on Location {
        id
        images {
          ...EntityDocumentsTable_files
        }
        files {
          ...EntityDocumentsTable_files
        }
        hyperlinks {
          ...EntityDocumentsTable_hyperlinks
        }
      }
    `,
    location,
  );

  const documents = useMemo(
    () => [...data.files.filter(Boolean), ...data.images.filter(Boolean)],
    [data],
  );

  return (
    <Card className={className}>
      <CardHeader
        className={classNames({
          [classes.cardHasNoContent]: documents.length === 0,
        })}
        rightContent={
          <div className={classes.actionButtonsContainer}>
            <AddHyperlinkButton entityType="LOCATION" entityId={data.id} />
            <DocumentsAddButton entityType="LOCATION" entityId={data.id} />
          </div>
        }>
        {InventoryStrings.documents.viewHeader}
      </CardHeader>
      <EntityDocumentsTable
        entityType="LOCATION"
        entityId={data.id}
        files={documents}
        hyperlinks={data.hyperlinks}
      />
    </Card>
  );
};

export default LocationDocumentsCard;
