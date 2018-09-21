-- Part 1: Hailstone Length

hailLen n 
  | n == 1    = 0
  | otherwise = 1 + hailLen (hailstone n) 

hailstone n
  | even n    = n `div` 2
  | otherwise = 3*n+1

-- Part 2: Divisors and Primes
divisors :: Int -> [Int]
divisors n = [i | i <- [2..(n `div` 2)], n `mod` i == 0]

primes :: Int -> [Int]
primes n = [i | i <- [2..n], divisors i == []]

-- Part 3: Joining Strings

join sep list(x:xs) = x