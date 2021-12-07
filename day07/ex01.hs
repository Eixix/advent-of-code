import System.IO
    ( hClose, openFile, hGetContents, IOMode(ReadMode) )
import Data.List.Split ( splitOn )

quicksort :: Ord a => [a] -> [a]
quicksort [] = []
quicksort (p:xs) = lesser ++ p:greater
  where
    lesser = quicksort $ filter (< p) xs
    greater = quicksort $ filter (>= p) xs

gauss :: (Fractional a1, Integral a2) => a2 -> a1
gauss x = (y * y + y) / 2
    where
        y = fromIntegral x

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
        print $ foldl (\x y -> x + gauss (abs (y - medianElement))) 0 list
        hClose handle