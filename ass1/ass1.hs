import RainbowAssign


pwLength, nLetters, width, height :: Int
filename :: FilePath
pwLength = 8            -- length of each password
nLetters = 5            -- number of letters to use in passwords: 5 -> a-e
width = 40              -- length of each chain in the table
height = 1000           -- number of "rows" in the table
filename = "table.txt"  -- filename to store the table

listOfDivisions x = take pwLength (reverse (iterate (`div` nLetters) x))

-- listOfMods :: [Int] -> [Int]
-- listOfMods x = map reverse x




-- While i < pwLength
--     answer = (hash mod nLetters) ++ answer
--     hash = hash / nLetters
--     i++
-- return answer
    

-- 345298305   3
-- 69059661    0
-- 13811932    1
-- 2762386     2
-- 552477      1
-- 110495      2
-- 22099       0
-- 4419        4