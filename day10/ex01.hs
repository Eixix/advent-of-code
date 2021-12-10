import Control.Monad ()
import System.IO
  ( IOMode (ReadMode),
    hClose,
    hGetContents,
    openFile,
  )

initBrackets :: Num a => [Char] -> a
initBrackets [] = error "Obacht"
initBrackets (x : xs) = brackets [x] xs 0

brackets :: Num a => [Char] -> [Char] -> a -> a
brackets _ [] a = a
brackets [] (y : ys) a
  | y == '(' = brackets [y] ys a
  | y == '[' = brackets [y] ys a
  | y == '{' = brackets [y] ys a
  | y == '<' = brackets [y] ys a
  | y == ')' = 3
  | y == ']' = 57
  | y == '}' = 1197
  | y == '>' = 25137
brackets (x : xs) (y : ys) a
  | y == '(' = brackets (y : x : xs) ys a
  | y == '[' = brackets (y : x : xs) ys a
  | y == '{' = brackets (y : x : xs) ys a
  | y == '<' = brackets (y : x : xs) ys a
  | y == ')' && x /= '(' = a + 3
  | y == ']' && x /= '[' = a + 57
  | y == '}' && x /= '{' = a + 1197
  | y == '>' && x /= '<' = a + 25137
  | otherwise = brackets xs ys a

main :: IO ()
main = do
  let list = []
  handle <- openFile "input.csv" ReadMode
  contents <- hGetContents handle
  let singlelines = lines contents
  print $ sum $ map initBrackets singlelines
  hClose handle