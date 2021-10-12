# Fermat's Last Theorem
#
# 0 < x < y < z
# x**n + y**n  = z**n (n = 1,2)
# x**n + y**n != z**n (n >= 3)
#
# Input: 10, 1 => []
# Input: 10, 2 => [(3, 4, 5), (6, 8, 10)]
# Input: 10, 3 => []
# Input: 10, 4 => []

from typing import List, Tuple
import sys
import time


def fermat_last_theorem_v1(max_num: int, square_num: int) -> List[Tuple[int, int, int]]:
    result = []
    max_z = int(pow((max_num-1)**square_num +
                max_num**square_num, 1.0 / square_num))
    for x in range(1, max_num+1):
        for y in range(x+1, max_num+1):
            for z in range(y+1, max_z):
                if pow(x, square_num) + pow(y, square_num) == pow(z, square_num):
                    result.append((x, y, z))
    return result


def fermat_last_theorem_v2(max_num: int, square_num: int) -> List[Tuple[int, int, int]]:
    result = []
    for x in range(1, max_num+1):
        for y in range(x+1, max_num+1):
            pow_sum = pow(x, square_num) + pow(y, square_num)
            # if pow_sum > sys.maxsize:
            #    raise ValueError(x, y, z, square_num, pow_sum)

            z = pow(pow_sum, 1.0/square_num)
            if not z.is_integer():
                continue

            # for rounding error
            # ex) z = 5.00000000000000000001
            z = int(z)
            z_pow = pow(z, square_num)
            if z_pow == pow_sum:
                result.append((x, y, z))
    return result


if __name__ == '__main__':
    start = time.time()
    print('v1_**1', fermat_last_theorem_v1(10, 1))
    print('v1_**2', fermat_last_theorem_v1(10, 2))
    print('v1_**3', fermat_last_theorem_v1(10, 3))
    print('v1_**4', fermat_last_theorem_v1(10, 4))
    print('v1_**5', fermat_last_theorem_v1(10, 5))
    print('v1_**18', fermat_last_theorem_v1(200, 5))
    end = time.time()
    print(end-start)

    start = time.time()
    print('v2_**1', fermat_last_theorem_v2(10, 1))
    print('v2_**2', fermat_last_theorem_v2(10, 2))
    print('v2_**3', fermat_last_theorem_v2(10, 3))
    print('v2_**4', fermat_last_theorem_v2(10, 4))
    print('v2_**5', fermat_last_theorem_v2(10, 5))
    print('v2_**18', fermat_last_theorem_v2(200, 5))
    end = time.time()
    print(end-start)
