func isPalindrome(x int) bool {
    str := strconv.Itoa(x)
    for i := 0; i < len(str) / 2; i++ {
        if(str[i] != str[len(str) - i - 1]) {
            return false
        }
    }
    return true
}