"""A simple FastAPI application that responds with a greeting message."""

import uvicorn
from fastapi import FastAPI


def main():
    """Main entry point for the FastAPI application."""
    app = create_app()

    uvicorn.run(app, host="0.0.0.0", port=8080)  # noqa: S104 #nosec B104


def create_app() -> FastAPI:
    """Create and return a FastAPI application instance.

    Returns:
        FastAPI: The FastAPI application instance.
    """
    app = FastAPI()

    # Map GET requests to the root URL "/" to this function
    app.add_api_route("/", endpoint=read_root, methods=["GET"])

    return app


async def read_root() -> dict[str, str]:
    """Handle GET requests to the root URL and return a greeting message.

    Returns:
        dict[str, str]: A dictionary containing a greeting message.
    """
    return {"message": "Hello World"}


if __name__ == "__main__":
    main()
