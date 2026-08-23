"""Pytest fixtures for the fastapi project."""

from typing import Iterator

import pytest
from fastapi.testclient import TestClient
from main import create_app


@pytest.fixture()
def client() -> Iterator[TestClient]:
    """Create and yield a FastAPI test client.

    Yields:
        TestClient: The FastAPI test client.
    """
    app = create_app()
    with TestClient(app) as client:
        yield client
