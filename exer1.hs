det a b c = b^2 - 4*a*c
quadsol1 a b c = (-b - sqrt (det a b c))/2*a
quadsol2 a b c = (-b + sqrt (det a b c))/2*a

-- Solutions to exercise --

third_a list = list !! 2

third_b (_:_:x:xs) = x
