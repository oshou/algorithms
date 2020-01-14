<?php
function bubble_sort($array){
    for ($i=0;$i<count($array);$i++){
        for($n=1;$n<count($array);$n++){
            if($array[$n-1] > $array[$n]){
                $temp = $array[$n];
                $array[$n] = $array[n-1];
                $array[$n-1] = $temp;
            }
        }
    }
    return $array;
}
