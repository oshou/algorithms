#!/usr/bin/env python
# -*- coding:utf-8

def insertion_sort(arr):
    for index in range(1, len(arr)):
        while 0 < index and arr[index] < arr[index-1]:
            arr[index], arr[index-1] = arr[index-1], colleciton[index]
            index -= 1
    return arr

if __name__ == '__main__':
    unsorted = [1,8,5,3,126,33,25,26,35,14,52]
    print(insertion_sort(unsorted))
