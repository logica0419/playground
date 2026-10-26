-- 問1

zipWith' :: (a -> b -> c) -> [a] -> [b] -> [c]
zipWith' _ [] _ = []
zipWith' _ _ [] = []
zipWith' f (x : xs) (y : ys) = f x y : zipWith' f xs ys

-- 問2

map' :: (a -> b) -> [a] -> [b]
map' f = foldr (\x acc -> f x : acc) []

filter' :: (a -> Bool) -> [a] -> [a]
filter' f = foldr (\x acc -> if f x then x : acc else acc) []

-- 問3

oddSquareSum :: Integer
oddSquareSum = sum . takeWhile (< 10000) . map (^ 2) . filter odd $ [1 ..]

-- 問4

deriv :: (Double -> Double) -> Double -> Double
deriv f = (\x -> (f (x + h) - f x) / h)
  where
    h = 0.0001

-- 問5

newton :: (Fractional a) => (a -> a) -> a -> a
newton f x 
