hailSeq :: Int -> [Int]
hailSeq n = hailTail n []
    where
        hailTail 1 list = list
        hailTail n list = list ++ (hailTail (hailstone n) list )

hailstone n
  | even n    = n `div` 2
  | otherwise = 3*n+1