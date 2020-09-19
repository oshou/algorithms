from main import Solution


class TestSolution:
    def test_gcd(self):
        solution = Solution()
        cases = [
            [15, 5, 5],
            [8, 6, 2],
            [18, 12, 6],
            [1, 1, 1],
        ]
        for a, b, expected in cases:
            assert solution.gcd(a, b) == expected
