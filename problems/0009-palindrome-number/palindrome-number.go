func isPalindrome(x int) bool {
    reverseX := 0
    for num := x; num > 0; {
        rev := num % 10
        reverseX = reverseX * 10 + rev
        num = num / 10
    }
    return reverseX == x
}