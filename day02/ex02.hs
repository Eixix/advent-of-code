import System.IO
    ( hClose, openFile, hGetContents, IOMode(ReadMode) )  
import Control.Monad ()
import qualified Data.Char as Char

data Direction = Forward | Down | Up deriving (Eq, Show, Read)


calculateAdvanced :: Num a => a -> a -> a -> [(Direction, a)] -> a
calculateAdvanced _ d p [] = d * p
calculateAdvanced a d p ((x,y):xs) 
  | x == Down       = calculateAdvanced (a + y) d p xs
  | x == Up         = calculateAdvanced (a - y) d p xs
  | x == Forward    = calculateAdvanced a (d + (a * y)) (p + y) xs

tuplify2 :: [a] -> (a,a)
tuplify2 [x,y] = (x,y)

capitalized :: String -> String
capitalized (head:tail) = Char.toUpper head : map Char.toLower tail
capitalized [] = []

main :: IO ()
main = do  
        let list = []
        handle <- openFile "input.csv" ReadMode
        contents <- hGetContents handle
        let singlelines = lines contents
            list = map ((\(x,y) -> (read x :: Direction, read y :: Integer)) . (\(x,y) -> (capitalized x, y)) . tuplify2 . words) singlelines
        print list
        print $ calculateAdvanced 0 0 0 list
        hClose handle   
