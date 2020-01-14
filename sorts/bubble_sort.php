<?php

function bubble_sort(array $arr): array
{
    for ($i = count($arr) - 1; $i >= 0; $i--) {
        for ($j = count($arr) - 1; $j >= count($arr) - $i; $j--) {
            if ($arr[$j - 1] > $arr[$j]) {
                $tmp = $arr[$j];
                $arr[$j] = $arr[$j - 1];
                $arr[$j - 1] = $tmp;
            }
        }
    }
    return $arr;
}

# Random Array
$arr = range(1, 1000, 1);
shuffle($arr);

# Sorted Array
$startTime = microtime(true);
$sorted = bubble_sort($arr);
#print_r($sorted);
$endTime = microtime(true);
echo "time: " . ($endTime - $startTime) . PHP_EOL;
