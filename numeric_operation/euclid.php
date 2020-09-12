<?php

require __DIR__ . "/../vendor/autoload.php";

use Algorithms\NumericOperation;

class Euclid
{
    static function run(int $a, $b): int
    {
        while ($b > 0) {
            $tmpa = $a;
            $tmpb = $b;
            $a = $tmpb;
            $b = $tmpa % $tmpb;
        }
        return $a;
    }
}

echo Euclid::run(1920, 1080) . PHP_EOL;
echo Euclid::run(1900, 1080) . PHP_EOL;
