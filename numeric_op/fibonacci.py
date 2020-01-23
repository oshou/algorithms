#!/usr/bin/env python
# -*- coding:utf-8


def fibonacci(n):
    if n <= 1:
        return 1
    return fibonacci(n-1)+fibonacci(n-2)


if __name__ == '__main__':
    n = input("n: ")
    print(fibonacci(int(n)))
