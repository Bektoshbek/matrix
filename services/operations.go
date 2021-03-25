package services

func (m1 *Matrix) Add(m2 Matrix) Matrix {
	var (
		res Matrix
	)
	for rowInd := 0; rowInd < m1.Rows; rowInd++ {
		var nums []int
		for i := 0; i < m1.Columns; i++ {
			nums = append(nums, m1.Content[rowInd][i] + m2.Content[rowInd][i])
		}
		res.Content = append(res.Content, nums)
	}
	res.Columns = m1.Columns
	res.Rows = m1.Rows
	return res
}
