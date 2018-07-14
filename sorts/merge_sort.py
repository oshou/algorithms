#!/usr/bin/env python
# -*- coding:utf-8

def merge_sort(arr):
    if len(arr) > 1:
        mid = len(arr) / 2
        left = merge_sort(arr[:mid])
        right = merge_sort(arr[mid:])
        arr = []

        while len(left) != 0 and len(right) != 0:
            if left[0] < right[0]:
                arr.apend(left.pop(0))
            else:
                arr.apend(right.pop(0))

            if len(left) != 0:
                arr.extend(left)

            if len(right) != 0:
                arr.extend(right)

        return arr

if __name__ == '__main__':
    unsorted = [1,8,5,3,126,33,25,26,35,14,52]
    print(merge_sort(unsorted))
