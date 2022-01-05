#!/usr/bin/env python3
# Copyright 2022 Canonical Ltd.
# See LICENSE file for licensing details.


import logging
from pathlib import Path

import pytest
import yaml
from helpers import IPAddressWorkaround, is_loki_up  # type: ignore[import]

logger = logging.getLogger(__name__)

METADATA = yaml.safe_load(Path("./metadata.yaml").read_text())
app_name = METADATA["name"]
resources = {"loki-image": METADATA["resources"]["loki-image"]["upstream-source"]}


@pytest.mark.abort_on_fail
async def test_deploy_from_edge_and_upgrade_from_local_path(ops_test, loki_charm):
    """Deploy from charmhub and then upgrade with the charm-under-test."""
    async with IPAddressWorkaround(ops_test):
        logger.debug("deploy charm from charmhub")
        await ops_test.model.deploy(f"ch:{app_name}", application_name=app_name, channel="edge")
        await ops_test.model.wait_for_idle(apps=[app_name], status="active", timeout=1000)

        logger.debug("upgrade deployed charm with local charm %s", loki_charm)
        await ops_test.model.applications[app_name].refresh(path=loki_charm, resources=resources)
        await ops_test.model.wait_for_idle(apps=[app_name], status="active", timeout=1000)
        assert await is_loki_up(ops_test, app_name)
