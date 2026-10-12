specialOnSeven :: Int -> String
specialOnSeven 7 = "It's seven!"
specialOnSeven x = "Not seven, it's " ++ show x

recFactorial :: Int -> Int
recFactorial n | n < 0 = -1
recFactorial 0 = 1
recFactorial n = n * recFactorial (n - 1)

tellBMICategory :: Double -> Double -> String
tellBMICategory weight height
  | bmi < 18.5 = mesPrefix ++ "Underweight"
  | bmi < 25.0 = mesPrefix ++ "Healthy weight"
  | bmi < 30.0 = mesPrefix ++ "Overweight"
  | bmi < 35 = mesPrefix ++ "Class 1 obesity"
  | bmi < 40 = mesPrefix ++ "Class 2 obesity"
  | otherwise = mesPrefix ++ "Class 3 obesity"
  where
    bmi = weight / ((height / 100) ^ 2)
    mesPrefix = "BMI: " ++ show bmi ++ " / "

main :: IO ()
main = do
  let arg = 7
  print $ specialOnSeven arg
  print $ recFactorial 3
  print $ recFactorial (-4)
  print $ tellBMICategory 61 172
