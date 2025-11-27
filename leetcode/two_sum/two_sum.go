package two_sum

/// things need to be done
/*
1. take a array
2. choose the traget sum number
3. pick the first number, pair with every other number in array
4. pick the second number, pair with every other number in array
5. continue until you find a pair that equals the target

*/
func TwoSumBruteForce(arr []int, target int) []int {
	// loop through each number
	for i := 0; i < len(arr); i++ {
		// for each number , loop through all numbers after it
		for j := i + 1; j < len(arr); j++ {
			// check if two numbers add to the target or not
			if arr[i]+arr[j] == target {

				return []int{i, j}
			}
		}
	}
	return nil
}

func TwoSumHash(arr []int, target int) []int {
	alreadySeen := make(map[int]int)

	for index, value := range arr {
		isComplement := target - value

		if idx, ok := alreadySeen[isComplement]; ok {
			return []int{idx, index}
		}

		alreadySeen[value] = index
	}
	return []int{}
}
