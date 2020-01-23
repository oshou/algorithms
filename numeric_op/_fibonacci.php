<?php
class Solution
{
    function getFib(int $n): int
    {
        if ($n == 0 || $n == 1) {
            return 1;
        }
        return $this->getFib($n - 1) + $this->getFib($n - 2);
    }
}

$sol = new Solution;
fscanf(STDIN, "%d", $n);
$time_start = microtime(true);
echo $sol->getFib($n) . PHP_EOL;
echo microtime(true) - $time_start;
