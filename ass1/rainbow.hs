import RainbowAssign
import Data.List()
import qualified Data.Map as Map

pwReduce :: Hash -> Passwd
rainbowTable :: Int -> [Passwd] -> Map.Map Hash Passwd
generateHashList :: [Passwd] -> [[Hash]]
getHashRow :: Int -> [Passwd] -> [Hash]
generateTable :: IO ()
findPassword :: (Eq a) => Map.Map Hash a -> Int -> Hash -> Maybe a
findPasswordHelper :: (Eq a) => Map.Map Hash a -> Int -> Hash -> Int -> Maybe a
-- generateHashes :: Int -> Hash -> [Hash]
nextRow :: Hash -> Hash
test1 :: IO (Maybe Passwd)
-- test2 :: Int -> IO ([Passwd], Int)
main :: IO ()

main = do
  print "Test 1"
  generateTable
  res <- test1
  print res

pwLength, nLetters, width, height :: Int
filename :: FilePath
pwLength = 8            -- length of each password
nLetters = 5            -- number of letters to use in passwords: 5 -> a-e
width = 40              -- length of each chain in the table
height = 1000           -- number of "rows" in the table
filename = "table.txt"  -- filename to store the table

-- Reducing the Hash Values to Passwords
pwReduce x = map toLetter $ reverse $ map ( `mod` nLetters ) $ take pwLength $ iterate ( `div` nLetters ) $ fromEnum x 

-- Building the Rainbow Table
rainbowTable w passwords = Map.fromList $ zip (getHashRow w passwords) passwords

-- Helpers for Building Table
generateHashList passwords = iterate (map pwHash . map pwReduce) $ map pwHash passwords 
getHashRow w passwords =  take (w + 1) (generateHashList passwords) !!w

-- Writing table to text file
generateTable = do
  table <- buildTable rainbowTable nLetters pwLength width height
  writeTable table filename

-- Password Finding
findPassword table w x = findPasswordHelper table w x 0

findPasswordHelper table w x counter
  | counter == w = Nothing
  | Map.lookup x table == Nothing = findPasswordHelper table w (nextRow x) (counter + 1)
  | otherwise = Map.lookup x table

-- Finds the next row's hash value
nextRow x = (pwHash . pwReduce) x

-- generateHashes w x =  take w $ iterate nextRow x

  -- Testing 
test1 = do
    table <- readTable filename
    return (findPassword table 40 1726491528)

-- test2 n = do
--     table <- readTable filename
--     pws <- randomPasswords nLetters pwLength n
--     let hs = map pwHash pws
--     let result = Map.mapMaybe (findPassword table width) hs
--     return (result, length result)



    