<?php

require __DIR__ . "./../../vendor/autoload.php";

use Algorithms\Sort;

class QuickSort
{
    static function run(array $arr): array
    {
        $lastIndex = count($arr) - 1;
        for ($i = 0; $i < $lastIndex; $i++) {
            for ($j = 0; $j < $lastIndex - $i; $j++) {
                if ($arr[$j] > $arr[$j + 1]) {
                    $tmp = $arr[$j];
                    $arr[$j] = $arr[$j + 1];
                    $arr[$j + 1] = $tmp;
                }
            }
        }
        return $arr;
    }
}

$unsorted = [1, 8, 5, 3, 126, 33, 25, 26, 35, 14, 52];
$sorted = QuickSort::run($unsorted);
print_r($sorted);
