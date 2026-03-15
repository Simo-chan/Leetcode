func replaceElements(arr []int) []int {
	if len(arr) < 2 {
		arr[0] = -1
		return arr
	}

	biggest := -1

	for i := len(arr) - 1; i >= 0; i-- {
		curVal:= arr[i]
        arr[i] = biggest

        if curVal > biggest {
            biggest = curVal
        }
	}

	return arr
}

//Time complexity is O(n)
//Space complaxity is O(1)