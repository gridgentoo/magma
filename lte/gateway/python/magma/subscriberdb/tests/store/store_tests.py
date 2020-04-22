"""
Copyright (c) 2016-present, Facebook, Inc.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree. An additional grant
of patent rights can be found in the PATENTS file in the same directory.
"""

import unittest

from lte.protos.subscriberdb_pb2 import SubscriberData
from magma.subscriberdb.store.base import DuplicateSubscriberError, \
    SubscriberNotFoundError
from magma.subscriberdb.store.sqlite import SqliteStore

from magma.subscriberdb.sid import SIDUtils


class StoreTests(unittest.TestCase):
    """
    Test class for subscriber storage
    """

    def setUp(self):
        # Create an in-memory sqlite3 database for testing
        self._store = SqliteStore("file::memory:")

    def _add_subscriber(self, sid):
        sub = SubscriberData(sid=SIDUtils.to_pb(sid))
        self._store.add_subscriber(sub)
        return (sid, sub)

    def test_subscriber_addition(self):
        """
        Test if subscriber addition works as expected
        """
        self.assertEqual(self._store.list_subscribers(), [])
        (sid1, _) = self._add_subscriber('IMSI11111')
        self.assertEqual(self._store.list_subscribers(), [sid1])
        (sid2, sub2) = self._add_subscriber('IMSI22222')
        self.assertEqual(self._store.list_subscribers(), [sid1, sid2])

        # Check if adding an existing user throws an exception
        with self.assertRaises(DuplicateSubscriberError):
            self._store.add_subscriber(sub2)
        self.assertEqual(self._store.list_subscribers(), [sid1, sid2])

        self._store.delete_all_subscribers()
        self.assertEqual(self._store.list_subscribers(), [])

    def test_subscriber_deletion(self):
        """
        Test if subscriber deletion works as expected
        """
        (sid1, _) = self._add_subscriber('IMSI11111')
        (sid2, _) = self._add_subscriber('IMSI22222')
        self.assertEqual(self._store.list_subscribers(), [sid1, sid2])

        self._store.delete_subscriber(sid2)
        self.assertEqual(self._store.list_subscribers(), [sid1])

        # Deleting a non-existent user would be ignored
        self._store.delete_subscriber(sid2)
        self.assertEqual(self._store.list_subscribers(), [sid1])

        self._store.delete_subscriber(sid1)
        self.assertEqual(self._store.list_subscribers(), [])

    def test_subscriber_retrieval(self):
        """
        Test if subscriber retrieval works as expected
        """
        (sid1, sub1) = self._add_subscriber('IMSI11111')
        self.assertEqual(self._store.list_subscribers(), [sid1])
        self.assertEqual(self._store.get_subscriber_data(sid1), sub1)

        with self.assertRaises(SubscriberNotFoundError):
            self._store.get_subscriber_data('IMSI30000')

        self._store.delete_all_subscribers()
        self.assertEqual(self._store.list_subscribers(), [])

    def test_subscriber_edit(self):
        """
        Test if subscriber edit works as expected
        """
        (sid1, sub1) = self._add_subscriber('IMSI11111')
        self.assertEqual(self._store.get_subscriber_data(sid1), sub1)

        sub1.lte.auth_key = b'1234'
        self._store.update_subscriber(sub1)
        self.assertEqual(self._store.get_subscriber_data(sid1).lte.auth_key,
                         b'1234')

        with self._store.edit_subscriber(sid1) as subs:
            subs.lte.auth_key = b'5678'
        self.assertEqual(self._store.get_subscriber_data(sid1).lte.auth_key,
                         b'5678')

        with self.assertRaises(SubscriberNotFoundError):
            sub1.sid.id = '30000'
            self._store.update_subscriber(sub1)
        with self.assertRaises(SubscriberNotFoundError):
            with self._store.edit_subscriber('IMSI3000') as subs:
                pass


if __name__ == "__main__":
    unittest.main()
