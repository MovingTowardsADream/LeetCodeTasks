package main

func diagonalSum(mat [][]int) int {
	sum, N := 0, len(mat)
	for i := 0; i < N; i++ {
		sum += mat[i][i] + mat[i][N-i-1]
	}
	if N%2 == 1 {
		sum -= mat[N/2][N/2]
	}
	return sum
}
