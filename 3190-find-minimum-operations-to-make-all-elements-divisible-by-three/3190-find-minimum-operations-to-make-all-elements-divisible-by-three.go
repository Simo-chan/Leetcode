//If a number is not divisible by 3, it always requires only one operation to make it divisible by 3.
//So in this case the only thing we have to do is to count numbers that are not divisible by 3. 

func minimumOperations(nums []int) int {
    var count int 

    for _, num := range nums {
        if num % 3 != 0 {
            count++
        }
    }

    return count    
}

//Time complexity is O(n)
//Space complexity is O(1)