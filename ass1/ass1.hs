import RainbowAssign
import qualified Data.Map as Map

pwLength, nLetters, width, height :: Int
filename :: FilePath
pwLength = 8            -- length of each password
nLetters = 5            -- number of letters to use in passwords: 5 -> a-e
width = 40              -- length of each chain in the table
height = 1000           -- number of "rows" in the table
filename = "table.txt"  -- filename to store the table

-- Reducing the Hash Values to Passwords
pwReduce :: Hash -> Passwd
pwReduce x = map toLetter $ reverse $ map ( `mod` nLetters ) $ take pwLength $ iterate ( `div` nLetters ) $ fromEnum x 

-- Building the Rainbow Table
rainbowTable :: Int -> [Passwd] -> Map.Map Hash Passwd
generateHashList passwords = iterate (map pwHash . map pwReduce) $ map pwHash passwords 
getHashRow width passwords =  take (width + 1) (generateHashList passwords) !!width
rainbowTable width passwords = Map.fromList $ zip (getHashRow width passwords) passwords






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