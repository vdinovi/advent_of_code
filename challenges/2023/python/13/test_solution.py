import unittest
import pathlib

from .solution import solve_p1, solve_p2

inputs = pathlib.Path(__file__).resolve().parent / "inputs"


class TestSolution(unittest.TestCase):
    def test_solution_p1(self):
        tests = {
            "sample 1": {
                "file": "sample_1.txt",
                "expected": 405,
            },
            "puzzle 1": {
                "file": "puzzle.txt",
                "expected": 37718,
            },
        }
        for name, test in tests.items():
            with self.subTest(test=name):
                with open(inputs / test["file"]) as f:
                    lines = [l.strip() for l in f.readlines()]
                    answer = solve_p1(lines)
                    self.assertEqual(answer, test["expected"])

    def test_solution_p2(self):
        pass
