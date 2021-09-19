<?php

require __DIR__ . "./../../vendor/autoload.php";

use Algorithms\Sort;

class SelectionSort
{
    static function run(array $arr): array
    {
        $size = count($arr) - 1;
        $minIndex = -1;
        for ($i = 0; $i < $size; $i++) {
            for ($j = $i; $j <= $size; $j++) {
                if ($minIndex < 0) {
                    $minIndex = $j;
                }
                if ($arr[$minIndex] > $arr[$j]) {
                    $minIndex = $j;
                }
            }
            $tmp = $arr[$i];
            $arr[$i] = $arr[$minIndex];
            $arr[$minIndex] = $tmp;
            $minIndex = -1;
            print_r($arr);
        }
        return $arr;
    }
}

$unsorted = [1, 8, 5, 3, 126, 33, 25, 26, 35, 14, 52];
$sorted = SelectionSort::run($unsorted);
print_r($sorted);
