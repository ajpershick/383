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

join sep [] = ""
join sep [x] = x
join sep (x:xs) = x ++ sep ++ join sep xs

-- Part 4: Pythagorean Triples

pythagorean :: Int -> [(Int, Int, Int)]
pythagorean 0..4 = 0
pythagorean x = 