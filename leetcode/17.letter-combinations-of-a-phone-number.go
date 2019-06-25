/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */
func letterCombinations(digits string) []string {
	var res []string
	mapping := map[string]string{
		"0": "0",
		"1": "1",
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz"}
	local := ""
	if len(digits) == 0 {
		return res
	}
	res = backtracking(res, &digits, 0, mapping, &local)
	return res
}

func backtracking(
	res []string,
	digits *string,
	level int,
	mapping map[string]string,
	local *string) []string {
	if level == len(*digits) {
		res = append(res, *local)
	} else {
		dgt := string((*digits)[level])
		characters := mapping[dgt]
		for _, char := range characters {
			*local += string(char)
			res = backtracking(res, digits, level+1, mapping, local)
			*local = (*local)[:len(*local)-1]
		}
	}
	return res
}

