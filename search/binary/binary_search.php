<?php
function binary_search(array $arr, int $v): int
{
    // 中央値のindexを決める
    $mid = intval(count($arr) / 2);
    echo "mid:" . $mid . PHP_EOL;
    if ($arr[$mid] === $v) {
        return $mid;
    }

    if ($arr[$mid] > $v) {
        binary_search(array_slice($arr, 0, $mid), $v);
    } else {
        binary_search(array_slice($arr, $mid), $v);
    }
}

# Random Array
$arr = range(1, 10000, 1);
sort($arr);
$searchValue = 500;

# Sorted Array
$startTime = microtime(true);
$idx = binary_search($arr, $searchValue);
echo $idx . PHP_EOL;
$endTime = microtime(true);
echo "time: " . ($endTime - $startTime) . PHP_EOL;
