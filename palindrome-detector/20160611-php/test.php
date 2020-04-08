<?php

class PalindromeDetectorTest extends PHPUnit_Framework_TestCase
{
    public function testItCanDetectAPalindrome()
    {
        $this->assertEquals(isPalindrome('racecar'), true);
        $this->assertEquals(isPalindrome('12321'), true);
        $this->assertEquals(isPalindrome('1AB232BA1'), true);
        $this->assertEquals(isPalindrome('foobar'), false);
    }

    public function testItNormalisesStringBeforeDetectingPalindrome()
    {
        $this->assertEquals(isPalindrome('A man, a plan, a canal, Panama!'), true);
        $this->assertEquals(isPalindrome('Was it a car or a cat I saw?'), true);
    }
}
