package binary_search

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */
func guess(num int) int {
	return num
}
func guessNumber(n int) int {
	L, R := 1, n
	for L < R {
		mid := L + (R-L)>>1
		if guess(mid) <= 0 {
			R = mid
		} else {
			L = mid + 1
		}
	}
	return L
}
