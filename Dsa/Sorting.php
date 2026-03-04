<?php

namespace Dsa;

use Dsa\Utils;

class Sorting
{
    public static function selectionSort($array = [])
    {
        for ($i = 0; $i < count($array); $i++) {
            $min = $i;

            for ($j = $i + 1; $j < count($array); $j++) {
                if ($array[$j] < $array[$min]) {
                    $min = $j;
                }
            }
            Utils::swap($array[$i], $array[$min]);
        }

        return $array;
    }
}
