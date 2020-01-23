<?php
class Solution
{
    public $memo = [];

    function getFib(int $n): int
    {
        if ($n == 0 || $n == 1) {
            $memo[$n] = 1;
            return $memo[$n];
        }
        if (isset($memo[$n])) {
            return $memo[$n];
        }
        $memo[$n] = $this->getFib($n - 1) + $this->getFib($n - 2);
        return $memo[$n];
    }
}

$sol = new Solution;
fscanf(STDIN, "%d", $n);
$time_start = microtime(true);
echo $sol->getFib($n) . PHP_EOL;
echo microtime(true) - $time_start;
