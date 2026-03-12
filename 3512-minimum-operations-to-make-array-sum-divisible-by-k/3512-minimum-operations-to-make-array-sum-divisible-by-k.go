func minOperations(nums []int, k int) int {
    var sum int

    for _, num := range nums {
        sum += num
    }

    //Modulo operator will return the remainder which will be exactly the amount of operations needed.
    return sum % k
}

//Time complexity is O(n)
//Space complexity is O(1)
