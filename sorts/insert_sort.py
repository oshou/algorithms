#!/usr/bin/env python
# -*- coding:utf-8


def insert_sort(arr):
    """
    TimeComplexity: O(n^2)
    """
    n = len(arr)

    # ソート不要であれば即返却
    if n <= 1:
        return arr

    # InsertSort実行
    for i in range(1, n):
        for j in range(i, 0, -1):
            if arr[j] < arr[j-1]:
                arr[j], arr[j-1] = arr[j-1], arr[j]
    return arr


if __name__ == '__main__':
    unsorted = [8, 1, 5, 3, 126, 33, 25, 26, 35, 14, 52]
    print(insert_sort(unsorted))
