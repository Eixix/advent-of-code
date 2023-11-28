
import System.IO
    ( hClose, openFile, hGetContents, IOMode(ReadMode) )  
import Control.Monad ()
import qualified Data.Char as Char

stringToInt :: String -> Int
stringToInt = read

findPair :: (Num a, Eq a) => [a] -> [a] -> a
findPair [] _ = 0
findPair xs (y:ys) = if head xs + y == 2020 then head xs*y else findPair xs ys
findPair (x:xs) [] = findPair xs (x:xs)

main :: IO ()
main = do  
        handle <- openFile "challenge.txt" ReadMode
        contents <- hGetContents handle
        let singlelines = lines contents
            numberList = map stringToInt singlelines
            result = findPair numberList numberList
        print result
        hClose handle   