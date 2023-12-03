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

validateGame :: [[Char]] -> Int -> Int -> Int -> Int
validateGame [] r g b = r*g*b
validateGame (x : xs) r g b
  | isDigit (head x) = case head xs of
      "red" -> validateGame (tail xs) (max (read x :: Int) r) g b
      "green" -> validateGame (tail xs) r (max (read x :: Int) g) b
      "blue" -> validateGame (tail xs) r g (max (read x :: Int) b)
      _ -> error "Should never happen"
  | otherwise = validateGame (tail xs) r g b

main :: IO ()
main = do
  handle <- openFile "challenge.txt" ReadMode
  contents <- hGetContents handle
  let singlelines = lines contents
      result = sum $ map ((\x -> validateGame x 0 0 0) . words . pruneSymbol) singlelines
  print result
  hClose handle