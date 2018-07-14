#!/usr/bin/env python
# -*- coding:utf-8

def bubble_sort(arr):
    length = len(arr)
    for i in range(length):
        for j in range(length-1):
            if arr[j] > arr[j+1]:
                arr[j],arr[j+1] = arr[j+1],arr[j]
    return arr

if __name__ == '__main__':
    unsorted = [1,8,5,3,126,33,25,26,35,14,52]
    print(bubble_sort(unsorted))
