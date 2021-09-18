#!/usr/bin/env python
# -*- coding:utf-8


def bubble_sort(arr):
    """
    TimeComplexity: O(n^2)
    """
    n = len(arr)

    # ソート不要であれば即返却
    if n <= 1:
        return arr

    # BubbleSort実行
    for i in range(n-1):
        for j in range(n-1, i, -1):
            if arr[j] < arr[j-1]:
                arr[j], arr[j-1] = arr[j-1], arr[j]
    return arr


if __name__ == '__main__':
    unsorted = [1, 8, 5, 3, 126, 33, 25, 26, 35, 14, 52]
    print(bubble_sort(unsorted))
