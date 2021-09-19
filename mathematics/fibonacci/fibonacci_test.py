from fibonacci import Solution


class TestSolution:
    def test_fib(self):
        solution = Solution()
        cases = [
            [1, 0],
            [2, 1],
            [3, 2],
        ]
        for arg, expected in cases:
            assert solution.fib(arg) == expected
