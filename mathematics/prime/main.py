import math


def prime(n: int) -> bool:
    ans = []
    max = int(math.sqrt(n))
    for i in range(2, max+1):
        for j in range(1, i*j)
        if n % i == 0:
            ans = False
    return ans


if __name__ == '__main__':
    print(prime(1000))
