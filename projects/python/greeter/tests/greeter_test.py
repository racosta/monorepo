"""Unit tests for the greeter module."""

import pytest

from projects.python.greeter import greet


@pytest.mark.parametrize(
    "name, want",
    [
        ("Alice", "Hello Alice!"),
        ("Bob", "Hello Bob!"),
        (None, "Hello World!"),
    ],
    ids=[
        "greet Alice",
        "greet Bob",
        "greet None (default)",
    ],
)
def test_greet(name: str, want: str):
    """Test the greet function.

    Args:
        name (str): The name to greet.
        want (str): The expected greeting message.
    """
    got = greet(name) if name is not None else greet()

    assert got == want, f"Want {want}, but got {got}"
