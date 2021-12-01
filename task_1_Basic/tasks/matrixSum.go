package tasks

func MatrixElementsSum(mat [][]int) int {
	total := 0
	for j := range mat[0] {
		for i := range mat {

			if mat[i][j] == 0 {
				break
			} else {
				total += mat[i][j]
			}
		}
	}
	return total
}
