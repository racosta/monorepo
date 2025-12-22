import pytest
from is_leap_year import is_leap_year


@pytest.mark.parametrize(
    "year, expected",
    [
        (2016, True),
        (2020, True),
        (2024, True),
        (2100, False),
        (2200, False),
        (2300, False),
        (2000, True),
        (2400, True),
        (2019, False),
    ],
    ids=[
        "2016 :: divisible by 4",
        "2020 :: divisible by 4",
        "2024 :: divisible by 4",
        "2100 :: divisible by 4 AND 100",
        "2200 :: divisible by 4 AND 100",
        "2300 :: divisible by 4 AND 100",
        "2000 :: divisible by 400",
        "2400 :: divisible by 400",
        "2019 :: not divisible by 4",
    ],
)
def test_is_leap_year(year, expected):
    actual = is_leap_year(year)
    assert actual == expected, f"Expected {expected} for year {year}, but got {actual}."


if __name__ == "__main__":
    raise SystemExit(pytest.main([__file__]))
