import Data.Char qualified as Char (chr, ord)

encode :: Int -> String -> String
encode offset = map (Char.chr . (+ 32) . (`mod` 95) . (+ offset) . (subtract 32) . Char.ord)

decode :: Int -> String -> String
decode offset = encode . negate $ offset
