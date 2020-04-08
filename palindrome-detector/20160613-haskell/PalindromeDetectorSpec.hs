module PalindromeDetectorSpec where

import Test.Hspec
import PalindromeDetector

main :: IO ()
main = hspec $ do
  describe "PalindromeDetector" $ do
    it "detects a palindrome" $ do
      isPalindrome "racecar" `shouldBe` True
      isPalindrome "12321" `shouldBe` True
      isPalindrome "1AB232BA1" `shouldBe` True
      isPalindrome "foobar" `shouldBe` False

    it "normalises the string before attempting to detect a palindrome" $ do
      isPalindrome "A man, a plan, a canal, Panama!" `shouldBe` True
      isPalindrome "Was it a car or a cat I saw?" `shouldBe` True

    it "does not crash when passed an empty string" $
      isPalindrome "" `shouldBe` True
