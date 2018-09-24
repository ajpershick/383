-- Merge 
merge x [] = x
merge [] y = y
merge (xi:xf) (yi:yf) 
    | xi <= yi = xi : merge xf (yi:yf)
    | xi > yi = yi : merge (xi:xf) yf

-- Tail Recursive Hailstone
hailLen n = hailTail a n
    where
        hailTail a 0 = a
        hailTail a n = hailTail 

-- stupidAdder x 0 = x
-- stupidAdder x y = stupidAdder (x+1) (y-1)

-- Factorial