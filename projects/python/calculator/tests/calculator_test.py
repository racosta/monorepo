"""Unit tests for the Calculator class."""

import pytest

from projects.python.calculator import Calculator


@pytest.mark.parametrize(
    "a, b, want",
    [
        (1, 2, 3),
        (-4, -11, -15),
        (3, -7, -4),
    ],
    ids=[
        "positive numbers :: 1 + 2 == 3",
        "negative numbers :: -4 + -11 == -15",
        "mixed numbers :: 3 + -7 == -4",
    ],
)
def test_sum(a, b, want):
    """Test computing a sum."""
    c = Calculator()
    got = c.add(a, b)

    assert got == want, f"Want {want}, but got {got}"
