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

appendDigits :: [Char] -> [Char]
appendDigits xs = head (filter isDigit xs) : [last (filter isDigit xs)]

main :: IO ()
main = do
  handle <- openFile "challenge.txt" ReadMode
  contents <- hGetContents handle
  let singlelines = lines contents
      result = sum $ map (stringToInt . appendDigits) singlelines
  print result
  hClose handle