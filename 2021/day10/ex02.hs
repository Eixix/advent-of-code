import Control.Monad ()
import System.IO
  ( IOMode (ReadMode),
    hClose,
    hGetContents,
    openFile,
  )

quicksort :: Ord a => [a] -> [a]
quicksort [] = []
quicksort (p : xs) = lesser ++ p : greater
  where
    lesser = quicksort $ filter (< p) xs
    greater = quicksort $ filter (>= p) xs

initBrackets :: Num a => [Char] -> a
initBrackets [] = error "Obacht"
initBrackets (x : xs) = brackets [x] xs 0

brackets :: Num a => [Char] -> [Char] -> a -> a
brackets [] [] a = a
brackets (xs) [] a
  | last xs == '(' = brackets (init xs) [] a * 5 + 1
  | last xs == '[' = brackets (init xs) [] a * 5 + 2
  | last xs == '{' = brackets (init xs) [] a * 5 + 3
  | last xs == '<' = brackets (init xs) [] a * 5 + 4
brackets [] (y : ys) a = brackets [y] ys a
brackets (x : xs) (y : ys) a
  | y == '(' = brackets (y : x : xs) ys a
  | y == '[' = brackets (y : x : xs) ys a
  | y == '{' = brackets (y : x : xs) ys a
  | y == '<' = brackets (y : x : xs) ys a
  | y == ')' && x /= '(' = 0
  | y == ']' && x /= '[' = 0
  | y == '}' && x /= '{' = 0
  | y == '>' && x /= '<' = 0
  | otherwise = brackets xs ys a

main :: IO ()
main = do
  let list = []
  handle <- openFile "input.csv" ReadMode
  contents <- hGetContents handle
  let singlelines = lines contents
      resultSorted = quicksort $ filter (/= 0) (map initBrackets singlelines)
  print resultSorted
  print $ length resultSorted
  print $ resultSorted !! (length resultSorted `div` 2)
  hClose handle