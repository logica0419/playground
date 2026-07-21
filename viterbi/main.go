package main

import "fmt"

// 指定された畳み込み符号のプログラミング的表現
func ConvolutionCode(reg uint8, inputBit uint8) ([2]uint8, uint8) {
	// レジスタの各ビットを取り出す
	reg1 := (reg >> 2) & 1
	reg2 := (reg >> 1) & 1
	reg3 := reg & 1

	// フロー図に従い、2bitの出力を計算
	out1 := inputBit ^ reg1 ^ reg2 ^ reg3
	out2 := inputBit ^ reg1 ^ reg3

	// 次のレジスタを計算
	nextReg := (inputBit << 2) | (reg1 << 1) | reg2

	return [2]uint8{out1, out2}, nextReg
}

// ハミング距離の計算
func HammingDistance(left []uint8, right []uint8) int {
	if len(left) != len(right) {
		panic("ハミング距離を計算する配列の長さが異なります")
	}

	// 1bitずつ見比べ、異なるビットの数を数える
	distance := 0
	for index := range left {
		if left[index] != right[index] {
			distance++
		}
	}

	return distance
}

// 入力ビット列の畳み込み符号化
func ConvolutionEncode(inputBits []uint8) []uint8 {
	// 出力ビット列とレジスタの初期化
	encodedBits := make([]uint8, 0, len(inputBits)*2)
	reg := uint8(0)

	// 各入力bitを順番に符号化
	for _, inputBit := range inputBits {
		outputBits, nextReg := ConvolutionCode(reg, inputBit)
		encodedBits = append(encodedBits, outputBits[0], outputBits[1])
		reg = nextReg
	}

	return encodedBits
}

// 最も単純な最尤系列推定
func MaximumLikelihood(received []uint8) []uint8 {
	// 受信列は2bitずつのシンボル列なので、入力全体の長さはその半分
	totalInputLength := len(received) / 2
	// 終端3bitを除いた情報ビット部分だけを総当たりする
	infoLength := totalInputLength - 3
	totalCandidates := 1 << uint(infoLength)
	// 最良候補の入力ビット列を保持
	bestInput := make([]uint8, totalInputLength)
	// 最初は十分大きい値 (受信列長 + 1) を最良距離として設定
	bestDistance := len(received) + 1
	// 候補ビット列を生成するための作業領域
	currentInput := make([]uint8, totalInputLength)

	for candidateIndex := range totalCandidates {
		// candidateIndexを情報ビット列に変換し、終端3bitは0のままにする
		for position := range infoLength {
			shift := uint(infoLength - 1 - position)
			currentInput[position] = uint8((candidateIndex >> shift) & 1)
		}

		// 候補を符号化
		currentOutput := ConvolutionEncode(currentInput)

		// 受信列とのハミング距離を計算し、最良候補を更新
		distance := HammingDistance(received, currentOutput)
		if distance < bestDistance {
			bestDistance = distance
			copy(bestInput, currentInput)
		}
	}

	// 終端ビット3bitを除いて、元の情報ビット列だけ返す
	return bestInput[:infoLength]
}

// ビタビ複合
func ViterbiDecode(receivedBits []uint8) []uint8 {
	// 受信列を2bitごとのシンボルに分割
	symbols := make([][2]uint8, len(receivedBits)/2)
	for index := range symbols {
		base := index * 2
		symbols[index] = [2]uint8{receivedBits[base], receivedBits[base+1]}
	}

	// レジスタは3bitなので、状態数は2^3 = 8
	numStates := 8
	// 受信列長 + 1を最大パスメトリック (到達不能状態) として設定
	maxPathMetric := len(receivedBits) + 1
	// 先頭はレジスタ000のみをパスメトリック0、他は到達不能として初期化
	pathMetrics := make([]int, numStates)
	for state := range pathMetrics {
		pathMetrics[state] = maxPathMetric
	}
	pathMetrics[0] = 0

	// あるステップまでで、各レジスタ状態に到達するまでの最良入力ビット列を保持
	survivorPaths := make([][]uint8, numStates)

	// 各シンボルについて、全レジスタ状態から次レジスタ状態への候補を評価s
	for _, symbol := range symbols {
		// 次ステップ用のパスメトリックを、到達不能として初期化
		newPathMetrics := make([]int, numStates)
		for state := range newPathMetrics {
			newPathMetrics[state] = maxPathMetric
		}

		// 次ステップの生存パスを保持
		newSurvivorPaths := make([][]uint8, numStates)

		for state := range numStates {
			// 現在到達できない状態は無視
			if pathMetrics[state] == maxPathMetric {
				continue
			}

			// 今のレジスタ状態から、入力bit 0/1 の2つの分岐を評価する
			for inputBit := uint8(0); inputBit <= 1; inputBit++ {
				// 畳み込み計算
				outputBits, nextReg := ConvolutionCode(uint8(state), inputBit)
				// 累積パスメトリックを計算
				currentMetric := pathMetrics[state] + HammingDistance(symbol[:], outputBits[:])

				// これまでの候補よりパスメトリックが小さければ、その状態の生存パスを更新
				if currentMetric < newPathMetrics[nextReg] {
					newPathMetrics[nextReg] = currentMetric
					newSurvivorPaths[nextReg] = append(append([]uint8(nil), survivorPaths[state]...), inputBit)
				}
			}
		}

		// パスメトリックと生存パスを更新
		pathMetrics = newPathMetrics
		survivorPaths = newSurvivorPaths
	}

	// 最終的にレジスタ000に残った経路が最良解
	bestPath := survivorPaths[0]

	// 終端ビット3bitを除いて、元の情報ビット列だけを返す
	return bestPath[:len(bestPath)-3]
}

// 文字列をビット配列に変換
func ParseBitString(bitString string) []uint8 {
	bits := make([]uint8, len(bitString))
	for index := range bitString {
		bits[index] = bitString[index] - '0'
	}

	return bits
}

// ビット配列を文字列に変換
func BitsToString(bits []uint8) string {
	result := make([]byte, len(bits))
	for index, bit := range bits {
		result[index] = byte('0' + bit)
	}

	return string(result)
}

// 例の受信ビット列
const received = "110101001100010101110110010100101101111011"

func main() {
	maximumLikelihoodDecoded := MaximumLikelihood(ParseBitString(received))
	viterbiDecoded := ViterbiDecode(ParseBitString(received))

	fmt.Println("受信ビット列: ", received)
	fmt.Println("最尤系列推定による復号結果: ", BitsToString(maximumLikelihoodDecoded))
	fmt.Println("ビタビ復号による復号結果: ", BitsToString(viterbiDecoded))
}
