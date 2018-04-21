<?php

list($a,$b) = explode(" ",trim(fgets(STDIN)));
if ($a > $b)
    echo "a > b".PHP_EOL;
elseif ($a < $b)
    echo "a < b".PHP_EOL;
else
    echo "a == b".PHP_EOL;
