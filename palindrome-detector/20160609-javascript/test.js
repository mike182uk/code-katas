const test = require('tape')

const isPalindrome = require('./')

test('it can detect a palindrome', assert => {
  assert.plan(4)

  assert.equals(isPalindrome('racecar'), true, 'it can detect a word palindrome')
  assert.equals(isPalindrome('12321'), true, 'it can detect a number palindrome')
  assert.equals(isPalindrome('1AB232BA1'), true, 'it can detect a mixed character palindrome')
  assert.equals(isPalindrome('foobar'), false, 'it can detect no palindrome')

  assert.end()
})

test('it removes capital letters, punctuation, and word dividers before attempting to detect a palindrome', assert => {
  assert.plan(1)

  assert.equals(isPalindrome('A man, a plan, a canal, Panama!'), true, 'capital letters, punctuation and word dividers are removed')

  assert.end()
})

test('it throws an error if anything other than a string is passed as an argument', assert => {
  assert.plan(4)

  let err = /Supplied palindrome is not a string/

  assert.throws(() => isPalindrome([]), err, 'error is thrown if an array is passed as an argument')
  assert.throws(() => isPalindrome({}), err, 'error is thrown if an object is passed as an argument')
  assert.throws(() => isPalindrome(123), err, 'error is thrown if an integer is passed as an argument')
  assert.doesNotThrow(() => isPalindrome('racecar'), 'error is not thrown if a string is passed as an argument')

  assert.end()
})
