from functools import lru_cache

@lru_cache(maxsize=1024)
def getFib(n):
    global cnt
    cnt += 1
    if (n == 0): return 1;
    if (n == 1): return 2;
    return getFib(n-1) + getFib(n-2)

n = int(input())
cnt = 0
print("fib(n): ", getFib(n))
print("cnt: ", cnt)
