package main

import (
	"fmt"
)

func countNumberOfDecodings(inputStr string) int {
	if len(inputStr) == 0 || inputStr[0] == '0' {
		return 0
	}

	characters := len(inputStr)

	dp := make([]int, characters+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= characters; i++ {
		if inputStr[i-1] != '0' {
			dp[i] += dp[i-1]
		}

		if twoDigit := (inputStr[i-2]-'0')*10 + (inputStr[i-1] - '0'); twoDigit >= 10 && twoDigit <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[characters]
}

func main() {
	testCases := []struct {
		input    string
		expected int
	}{
		{"12", 2},    // "AB", "L"
		{"226", 3},   // "BZ", "VF", "BBF"
		{"0", 0},     // No valid decodings
		{"06", 0},    // No valid decodings
		{"11106", 2}, // "AAJF", "KJF"
		{"1234", 3},  // "ABCD", "LCD", "AWD"
	}

	for _, tc := range testCases {
		result := countNumberOfDecodings(tc.input)
		fmt.Printf("Input: %s, Expected: %d, Got: %d\n", tc.input, tc.expected, result)
	}
}
