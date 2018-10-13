-- Recursive Iterate
myIterate :: (a -> a) -> a -> [a]
myIterate f x = x : myIterate f (f x)

-- Recursive TakeWhile
myTakeWhile :: (Eq a) => [a] -> [a] -> [a]
myTakeWhile cond list
    | cond == (take 1 list) = myTakeWhile cond list
    | otherwise = myTakeWhile cond (drop 1 list)

-- Pascal's Triangle
pascal :: Int -> [Int]
pascal 0 = [1]
pascal 1 = [1,1]
pascal n = [1] ++ map (uncurry (+)) (zip (pascal (n-1)) (tail (pascal (n-1)))) ++ [1]

-- Pointfree Addition
addPair :: (Int, Int) -> Int
addPair = uncurry (+)

-- Pointfree Filtering
withoutZeros :: (Eq a, Num a) => [a] -> [a]
withoutZeros = filter (/=0)

-- Exploring Fibonacci
fib :: Int -> Int
fib 0 = 0
fib 1 = 1
fib n = fib (n-1) + fib (n-2)

-- Infinite sequence
fibs = map fib [0..]

-- Something else
things :: [Integer]
things = 0 : 1 : zipWith (+) things (tail things)