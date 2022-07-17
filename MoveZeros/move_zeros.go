package movezeros

func MoveZeros(arr []int) []int {
	arrLen := len(arr)
	// Creating another array to insert at the correct order
	nonZeroArr := make([]int, arrLen)

	nonZeroIndex := 0
	lastZeroIndex := arrLen - 1

	// Iterating through the arr and populate the new array with correct positions
	for i := 0; i < len(arr); i++ {
		switch arr[i] {
		case 0:
			nonZeroArr[lastZeroIndex] = arr[i]
			lastZeroIndex -= 1
		default:
			nonZeroArr[nonZeroIndex] = arr[i]
			nonZeroIndex += 1
		}
	}

	return nonZeroArr
}
