class Solution:
    def gcd(self, a, b: int) -> int:
        while b:
            a, b = b, a % b
        return a

    def lcm(self, a, b: int) -> int:
        return a*b // self.gcd(a, b)
