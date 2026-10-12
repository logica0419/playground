-- 問1
maximum' :: (Ord a) => [a] -> a
maximum' [x] = x
maximum' (x : xs) = max x (maximum' xs)

-- 問2
reverse' :: [a] -> [a]
reverse' [] = []
reverse' (x : xs) = reverse' xs ++ [x]

-- 問3
quickSort :: (Ord a) => [a] -> [a]
quickSort [] = []
quickSort (pivot : remaining) = quickSort left ++ [pivot] ++ quickSort right
  where
    left = [x | x <- remaining, x <= pivot]
    right = [x | x <- remaining, x > pivot]

-- 問4
pow :: Integer -> Integer -> Integer
pow x n
  | x < 1 || n < 0 = error "Negative exponent"
  | n == 0 = 1
  | otherwise = x * pow x (n - 1)

-- 問5
quickPow :: Integer -> Integer -> Integer
quickPow x n
  | x < 1 || n < 0 = error "Negative exponent"
  | n == 0 = 1
  | even n = quickPow (x * x) (div n 2)
  | otherwise = x * quickPow x (n - 1)

-- 問6
tailPow :: Integer -> Integer -> Integer
tailPow x n
  | x < 1 || n < 0 = error "Negative exponent"
  | otherwise = iter 0 1
  where
    iter i acc
      | i == n = acc
      | otherwise = iter (i + 1) (x * acc)

-- 問7
myFList :: (a -> a) -> [a] -> [a]
myFList _ [] = []
myFList f (x : xs) = f x : myFList f xs

-- 問7 Alt
myFList' :: (a -> a) -> [a] -> [a]
myFList' f xs = [f x | x <- xs]

main :: IO ()
main = do
  print $ tailPow 2 10
  print $ myFList (pow 2) [1, 2, 3]
  print $ myFList' (pow 2) [1, 2, 3]
