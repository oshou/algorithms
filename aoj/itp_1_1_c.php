<?php

$x = trim(fgets(STDIN));
List($a,$b) = explode(" ",$x);
$s = $a * $b;
$l = ($a + $b) * 2;
echo "{$s} {$l}".PHP_EOL;
