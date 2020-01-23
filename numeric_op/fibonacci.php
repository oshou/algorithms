<?php

require __DIR__ . "/../vendor/autoload.php";

use Algorithms\NumericOperation;

class Fibonacci
{
    static function run(int $n): int
    {
        if ($n === 1 || $n === 2) {
            return 1;
        }
        return self::run($n - 1) + self::run($n - 2);
    }
}

$startTime = microtime(true);
echo Fibonacci::run(3) . PHP_EOL;
echo Fibonacci::run(4) . PHP_EOL;
echo Fibonacci::run(5) . PHP_EOL;
echo Fibonacci::run(6) . PHP_EOL;
echo Fibonacci::run(7) . PHP_EOL;
echo Fibonacci::run(9) . PHP_EOL;
echo Fibonacci::run(10) . PHP_EOL;
echo Fibonacci::run(11) . PHP_EOL;
echo Fibonacci::run(12) . PHP_EOL;
echo Fibonacci::run(13) . PHP_EOL;
$endTime = microtime(true);
echo "time: " . ($endTime - $startTime) . PHP_EOL;
