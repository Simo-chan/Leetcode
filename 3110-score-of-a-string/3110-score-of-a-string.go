func scoreOfString(s string) int {
    var result int

    //We iterate only until the second to last elenemt to avoid going out of bounds.
    //We convert byte to int right away because byte can hold values only from 0 to 255.
    // Then we need to find the absolute value, because sometimes substruction results in a 
    //negative number. If it's the case we substruct negative number from the result variable, 
    //because double minus equals to plus.
    for i:= 0; i < len(s)-1; i++ {
        difference:= int(s[i]) - int(s[i+1])

        if difference > 0 {
            result+= difference
        } else {
            result-= difference
        }
    }

    return result
}

//Time complexity is O(n) - we iterate through the string only once.
//Space complexity is O(1) - we don't create any data scructure.
