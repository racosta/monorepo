"""Test the root endpoint of the FastAPI application."""

from __future__ import annotations

import typing

if typing.TYPE_CHECKING:
    from fastapi.testclient import TestClient


def test_read_root(client: TestClient):
    """Test the root endpoint of the FastAPI application.

    Args:
        client (TestClient): The FastAPI test client fixture.
    """
    response = client.get("/")
    assert response.status_code == 200
    assert response.json() == {"message": "Hello World"}
