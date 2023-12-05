import Control.Monad ()
import Data.Char qualified as Char
import System.IO
  ( IOMode (ReadMode),
    hClose,
    hGetContents,
    openFile,
  )

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

splitInIndividualLists :: [String] -> [[String]]
splitInIndividualLists [] = []
splitInIndividualLists xs = takeWhile (/= []) xs : splitInIndividualLists (drop 1 (dropWhile (/= "") xs))

getEntries :: [String] -> [Int]
getEntries xs = filter (/= 0) (map (\x -> if isDigit (head x) then (read x :: Int) else 0) xs)

main :: IO ()
main = do
  handle <- openFile "small.txt" ReadMode
  contents <- hGetContents handle
  let singlelines = lines contents
      filteredElements = map (filter (not . null) . map (getEntries . words)) (splitInIndividualLists singlelines)
  print filteredElements
  hClose handle