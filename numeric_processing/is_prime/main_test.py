from main import Solution


class TestSolution:
    def test_is_prime(self):
        solution = Solution()
        cases = [
            [17, True],
            [211, True],
            [2200, False],
        ]
        for n, expected in cases:
            assert solution.is_prime(n) == expected
