import Data.List.Split (splitOn)
import System.IO
  ( IOMode (ReadMode),
    hClose,
    hGetContents,
    openFile,
  )

gauss :: (Fractional a1, Integral a2) => a2 -> a1
gauss x = (y * y + y) / 2
  where
    y = fromIntegral x

main :: IO ()
main = do
  let list = []
  handle <- openFile "input.csv" ReadMode
  contents <- hGetContents handle
  let singlelines = contents
      list = map (read :: String -> Int) (splitOn "," singlelines)
  print $ minimum $ map (\x -> foldl (\y z -> y + gauss (abs (x - z))) 0 list) list
  hClose handle