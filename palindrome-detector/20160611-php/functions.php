<?php

define('INVALID_CHARS_PATTERN', '/[^a-z0-9]/');

function isPalindrome(string $str): bool
{
    $str = preg_replace(INVALID_CHARS_PATTERN, '', strtolower($str));

    return $str === strrev($str);
}
