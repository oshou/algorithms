<?php

require __DIR__ . "/../vendor/autoload.php";

use Algorithms\NumericOperation;

class Factorial
{
    static function run(int $n)
    {
        if ($n === 1) {
            return 1;
        }
        return $n * self::run($n - 1);
    }
}

echo Factorial::run(1) . PHP_EOL;
echo Factorial::run(3) . PHP_EOL;
echo Factorial::run(4) . PHP_EOL;
echo Factorial::run(5) . PHP_EOL;
