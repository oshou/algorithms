<?php
function linear_search(array $arr, int $n): int
{
    for ($i = 0; $i < count($arr); $i++) {
        if ($arr[$i] === $n) {
            return $i;
        }
    }
    return -1;
}

# Random Array
$arr = range(1, 10000, 1);
sort($arr);
$searchValue = 500;

# Sorted Array
$startTime = microtime(true);
$index = linear_search($arr, $searchValue);
echo $index . PHP_EOL;
$endTime = microtime(true);
echo "time: " . ($endTime - $startTime) . PHP_EOL;
