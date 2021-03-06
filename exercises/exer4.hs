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

-- Join using fold
join :: String -> [String] -> String
join sep [] = ""
join sep [x] = x
join sep list = foldl1 (++) (map (\x -> x ++ sep) (init list)) ++ (last list)

-- Merge Sort
merge :: Ord a => [a] -> [a] -> [a]
merge x [] = x
merge [] y = y
merge (xi:xf) (yi:yf) 
    | xi <= yi = xi : merge xf (yi:yf)
    | xi > yi = yi : merge (xi:xf) yf

mergeSort :: Ord a => [a] -> [a]
mergeSort [] = []
mergeSort [x] = [x]
mergeSort list = merge (mergeSort left) (mergeSort right)
    where 
        left = take split list
        right = drop split list
        split = length list `div` 2

-- Maybe
maybeToList :: Maybe a -> [a]
maybeToList Nothing  = []
maybeToList (Just x) = [x]
-- findElt :: a -> Maybe a
findElt x list = do 
   case x of
    Nothing -> x:list
    Just x -> list
