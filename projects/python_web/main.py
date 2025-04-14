from random import randint

from flask import Flask

from projects.python_calculator.calculator import Calculator


def create_app():
    app = Flask(__name__)
    my_calculator = Calculator()

    @app.route("/")
    def add() -> str:
        """Generate two random integers and add them.

        Returns:
            str - Message containing the addition expression.
        """
        num1 = randint(0, 100)
        num2 = randint(0, 100)
        message = f"{num1} + {num2} = {my_calculator.add(num1, num2)}"
        return message

    return app


if __name__ == "__main__":
    app = create_app()
    app.run()
