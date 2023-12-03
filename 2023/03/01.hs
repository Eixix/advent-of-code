import Control.Monad ()
import Data.Char qualified as Char
import System.IO
  ( IOMode (ReadMode),
    hClose,
    hGetContents,
    openFile,
  )

isSymbol :: Char -> Bool
isSymbol x
  | isDigit x = False
  | x == '.' = False
  | otherwise = True

customHeadEmptyTuple :: [(a, Char)] -> Char
customHeadEmptyTuple [] = '.'
customHeadEmptyTuple ((_,b):xs) = b

isSymbolAtIndex :: Eq a => [(a, Char)] -> a -> Bool
isSymbolAtIndex xs n = isSymbol (customHeadEmptyTuple (filter (\x -> fst x == n) xs))

mapToIndexedTupleList :: [b] -> [(Int, b)]
mapToIndexedTupleList = zip [1..]

isDigit :: Char -> Bool
isDigit x
  | x == '1' = True
  | x == '2' = True
  | x == '3' = True
  | x == '4' = True
  | x == '5' = True
  | x == '6' = True
  | x == '7' = True
  | x == '8' = True
  | x == '9' = True
  | otherwise = False

isSymbolAroundNumber :: [(Int, Char)] -> Int -> Int -> Bool
isSymbolAroundNumber xs lineCount n = isSymbolAtIndex xs (n-lineCount-1) || isSymbolAtIndex xs (n-lineCount) || isSymbolAtIndex xs (n-lineCount+1)|| isSymbolAtIndex xs (n-1)|| isSymbolAtIndex xs (n+1)|| isSymbolAtIndex xs (n+lineCount-1)|| isSymbolAtIndex xs (n+lineCount)|| isSymbolAtIndex xs (n+lineCount-1)

onlySymbolNumbers :: [(Int, Char)] -> [(Int, Char)] -> Int -> Int
onlySymbolNumbers _ [] _ = 0
onlySymbolNumbers all ((a, x1):(b, x2):(c, x3):xs) lineCount
  | isDigit x1 && isDigit x2 && isDigit x3 = if any (isSymbolAroundNumber all lineCount) [a,b,c] then (read (x1:x2:[x3]) :: Int) + onlySymbolNumbers all xs lineCount else onlySymbolNumbers all xs lineCount
  | isDigit x1 && isDigit x2 = if any (isSymbolAroundNumber all lineCount) [a,b] then (read (x1:[x2]) :: Int) + onlySymbolNumbers all xs lineCount else onlySymbolNumbers all xs lineCount
  | isDigit x1 = if isSymbolAroundNumber all lineCount a then (read [x1] :: Int) + onlySymbolNumbers all xs lineCount else onlySymbolNumbers all xs lineCount
  | otherwise = onlySymbolNumbers all xs lineCount
onlySymbolNumbers all ((a, x1):(b, x2):xs) lineCount
  | isDigit x1 && isDigit x2 = if any (isSymbolAroundNumber all lineCount) [a,b] then (read (x1:[x2]) :: Int) + onlySymbolNumbers all xs lineCount else onlySymbolNumbers all xs lineCount
  | isDigit x1 = if isSymbolAroundNumber all lineCount a then (read [x1] :: Int) + onlySymbolNumbers all xs lineCount else onlySymbolNumbers all xs lineCount
  | otherwise = onlySymbolNumbers all xs lineCount
onlySymbolNumbers all ((a, x1):xs) lineCount
  | isDigit x1 = if isSymbolAroundNumber all lineCount a then (read [x1] :: Int) + onlySymbolNumbers all xs lineCount else onlySymbolNumbers all xs lineCount
  | otherwise = onlySymbolNumbers all xs lineCount

main :: IO ()
main = do
  handle <- openFile "small.txt" ReadMode
  contents <- hGetContents handle
  let singlelines = filter (/= '\n') contents
      lineCount = length singlelines
      tupleList = mapToIndexedTupleList singlelines
      result = onlySymbolNumbers tupleList tupleList lineCount
  print result
  hClose handle