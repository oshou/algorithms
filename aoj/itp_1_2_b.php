<?php

list($a,$b,$c) = explode(" ",trim(fgets(STDIN)));
if ($a < $b && $b < $c){
    echo "Yes".PHP_EOL;
} else {
    echo "No".PHP_EOL;
}
