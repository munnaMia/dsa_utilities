<?php

namespace Dsa;

class Utils
{
    public static function swap(&$elem1, &$elem2)
    {
        [$elem1, $elem2] = [$elem2, $elem1];
    }
}
