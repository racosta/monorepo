import unittest

from calculator import Calculator


class TestSum(unittest.TestCase):
    def test_sum(self):
        c = Calculator()
        self.assertEqual(c.add(1, 2), 3)

    def test_sum_negative(self):
        c = Calculator()
        self.assertEqual(c.add(-4, -11), -15)


if __name__ == "__main__":
    unittest.main()
