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

pruneSymbol :: [Char] -> [Char]
pruneSymbol = filter (\x -> x /= ',' && x /= ';' && x /= ':')

getId :: [a] -> a
getId xs = xs !! 1

isValidDraw :: (Ord a, Num a) => a -> String -> Bool
isValidDraw n xs
  | n <= 12 && xs == "red" = True
  | n <= 13 && xs == "green" = True
  | n <= 14 && xs == "blue" = True
  | otherwise = False

validateGame :: [[Char]] -> String -> Int
validateGame [] id = read id :: Int
validateGame (x : xs) id
  | isDigit (head x) = if isValidDraw (read x :: Int) (pruneSymbol (head xs)) then validateGame (tail xs) id else 0
  | otherwise = validateGame (tail xs) id

main :: IO ()
main = do
  handle <- openFile "challenge.txt" ReadMode
  contents <- hGetContents handle
  let singlelines = lines contents
      result = sum $ map ((\x -> validateGame x (getId x)) . words . pruneSymbol) singlelines
  print result
  hClose handle