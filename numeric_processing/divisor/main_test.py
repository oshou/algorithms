from main import Solution


class TestSolution:
    def test_divisor(self):
        solution = Solution()
        cases = [
            [6, [1, 2, 3, 6]],
            [10, [1, 2, 5, 10]],
            [12, [1, 2, 3, 4, 6, 12]]
        ]
        for n, expected in cases:
            assert solution.divisor(n) == expected
