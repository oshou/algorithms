class Solution:
    def gcd(self, a, b: int) -> int:
        while b:
            a, b = b, a % b
        return a
