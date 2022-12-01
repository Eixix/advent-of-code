import System.IO
    ( hClose, openFile, hGetContents, IOMode(ReadMode) )
import Data.List.Split ( splitOn )

quicksort :: Ord a => [a] -> [a]
quicksort [] = []
quicksort (p:xs) = lesser ++ p:greater
  where
    lesser = quicksort $ filter (< p) xs
    greater = quicksort $ filter (>= p) xs

main :: IO ()
main = do
        let list = []
        handle <- openFile "input.csv" ReadMode
        contents <- hGetContents handle
        let singlelines = contents
            list = map (read :: String -> Int) (splitOn "," singlelines)
            sorted = quicksort list
            listLength = length list
            position = listLength `div` 2
            medianElement = sorted !! position
        print $ foldl (\x y -> x + abs (y - medianElement)) 0 list
        hClose handle