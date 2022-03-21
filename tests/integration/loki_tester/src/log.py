#!/usr/bin/env python3
# Copyright 2022 Canonical Ltd.
# See LICENSE file for licensing details.

import json
import logging
import os
import sys
import syslog

from time import sleep, time
from typing import Optional
from urllib import request

# to run manually inside container:
# j ssh -container loki-tester loki-tester/0 bash
# python3 ./log.py syslog,file,loki \
#    http://loki-k8s-0.loki-k8s-endpoints.test-0.svc.cluster.local:3100/loki/api/v1/push \  # noqa
#    /loki_tester_msgs.txt

logger = logging.getLogger(__name__)
try:
    # to debug log.py when running inside a container; we dump the logs
    # to a logpy.log file; on localhost, this logger will print to stdout
    # and that is enough.
    logger.addHandler(logging.FileHandler("/logpy.log"))
    CONTAINER = True
except PermissionError:
    # not in a container; skip this
    CONTAINER = False
logger.setLevel("INFO")

LONG = False
DEBUG = False
TEST_JOB_NAME = "test-job"
SYSLOG_LOG_MSG = "LOG SYSLOG"
LOKI_LOG_MSG = "LOG LOKI"
FILE_LOG_MSG = "LOG FILE"
WAIT = 0.2


def _log_to_syslog():
    if CONTAINER:
        import socket
        try:
            from rfc5424logging import Rfc5424SysLogHandler
        except ImportError as e:
            # we'd normally do this by baking the requirement in the
            # workload image but here we don't really care. we just want the
            # import to succeed
            os.system('pip install rfc5424-logging-handler')
            try:
                return _log_to_syslog()
            except ImportError:
                # something went wrong;
                # package not installed?
                raise e

        syslogger = logging.getLogger("promtail")
        handler = Rfc5424SysLogHandler(address=('localhost', 1514),
                                       socktype=socket.SOCK_STREAM)
        handler.setLevel(logging.DEBUG)
        syslogger.addHandler(handler)
        syslogger.warning(SYSLOG_LOG_MSG, extra={"msgid": 1})

        logger.info('logged to syslog via rfc5424handler.')

    else:
        syslog.syslog(SYSLOG_LOG_MSG)
        logger.info(f"logged to syslog")


def _log_to_loki(loki_address):
    # `loki_address` is the address of the push endpoint e.g.
    # 'http://loki-k8s-0.loki-k8s-endpoints.cos.svc.cluster.local:3100/loki/api/v1/push'
    loki_base_url = loki_address[:-17]

    try:
        with request.urlopen(f"{loki_base_url}/ready") as resp:
            resp_text = resp.read()
    except Exception as e:
        try:
            if resp.getcode() == 503:
                raise RuntimeError(
                    f"loki gives Service Unavailable at {loki_base_url}")
        except:  # noqa
            pass
        raise RuntimeError(
            f"Could not contact loki api at {loki_base_url!r}; gotten {e!r};"
            f"loki_address={loki_address!r}"
        )

    if not resp_text.decode("ascii").strip() == "ready":
        raise RuntimeError("loki not ready yet... give it time.")

    data = {
        "streams": [
            {
                "stream": {"job": TEST_JOB_NAME},
                "values": [[str(int(time())) + "0" * 9, LOKI_LOG_MSG]],
            }
        ]
    }

    enc = json.dumps(data).encode("ascii")
    req = request.Request(loki_address)
    req.add_header("Content-Type", "application/json; charset=utf-8")
    req.add_header("Content-Length", str(len(enc)))

    try:
        resp = request.urlopen(req, enc)
    except Exception as e:
        raise RuntimeError(
            f"Could not contact loki api at {loki_address}; gotten {e}")

    if resp.getcode() != 204:
        raise RuntimeError(
            f"contacted loki, but request didn't go well: "
            f"gotten {resp.status_code}: {resp.reason}"
        )

    logger.info(f"logged to loki at {loki_address}")


def _log_to_file(fname: str):
    with open(fname, "a+") as f:
        f.write(FILE_LOG_MSG + "\n")
    logger.info(f"logged to file at {fname}")


def _log(mode, loki_address: str, fname: str):
    if mode == "NOOP":
        return
    if mode == "syslog":
        _log_to_syslog()
    elif mode == "loki":
        if not loki_address:
            raise RuntimeError(
                "loki_address needs to be provided when in `loki` mode")
        _log_to_loki(loki_address)
    elif mode == "file":
        if not fname:
            raise RuntimeError("fname needs to be provided when in `file` mode")
        _log_to_file(fname)
    else:
        raise ValueError(f"unknown logging mode: {fname!r}")


def main(modes: str, loki_address: Optional[str] = None,
         fname: Optional[str] = None):
    if modes == "ALL":
        log_modes = ["syslog", "loki", "file"]
    elif modes:
        log_modes = modes.split(",")
    else:
        raise ValueError(f"modes cannot be {modes!r}")

    # we do it a few more times so we have the time to inspect the running
    # process when debugging
    rng = 200 if LONG else 10  # TODO try lowering to 1
    for _ in range(rng):
        if not log_modes:
            break

        for mode in log_modes:
            try:
                _log(mode, loki_address, fname)
            except Exception as e:
                if DEBUG:
                    raise e
                logger.error(
                    f"failed to _log({mode!r}, {loki_address!r}, {fname!r}); "
                    f"got {e}")
        if LONG:
            sleep(WAIT)

    logger.info(f"executed: logpy {log_modes}({modes}) {loki_address} {fname}")
    print("DONE")


if __name__ == "__main__":
    main(*sys.argv[1:4])