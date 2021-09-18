<?php

require __DIR__ . "/../vendor/autoload.php";

use Algorithms\Sorts;

# Insert Sort
class InsertSort
{
    static function run(array $nums): array
    {
        $sorted = [];
        while (!empty($nums)) {
            $sorted = array_merge($sorted, array_splice($nums, array_keys($nums, min($nums))[0], 1));
        }
        return $sorted;
    }
}

# Random Array
$arr = range(1, 1000, 1);
shuffle($arr);
#print_r($arr);

# Sorted Array
$startTime = microtime(true);
$sorted = InsertSort::run($arr);
#print_r($sorted);
$endTime = microtime(true);
echo "time: " . ($endTime - $startTime) . PHP_EOL;
