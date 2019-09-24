#!/usr/bin/env python
# -*- cording:utf-8


def select_sort(arr):
    n = len(arr)

    # ソート不要であれば即返却
    if n <= 1:
        return arr

    # SelectSort実行
    for i in range(n-1):
        # 最小値のインデックスを保持
        min = i
        # iから末尾までで最小値を確認して最小値インデックスの値と差し替え
        for j in range(i+1, n):
            if arr[min] > arr[j]:
                arr[min], arr[j] = arr[j], arr[min]
    return arr


if __name__ == '__main__':
    unsorted = [1, 8, 5, 3, 126, 33, 25, 26, 35, 14, 52]
    print(select_sort(unsorted))
