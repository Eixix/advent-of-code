
import System.IO
    ( hClose, openFile, hGetContents, IOMode(ReadMode) )  
import Control.Monad ()
import qualified Data.Char as Char

findPair :: (Num a, Eq a) => [a] -> [a] -> a
findPair [] _ = 0
findPair (x:xs) [] = findPair xs (x:xs)
findPair (x:xs) (y:ys) = if x + y == 2020 then x * y else findPair (x:xs) ys

main :: IO ()
main = do  
        handle <- openFile "small.txt" ReadMode
        contents <- hGetContents handle
        let singlelines = lines contents
            numberList = map (read)
            result = findPair singlelines singlelines
        print result
        hClose handle   