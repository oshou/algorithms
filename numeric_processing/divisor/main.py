from typing import List


class Solution:
    def divisor(self, n: int) -> List[int]:
        ans = []
        for i in range(1, n+1):
            if n % i == 0:
                ans.append(i)
        return ans
