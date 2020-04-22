"""
Copyright (c) 2016-present, Facebook, Inc.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree. An additional grant
of patent rights can be found in the PATENTS file in the same directory.
"""

import logging
from logging.handlers import RotatingFileHandler


LOG_FILE = 'var/log/enodebd.log'
LOGGER_NAME = 'debug'
MAX_BYTES = 1024 * 1024 * 10  # 10MB
BACKUP_COUNT = 5  # 10MB, 5 files, 50MB total


class EnodebdLogger:
    """
    EnodebdLogger backs up debug logs with a RotatingFileHandler.

    Debug logs will be propagated to root level if the root logger is set to
    debug level.
    """

    _LOGGER = logging.getLogger(LOGGER_NAME)  # type: logging.Logger

    @staticmethod
    def init() -> None:
        if logging.root.level is not logging.DEBUG:
            EnodebdLogger._LOGGER.propagate = False
        handler = RotatingFileHandler(LOG_FILE,
                                      maxBytes=MAX_BYTES,
                                      backupCount=BACKUP_COUNT)
        formatter = logging.Formatter(fmt='%(asctime)s %(levelname)s %(message)s')
        handler.setFormatter(formatter)
        EnodebdLogger._LOGGER.addHandler(handler)
        EnodebdLogger._LOGGER.setLevel(logging.DEBUG)

    @staticmethod
    def debug(msg, *args, **kwargs):
        EnodebdLogger._LOGGER.debug(msg, *args, **kwargs)

    @staticmethod
    def info(msg, *args, **kwargs):
        if not EnodebdLogger._LOGGER.propagate:
            logging.info(msg, *args, **kwargs)
        EnodebdLogger._LOGGER.info(msg, *args, **kwargs)

    @staticmethod
    def warning(msg, *args, **kwargs):
        if not EnodebdLogger._LOGGER.propagate:
            logging.warning(msg, *args, **kwargs)
        EnodebdLogger._LOGGER.warning(msg, *args, **kwargs)

    @staticmethod
    def error(msg, *args, **kwargs):
        if not EnodebdLogger._LOGGER.propagate:
            logging.error(msg, *args, **kwargs)
        EnodebdLogger._LOGGER.error(msg, *args, **kwargs)

    @staticmethod
    def exception(msg, *args, **kwargs):
        if not EnodebdLogger._LOGGER.propagate:
            logging.exception(msg, *args, **kwargs)
        EnodebdLogger._LOGGER.exception(msg, *args, **kwargs)

    @staticmethod
    def critical(msg, *args, **kwargs):
        if not EnodebdLogger._LOGGER.propagate:
            logging.critical(msg, *args, **kwargs)
        EnodebdLogger._LOGGER.critical(msg, *args, **kwargs)
