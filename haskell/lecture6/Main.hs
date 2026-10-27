{-# OPTIONS_GHC -Wall #-}

import Data.Char qualified as Char (digitToInt, isDigit)
import Data.List qualified as List (findIndices, group, isPrefixOf, sort, tails, words)
import Data.Map qualified as Map (fromListWith, toList)
import Shape.Circle qualified as Circle (area)
import Shape.Rectangle qualified as Rectangle (area)

-- 問1

rectAreaFor :: (Num a) => [(a, a)] -> [a]
rectAreaFor = map (\(w, h) -> Rectangle.area w h)

circleAreaFor :: (Floating a) => [a] -> [a]
circleAreaFor = map Circle.area

-- 問2

duplicateWordCount :: String -> [(Int, String)]
duplicateWordCount =
  map
    ( \w -> case w of
        (x : _) -> (length w, x)
        [] -> error "unreachable"
    )
    . filter (\w -> length w > 1)
    . List.group
    . List.sort
    . List.words

-- 問3

infixHeadIndexList :: (Eq a) => [a] -> [a] -> [Int]
infixHeadIndexList part = List.findIndices (List.isPrefixOf part) . List.tails

-- 問4

firstTo' :: Int -> Integer
firstTo' n
  | n < 1 = error "n must be greater than 0"
  | n < 10 = toInteger n
  | otherwise = firstTo' (n - 9) * 10 + 9

-- 問5

string2digits :: String -> [Int]
string2digits = map Char.digitToInt . filter Char.isDigit

-- 問6

wordNums :: String -> [(String, Int)]
wordNums = Map.toList . Map.fromListWith (+) . map (\w -> (w, 1)) . List.words

-- 問7

dotProduct :: (Num a) => [a] -> [a] -> a
dotProduct [] [] = 0
dotProduct [] _ = error "either list must be non-empty"
dotProduct _ [] = error "either list must be non-empty"
dotProduct (x : xs) (y : ys) = x * y + dotProduct xs ys

-- 問8

dotProduct2 :: (Num a) => [a] -> [a] -> a
dotProduct2 a b = dotProductTail a b 0
  where
    dotProductTail [] [] acc = acc
    dotProductTail [] _ _ = error "either list must be non-empty"
    dotProductTail _ [] _ = error "either list must be non-empty"
    dotProductTail (x : xs) (y : ys) acc = dotProductTail xs ys (acc + x * y)

-- 問9

dotProduct3 :: (Num a) => [a] -> [a] -> a
dotProduct3 a b = dotProductTail a b 0
  where
    dotProductTail [] [] acc = acc
    dotProductTail [] _ _ = error "either list must be non-empty"
    dotProductTail _ [] _ = error "either list must be non-empty"
    dotProductTail (x : xs) (y : ys) acc =
      let acc' = acc + x * y
       in acc' `seq` dotProductTail xs ys acc'
