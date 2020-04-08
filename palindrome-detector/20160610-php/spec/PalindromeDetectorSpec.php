<?php

namespace spec\Mdb;

use PhpSpec\ObjectBehavior;
use Prophecy\Argument;

class PalindromeDetectorSpec extends ObjectBehavior
{
    function it_should_detect_a_palindrome()
    {
        $this::test('racecar')->shouldReturn(true);
        $this::test('12321')->shouldReturn(true);
        $this::test('1AB232BA1')->shouldReturn(true);
        $this::test('foobar')->shouldReturn(false);
    }

    function it_should_remove_capital_letters_punctuation_and_word_dividers_before_detecting_a_palindrome()
    {
      $this::test('A man, a plan, a canal, Panama!')->shouldReturn(true);
      $this::test('Was it a car or a cat I saw?')->shouldReturn(true);
    }
}
