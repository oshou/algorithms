<?php

function select_sort(array $arr): array
{
    $sorted = [];
    while (is_null($arr)) {
        print_r($arr);
        for ($i = 0; $i < count($arr); $i++) {
            array_merge($sorted, array_splice($arr, array_search(min($arr), $arr), 1));
        }
    }
    return $arr;
}

$arr = [1, 3, 8, 6, 2, 9, 11, 5, 19];
$sorted = select_sort($arr);
print_r($sorted);
