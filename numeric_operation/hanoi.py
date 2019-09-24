#!/usr/bin/env python


def hanoi(n, from_stack, to_stack, tmp_stack):
    if n <= 1:
        print(f"move {n} to disk from {from_stack} to {to_stack}")
        return
    hanoi(n-1, from_stack, tmp_stack, to_stack)
    print(f"move {n} to disk from {from_stack} to {to_stack}")
    hanoi(n-1, tmp_stack, to_stack, from_stack)


if __name__ == '__main__':
    n = int(input("n: "))
    from_stack = input("from: ")
    to_stack = input("to: ")
    tmp_stack = input("tmp: ")
    hanoi(n, from_stack, to_stack, tmp_stack)
