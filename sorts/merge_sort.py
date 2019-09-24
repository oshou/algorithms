#!/usr/bin/env python
# -*- coding: utf-8


def merge_sort(arr):
    n = len(arr)

    # ソート不要であれば即返却
    if n <= 1:
        return arr

    # 分割
    q, r = divmod(len(arr), 2)
    left = [arr.pop() for _ in range(q)]
    right = [arr.pop() for _ in range(q + r)]

    # ソート
    merge_sort(left)
    merge_sort(right)

    # 統合
    left.reverse()
    right.reverse()
    while left and right:
        if left[-1] <= right[-1]:
            arr.append(left.pop())
        else:
            arr.append(right.pop())
    while left:
        arr.append(left.pop())
    while right:
        arr.append(right.pop())

    return arr


if __name__ == '__main__':
    unsorted = [1, 8, 5, 3, 126, 33, 25, 26, 35, 14, 52]
    print(merge_sort(unsorted))
