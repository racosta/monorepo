"""is_leap_year returns True if the given year is a leap year, and False otherwise."""


def is_leap_year(year: int) -> bool:
    """Returns whether the given year is a leap year.

    Args:
      year (int): The year to check.

    Returns:
      bool: True if the year is a leap year, False otherwise.
    """
    return (year % 4 == 0 and year % 100 != 0) or year % 400 == 0
