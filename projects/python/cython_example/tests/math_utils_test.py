"""Unit tests for the math_utils Cython extension module."""

import pytest

from projects.python.cython_example import math_utils  # ty: ignore[unresolved-import]


@pytest.mark.parametrize(
    "num, expected",
    [
        (0, 1),
        (1, 1),
        (5, 120),
        (10, 3628800),
        (20, 2432902008176640000),
    ],
    ids=[
        "0! = 1",
        "1! = 1",
        "5! = 120",
        "10! = 3628800",
        "20! = 2432902008176640000",
    ],
)
def test_fast_factorial(num: int, expected: int):
    """Test the fast_factorial function from the math_utils Cython extension.

    Args:
        num (int): The number to compute the factorial of.
        expected (int): The expected result of the factorial computation.
    """
    actual = math_utils.fast_factorial(num)
    assert actual == expected, f"Expected {expected} for {num}!, but got {actual}."
