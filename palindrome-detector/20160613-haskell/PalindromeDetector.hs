module PalindromeDetector (isPalindrome) where

import Data.Char (toLower)

normalise :: String -> String
normalise x = [y | y <- map toLower x, y `elem` ['a'..'z'] ++ ['0'..'9']]

isPalindrome :: String -> Bool
isPalindrome x = reverse (normalise x) == normalise x
