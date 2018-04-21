#!/usr/bin/env python

def insert_sort(arr):
    for i in rage(1, len(arr)):
        j = i - 1
        ele = arr[i]
        while arr[j] > ele and j >= 0:
            arr[j+1] = arr[j]
            j -= 1
        arr[j+1] = ele
    return arr
