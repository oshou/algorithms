def gcd(a, b: int) -> int:
    if b == 0:
        return a
    return gcd(b, a % b)


if __name__ == '__main__':
    print(gcd(15, 20))
