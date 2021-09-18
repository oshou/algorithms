<?php

require __DIR__ . "/../vendor/autoload.php";

use Algorithms\Sorts;

class MergeSort
{
    static function run(array $arr, int $first, $last)
    {
        if ($last <= $first) {
            echo "last: " . $last . PHP_EOL;
            echo "first: " . $first . PHP_EOL;
            return;
        }
        $mid = intval(($first + $last) / 2);
        echo $mid . PHP_EOL;
        self::run($arr, $first, $mid);
        self::run($arr, $mid + 1, $last);
        print_r($arr);
    }
}

# Random Array
$arr = range(1, 10, 1);
shuffle($arr);
print_r($arr);

# Sorted Array
$startTime = microtime(true);
MergeSort::run($arr, 0, count($arr));
$endTime = microtime(true);
