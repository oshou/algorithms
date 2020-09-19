from main import Solution


class TestSolution:
    def test_lcm(self):
        solution = Solution()
        cases = [
            [2, 3, 6],
            [1, 1, 1],
            [18, 1, 18],
            [9, 6, 18],
        ]
        for a, b, expected in cases:
            assert solution.lcm(a, b) == expected
