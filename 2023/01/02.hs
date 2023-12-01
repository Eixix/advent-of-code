import Control.Monad ()
import Data.Char qualified as Char
import System.IO
  ( IOMode (ReadMode),
    hClose,
    hGetContents,
    openFile,
  )

stringToInt :: String -> Int
stringToInt = read

isDigit :: Char -> Bool
isDigit x
  | x == '1' = True
  | x == '2' = True
  | x == '3' = True
  | x == '4' = True
  | x == '5' = True
  | x == '6' = True
  | x == '7' = True
  | x == '8' = True
  | x == '9' = True
  | otherwise = False

toDigitWord :: [Char] -> Char
toDigitWord x
  | take 3 x == "one" = '1'
  | take 3 x == "two" = '2'
  | take 5 x == "three" = '3'
  | take 4 x == "four" = '4'
  | take 4 x == "five" = '5'
  | take 3 x == "six" = '6'
  | take 5 x == "seven" = '7'
  | take 5 x == "eight" = '8'
  | take 4 x == "nine" = '9'
  | otherwise = '0'

convertToPureDigits :: [Char] -> [Char]
convertToPureDigits [] = []
convertToPureDigits (x : xs)
  | isDigit x = x : convertToPureDigits xs
  | toDigitWord (take 5 (x : xs)) /= '0' = toDigitWord (take 5 (x : xs)) : convertToPureDigits xs
  | otherwise = convertToPureDigits xs

appendDigits :: [Char] -> [Char]
appendDigits xs = head (filter isDigit (convertToPureDigits xs)) : [last (filter isDigit (convertToPureDigits xs))]

main :: IO ()
main = do
  handle <- openFile "challenge.txt" ReadMode
  contents <- hGetContents handle
  let singlelines = lines contents
      result = sum $ map (stringToInt . appendDigits) singlelines
  print result
  hClose handle