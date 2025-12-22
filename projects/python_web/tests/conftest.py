"""Pytest fixtures for the python_web project."""

import pytest
from flask import Flask
from flask.testing import FlaskClient
from main import create_app


@pytest.fixture()
def app() -> Flask:
    """Create and yield a Flask app.

    Returns:
        Flask : The Flask application.
    """
    app = create_app()

    yield app


@pytest.fixture
def client(app) -> FlaskClient:
    """Access the Flask app's test client.

    Returns:
        FlaskClient : The Flask test client.
    """
    return app.test_client()
