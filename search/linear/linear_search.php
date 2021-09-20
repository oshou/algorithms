<?php

function linear_search(array $arr, int $v): bool
{
    for ($i = 0; $i < count($arr); $i++) {
        if ($arr[$i] === $v) {
            return true;
        }
    }
    return false;
}

$arr = [1, 3, 9, 2, 7, 3, 4, 7];
$value = 3;
echo linear_search($arr, $value) . PHP_EOL;
