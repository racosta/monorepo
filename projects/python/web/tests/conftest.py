"""Pytest fixtures for the python_web project."""

from typing import Iterator

import pytest
from flask import Flask
from flask.testing import FlaskClient
from main import create_app


@pytest.fixture()
def app() -> Iterator[Flask]:
    """Create and yield a Flask app.

    Yields:
        Flask: The Flask application.
    """
    app = create_app()

    yield app


@pytest.fixture
def client(app: Flask) -> FlaskClient:
    """Access the Flask app's test client.

    Args:
        app (Flask): The Flask application fixture.

    Returns:
        FlaskClient: The Flask test client.
    """
    return app.test_client()
