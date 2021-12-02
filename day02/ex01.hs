import System.IO
    ( hClose, openFile, hGetContents, IOMode(ReadMode) )  
import Control.Monad ()
import qualified Data.Char as Char

data Direction = Forward | Down | Up deriving (Eq, Show, Read)

calculateDepth :: Num a => [(Direction, a)] -> a
calculateDepth [] = 0
calculateDepth ((x,y):xs) 
  | x == Down    = calculateDepth xs + y
  | x == Up      = calculateDepth xs - y
  | otherwise    = calculateDepth xs

calculatePosition :: Num p => [(Direction, p)] -> p
calculatePosition [] = 0
calculatePosition ((x,y):xs)
  | x == Forward = y + calculatePosition xs
  | otherwise    = calculatePosition xs

tuplify2 :: [a] -> (a,a)
tuplify2 [x,y] = (x,y)

capitalized :: String -> String
capitalized (head:tail) = Char.toUpper head : map Char.toLower tail
capitalized [] = []

main :: IO ()
main = do  
        let list = []
        handle <- openFile "input01.csv" ReadMode
        contents <- hGetContents handle
        let singlelines = lines contents
            list = map ((\(x,y) -> (read x :: Direction, read y :: Integer)) . (\(x,y) -> (capitalized x, y)) . tuplify2 . words) singlelines
        print list
        print $ calculatePosition list
        print $ calculateDepth list
        print $ calculatePosition list * calculateDepth list
        hClose handle   
