def euclid1(a, b):
    while b > 0:
        a, b = b, a % b
    return a


def euclid2(a, b):
    if b == 0:
        return a
    return euclid2(b, a % b)


if __name__ == "__main__":
    print("euclid1: "+str(euclid1(1920, 1000)))
    print("euclid2: "+str(euclid2(1920, 1000)))
