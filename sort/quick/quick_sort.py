#!/usr/bin/env python


def quick_sort(arr):
    n = len(arr)

    # ソート不要であれば即返却
    if n <= 1:
        return arr

    # 分割
    low = []
    pivot = arr.pop()
    high = []
    while arr:
        e = arr.pop()
        if e <= pivot:
            low.append(e)
        else:
            high.append(e)

    # ソート
    quick_sort(low)
    quick_sort(high)

    # 結合
    low.reverse()
    while low:
        arr.append(low.pop())
    arr.append(pivot)
    high.reverse()
    while high:
        arr.append(high.pop())

    return arr


if __name__ == '__main__':
    unsorted = [1, 8, 5, 3, 126, 33, 25, 26, 35, 14, 52]
    print(quick_sort(unsorted))
