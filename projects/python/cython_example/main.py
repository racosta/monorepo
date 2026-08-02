"""The main Python script that demonstrates how to use the Cython extension.

It imports the math_utils module, which is the Cython extension, and calls the
fast_factorial function to compute the factorial of a number. The result is
printed to the console.
"""

import time

# The import name matches the cython_library target name defined in the BUILD file
from projects.python.cython_example import math_utils


def slow_factorial(n: int) -> int:
    """A slow factorial implementation using pure Python for comparison.

    Args:
        n (int): The number to compute the factorial of.

    Returns:
        int: The factorial of n.

    Raises:
        ValueError: If n is negative.
    """
    if n < 0:
        raise ValueError("Factorial not defined for negative numbers.")
    elif n == 0 or n == 1:
        return 1
    else:
        return n * slow_factorial(n - 1)


def main():
    """Main function to demonstrate the use of the math_utils Cython extension."""
    num = 20

    t1 = time.perf_counter()
    result = slow_factorial(num)
    t2 = time.perf_counter()
    print("🐍 Pure Python execution successful!")
    print(f"The factorial of {num} is: {result}")
    print(f"⏱️ Pure Python execution time: {t2 - t1:.6f} seconds")

    t1 = time.perf_counter()
    result = math_utils.fast_factorial(num)
    t2 = time.perf_counter()
    print("🐍 Cython execution successful!")
    print(f"The factorial of {num} is: {result}")
    print(f"⏱️ Cython execution time: {t2 - t1:.6f} seconds")


if __name__ == "__main__":
    main()
