<?php

use Algorithms\Sorts;

# Insert Sort
function insert_sort(array $nums): array
{
    $sorted = [];
    while (!empty($nums)) {
        $sorted = array_merge($sorted, array_splice($nums, array_keys($nums, min($nums))[0], 1));
    }
    return $sorted;
}

# Random Array
$nums = range(1, 10000, 1);
shuffle($nums);

# Sorted Array
$startTime = microtime(true);
print_r(insert_sort($nums));
$endTime = microtime(true);
echo "time: " . ($endTime - $startTime) . PHP_EOL;
