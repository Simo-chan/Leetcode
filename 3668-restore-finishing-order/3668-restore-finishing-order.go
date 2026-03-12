func recoverOrder(order []int, friends []int) []int {
    var result []int

    for i:= 0; i < len(order); i++ {
        for j:= 0; j < len(friends); j++ {
            if order[i] == friends[j] {
                result = append(result, order[i])
            }
        }
    }

    return result
}

//Time complexity is O(n^2)
//Space complexity O(1)