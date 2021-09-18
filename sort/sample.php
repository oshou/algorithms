<?php

$arr = [1, 2, 3, 4, 5];
$result = array_filter($arr, function ($num) {
    return ($num % 2 === 0);
});
print_r($result);
