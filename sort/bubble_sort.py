#!/usr/bin/env python
def bubble_sort(arr):
    is_change = True
    while is_change:
        is_change = False
        for i in range(len(arr) - 1):
            if arr[i] > arr[i+1]:
                arr[i], arr[i+1] = arr[i+1], arr[i]
                is_change = True
    return arr
