import Control.Monad ()
import Data.Char qualified as Char
import System.IO
  ( IOMode (ReadMode),
    hClose,
    hGetContents,
    openFile,
  )
import Data.List (elemIndex)
import Data.Maybe (fromJust)
import qualified Data.Bifunctor
import Debug.Trace

isDigit :: Char -> Bool
isDigit x
  | x == '0' = True
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

splitInIndividualLists :: [String] -> [[String]]
splitInIndividualLists [] = []
splitInIndividualLists xs = takeWhile (/= []) xs : splitInIndividualLists (drop 1 (dropWhile (/= "") xs))

getEntries :: [String] -> [Int]
getEntries xs = filter (/= -1) (map (\x -> if isDigit (head x) then (read x :: Int) else -1) xs)

convertToRange :: [Int] -> ([Int], [Int])
convertToRange (destination:source:range) = ([destination..(destination + (head range-1))], [source..(source + (head range-1))])

mergeRanges :: [([Int], [Int])] -> ([Int], [Int])
mergeRanges = foldr
      (\ x -> Data.Bifunctor.bimap (fst x ++) (snd x ++)) ([], [])

convertToRanges :: [[Int]] -> [([Int], [Int])]
convertToRanges = map convertToRange

getDestinationFromSource :: Int -> ([Int], [Int]) -> Int
getDestinationFromSource source (destinations, sources) = if source `elem` sources then destinations !! fromJust (elemIndex source sources) else source

mapSeed :: [([Int], [Int])] -> Int -> Int
mapSeed maps seed = foldl getDestinationFromSource seed maps


main :: IO ()
main = do
  handle <- openFile "challenge.txt" ReadMode
  contents <- hGetContents handle
  let singlelines = lines contents
      filteredElements = map (filter (not . null) . map (getEntries . words)) (splitInIndividualLists singlelines)
      seeds = head $ head filteredElements
      maps = map (mergeRanges . convertToRanges) (tail filteredElements)
  --print $ minimum (map (mapSeed maps) seeds)
  print $ map (mapSeed maps) seeds
  hClose handle