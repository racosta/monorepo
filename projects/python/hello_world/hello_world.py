"""A simple Python program that prints "Hello World!"."""

from projects.python.greeter import greet


def main() -> None:
    """Print a greeting message."""
    print(greet())


if __name__ == "__main__":
    main()
