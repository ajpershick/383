-- Merge 
merge x [] = x
merge [] y = y
merge (xi:xf) (yi:yf) 
    | xi <= yi = xi : merge xf (yi:yf)
    | xi > yi = yi : merge (xi:xf) yf

-- Tail Recursive Hailstone
-- hailLen n = hailTail a n
--     where
--         hailTail a 1 = a
--         hailTail a n = hailTail (3n + 1) (n `div` 2)

-- stupidAdder x 0 = x
-- stupidAdder x y = stupidAdder (x+1) (y-1)

-- Factorial

-- Recursively
fact 0 = 1
fact x = x * fact (x-1) 

-- Using foldl
fact' x = foldl (*) 1 [1..x]
