"""The main Python script that demonstrates how to use the Cython extension.

It imports the math_utils module, which is the Cython extension, and calls the
fast_factorial function to compute the factorial of a number. The result is
printed to the console.
"""

# The import name matches the cython_library target name defined in the BUILD file
from projects.python.cython_example import math_utils


def main():
    """Main function to demonstrate the use of the math_utils Cython extension."""
    num = 20
    result = math_utils.fast_factorial(num)
    print("🐍 Cython execution successful!")
    print(f"The factorial of {num} is: {result}")


if __name__ == "__main__":
    main()
