"""Unit tests for the add function."""

from __future__ import annotations

import re
import typing

if typing.TYPE_CHECKING:
    from flask.testing import FlaskClient


def test_add(client: FlaskClient):
    """Test the add function.

    Args:
        client (FlaskClient): The Flask test client.
    """
    response = client.get("/")
    assert response.status_code == 200
    assert bool(re.match(r"[1-9]\d* \+ [1-9]\d* = [1-9]\d*", response.data.decode()))
