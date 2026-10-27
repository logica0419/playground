module Shape.Rectangle (area, perimeter) where

area :: (Num a) => a -> a -> a
area width height = width * height

perimeter :: (Num a) => a -> a -> a
perimeter width height = (width + height) * 2
