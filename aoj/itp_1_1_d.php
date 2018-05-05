<?php

$S = fgets(STDIN);
$h = floor($S / 3600);
$m = floor(($S % 3600) / 60);
$s = ($S % 3600) % 60;

echo $h.':'.$m.':'.$s.PHP_EOL;
