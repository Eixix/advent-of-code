import Control.Monad ()
import Data.Char qualified as Char
import System.IO
  ( IOMode (ReadMode),
    hClose,
    hGetContents,
    openFile,
  )

-- 5 and 8 for small.txt
-- 10 and 25 for challenge.txt
getWinningLineAndCard :: [String] -> (Int, [Int], [Int])
getWinningLineAndCard xs = (1, map read (take 10 xs), map read (take 25 (reverse xs)))

numberOfOverlaps :: (Num a1, Foldable t, Eq a2) => t a2 -> [a2] -> a1
numberOfOverlaps xs [] = 0
numberOfOverlaps xs (y:ys) = if y `elem` xs then 1 + numberOfOverlaps xs ys else numberOfOverlaps xs ys

addDrawsToArray :: Int -> [(Int, [Int], [Int])] -> [(Int, [Int], [Int])]
addDrawsToArray _ [] = []
addDrawsToArray n all@((count, winners, draws):xs)
  | n <= 0 = all
  | otherwise =  (count+1, winners, draws) : addDrawsToArray (n-1) xs

callGame :: [(Int, [Int], [Int])] -> Int
callGame [] = 0
callGame ((count, winners, draws):xs) = count + callGame (iterate (addDrawsToArray (numberOfOverlaps winners draws)) xs !! count)

main :: IO ()
main = do
  handle <- openFile "challenge.txt" ReadMode
  contents <- hGetContents handle
  let singlelines = lines contents
      winnersAndDraws = map (getWinningLineAndCard . drop 2 . words) singlelines
      result = callGame winnersAndDraws
  print result
  hClose handle