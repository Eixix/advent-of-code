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
getWinningLineAndCard :: [String] -> ([Int], [Int])
getWinningLineAndCard xs = (map read (take 10 xs), map read (take 25 (reverse xs)))

numberOfOverlaps :: (Num a1, Foldable t, Eq a2) => t a2 -> Bool -> [a2] -> a1
numberOfOverlaps xs isWinner [] = if isWinner then 1 else 0
numberOfOverlaps xs isWinner (y:ys) = if y `elem` xs then 2 * numberOfOverlaps xs True ys else numberOfOverlaps xs (isWinner || False) ys

getPointsPerGame :: (Integral a, Foldable t, Eq a2) => (t a2, [a2]) -> a
getPointsPerGame (winners, draws) = div (numberOfOverlaps winners False draws) 2

main :: IO ()
main = do
  handle <- openFile "challenge.txt" ReadMode
  contents <- hGetContents handle
  let singlelines = lines contents
      winnersAndDraws = map (getWinningLineAndCard . drop 2 . words) singlelines
      result = sum $ map getPointsPerGame winnersAndDraws
  print result
  hClose handle