import System.IO
    ( hClose, openFile, hGetContents, IOMode(ReadMode) )
import Data.List.Split ( splitOn )

lanternfish :: (Enum b, Foldable t, Eq b, Num b) => t b -> [b]
lanternfish xs = map pred $ foldl (\xs x -> if x == 0 then 7:9:xs else x:xs) [] xs


main :: IO ()
main = do
        let list = []
        handle <- openFile "input.csv" ReadMode
        contents <- hGetContents handle
        let singlelines = contents
            list = map (read :: String -> Int) (splitOn "," singlelines)
        print $ length $ last $ take 257 $ iterate lanternfish list
        hClose handle
