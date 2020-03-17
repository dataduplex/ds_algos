package main

func main() {

}

func longestPalindrome(s string) string {

	if len(s) == 0 {
		return ""
	}

	if len(s) == 1 {
		return s
	}
	result := ""
	matrix := make([][]bool, len(s))
	for i := range matrix {
		matrix[i] = make([]bool, len(s))
	}
	for size := 1; size <= len(s); size++ {
		for i := 0; i+size-1 < len(s); i++ {

			if size == 1 {
				matrix[i][i] = true
				continue
			}

			if size == 2 && string(s[i]) == string(s[i+1]) {
				matrix[i][i+1] = true
				result = s[i : i+size]
				continue
			}
			begin := i
			end := i + size - 1
			if string(s[begin]) == string(s[end]) && matrix[begin+1][end-1] == true {
				matrix[begin][end] = true
				result = s[i : i+size]
			}

		}
	}

	if result == "" {
		return s[0:1]
	}

	return result
}
