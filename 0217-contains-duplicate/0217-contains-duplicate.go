func containsDuplicate(nums []int) bool {
    uniqueNums := make(map[int]struct{})

    for _, num:= range nums{
        if _, exists := uniqueNums[num]; !exists {
            uniqueNums[num] = struct{}{}
        } else {
            return true
        }
    }

    return false 
}