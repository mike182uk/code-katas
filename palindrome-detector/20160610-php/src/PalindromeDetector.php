<?php

namespace Mdb;

class PalindromeDetector
{
    const INVALID_CHARS_PATTERN = '/[^a-z0-9]/';

    public static function test(string $str): bool
    {
        $str = preg_replace(self::INVALID_CHARS_PATTERN, '', strtolower($str));

        return $str === strrev($str);
    }
}
