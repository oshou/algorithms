#!/usr/bin/env python
# -*- coding:utf-8


def fizzbuzz(n):
    for i in range(1, n):
        if i % 15 == 0:
            print("i:", i, "FizzBuzz!")
        elif i % 3 == 0:
            print("i:", i, "Fizz!")
        elif i % 5 == 0:
            print("i:", i, "Buzz!")
        else:
            print("i:", i)


if __name__ == '__main__':
    n = int(input("n: "))
    fizzbuzz(n)
