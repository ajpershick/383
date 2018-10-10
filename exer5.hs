-- Recursive Iterate
myIterate :: (a -> a) -> a -> [a]
myIterate f x = x : myIterate f (f x)

-- Recursive TakeWhile
myTakeWhile :: (Eq a) => [a] -> [a] -> [a]
myTakeWhile cond list
    | cond == (take 1 list) = myTakeWhile cond list
    | otherwise = myTakeWhile cond (drop 1 list)

-- Pascal's Triangle


pascal prev = map $ uncurry (+) (zip prev (tail prev))