#!/usr/bin/env python
# -*- coding:utf-8

def quick_sort(arr):
    arr_length = len(arr)
    if ( arr_length <= 1):
        return arr
    else:
        PIVOT = arr[0]
        GREATER = [ element for element in arr[1:] if element > PIVOT ]
        LESSER = [ element for element in arr[1:] if element <= PIVOT ]
        return quick_sort(LESSER) + [PIVOT] + quick_sort(GREATER)

if __name__ == '__main__':
    unsorted = [1,8,5,3,126,33,25,26,35,14,52]
    print(quick_sort(unsorted))
