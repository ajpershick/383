hailstone n
  | even n    = n `div` 2
  | otherwise = 3*n+1

-- Recursive Hailstone Generator
hailSeq :: Int -> [Int]
hailSeq n = hailTail [] n 
    where
        hailTail list 1 = list ++ [1]
        hailTail list n = hailTail (list ++ [n]) (hailstone n)


-- Iterative Hailstone Generator
hailSeq' :: Int -> [Int]
hailSeq' n = takeWhile(\x -> x /= 1) (iterate hailstone n) ++ [1]