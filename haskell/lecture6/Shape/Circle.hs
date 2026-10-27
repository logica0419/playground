module Shape.Circle (area, perimeter) where

area :: (Floating a) => a -> a
area r = r ** 2 * pi

perimeter :: (Floating a) => a -> a
perimeter r = 2 * r * pi
