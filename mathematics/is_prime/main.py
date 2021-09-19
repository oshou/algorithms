import math


class Solution:
    def is_prime(self, n: int) -> bool:
        max = int(math.sqrt(n))
        ans = True
        for i in range(2, max+1):
            if n % i == 0:
                ans = False
        return ans
