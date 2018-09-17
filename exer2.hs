hailstone n
  | even n    = n `div` 2
  | otherwise = 3*n+1
