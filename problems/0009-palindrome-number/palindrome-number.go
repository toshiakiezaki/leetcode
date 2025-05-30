func isPalindrome(x int) bool {
    reverseX := 0
    for num := x; num > 0; {
        reverseX = reverseX * 10 + num % 10
        num = num / 10
    }
    return reverseX == x
}