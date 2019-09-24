#!/usr/bin/env python
# -*- coding:utf-8


def fibonacci_memo(n):
    memo = [0]*(n+1)

    def _fib(n):
        if n <= 1:
            return 1
        if memo[n]:
            return memo[n]
        memo[n] = _fib(n-2)+_fib(n-1)
        return memo[n]

    return _fib(n)


if __name__ == '__main__':
    n = input("n: ")
    print(fibonacci_memo(int(n)))
