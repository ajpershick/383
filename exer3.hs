import Data.Time.Calendar
import Data.Time.Calendar.OrdinalDate

-- Merge 
merge :: [Int] -> [Int] -> [Int]
merge x [] = x
merge [] y = y
merge (xi:xf) (yi:yf) 
    | xi <= yi = xi : merge xf (yi:yf)
    | xi > yi = yi : merge (xi:xf) yf

-- Tail Recursive Hailstone
hailLen n = hailTail a n
    where
        hailTail a 1 = a
        hailTail a n = hailTail (3n + 1) (n `div` 2)

stupidAdder x 0 = x
stupidAdder x y = stupidAdder (x+1) (y-1)

-- Factorial recursively
fact :: Int -> Int
fact 0 = 1
fact x = x * fact (x-1) 

-- Factorial using foldl
fact' :: Int -> Int
fact' x = foldl (*) 1 [1..x]

-- Haskell Library and Dates
daysInYear :: Integer -> [Day]
daysInYear y = [jan1..dec31]
  where jan1 = fromGregorian y 1 1
        dec31 = fromGregorian y 12 31

-- isFriday function
isFriday :: Day -> Bool        
isFriday day 
    | snd (mondayStartWeek day) == 5 = True
    | otherwise = False

-- isPrimeDay function
isPrimeDay :: Day -> Bool
isPrimeDay day
    | divisors(getDay(toGregorian day)) == [] = True
    | otherwise = False

-- Helper functions for isPrimeDay
getDay (y,m,d) = d

divisors :: Int -> [Int]
divisors n = [i | i <- [2..(n `div` 2)], n `mod` i == 0]


