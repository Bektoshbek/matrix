package services

// IsCompatible is a method for checking whether matrixes are compatible for carrying out a particular operation
func (m1 *Matrix) IsCompatible(m2 Matrix, op string) bool {
	if op == "+" || op == "-" {
		if m1.Columns == m2.Columns && m1.Rows == m2.Rows {
			return true
		}
	} else {
		if m1.Columns == m2.Rows && m1.Rows == m2.Columns {
			return true
		}
	}
	return false
}

// Add is a method for adding the first matrix to the second
func (m1 *Matrix) Add(m2 Matrix) (Matrix, error) {
	if !m1.IsCompatible(m2, "+") {
		return m2, ErrNotCompatible
	}
	var (
		res Matrix
	)
	for rowInd := 0; rowInd < m1.Rows; rowInd++ {
		var nums []int
		for i := 0; i < m1.Columns; i++ {
			nums = append(nums, m1.Content[rowInd][i]+m2.Content[rowInd][i])
		}
		res.Content = append(res.Content, nums)
	}
	res.Columns = m1.Columns
	res.Rows = m1.Rows
	return res, nil
}

// Substract is a method for substracting the second matrix from the first one
func (m1 *Matrix) Substract(m2 Matrix) (Matrix, error) {
	if !m1.IsCompatible(m2, "-") {
		return m2, ErrNotCompatible
	}
	var res Matrix
	for rowInd := 0; rowInd < m1.Rows; rowInd++ {
		var nums []int
		for i := 0; i < m1.Columns; i++ {
			nums = append(nums, m1.Content[rowInd][i]*m2.Content[rowInd][i])
		}
		res.Content = append(res.Content, nums)
	}
	res.Columns = m1.Columns
	res.Rows = m1.Rows
	return res, nil
}

// Multiply is a method for multiplying matrixes
func (m1 *Matrix) Multiply(m2 Matrix) (Matrix, error) {
	if !m1.IsCompatible(m2, "*") {
		return m2, ErrNotCompatible
	}
	var res Matrix
	for rowInd := 0; rowInd < m1.Rows; rowInd++ {
		row := []int{}
		for m2ColInd := 0; m2ColInd < m2.Columns; m2ColInd++ {
			num := 0
			for m1ColInd := 0; m1ColInd < m1.Columns; m1ColInd++ {
				num += m1.Content[rowInd][m1ColInd] * m2.Content[m1ColInd][m2ColInd]
			}
			row = append(row, num)
		}
		res.Content = append(res.Content, row)
	}
	return res, nil
}
