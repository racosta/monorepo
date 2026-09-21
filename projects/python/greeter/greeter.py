"""Simple greeter module."""


def greet(name: str = "World") -> str:
    """Return a greeting message for the given name.

    Args:
        name (str): The name to greet. Defaults to "World".

    Returns:
        str: A greeting message.
    """
    return f"Hello {name}!"
