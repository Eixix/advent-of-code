import Control.Monad ()
import System.IO
  ( IOMode (ReadMode),
    hClose,
    hGetContents,
    openFile,
  )

only4 :: [[Char]] -> [[Char]]
only4 [] = []
only4 (x : xs)
  | x /= "|" = only4 xs
  | otherwise = xs

countWords :: (Foldable t, Num p) => t a -> p
countWords xs
  | length xs == 2 = 1
  | length xs == 3 = 1
  | length xs == 4 = 1
  | length xs == 7 = 1
  | otherwise = 0

main :: IO ()
main = do
  let list = []
  handle <- openFile "input.csv" ReadMode
  contents <- hGetContents handle
  let singlelines = lines contents
      wordsList = map words singlelines -- [["Test"], ["Test", "|"]]
  print $ map only4 wordsList
  print $ sum (map (foldl (\x y -> x + countWords y) 0 . only4) wordsList)
  hClose handle