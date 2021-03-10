package main

/**
* Author: zhi-xiang.meng
* Email: flycomte@qq.com
* Tel: 13372681480
* Date: 03/03/2021
**/

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// PossVal possible val in cell
type PossVal struct {
	i    int
	j    int
	vals []int
}

// RunSuduku run algorithm of suduku
func RunSuduku(data *[9][9]int) bool {
	curData := [9][9]int{}
	curData = *data

	minRow := 0
	minCol := 0

	preCantUpdateCount := 0
	for {
		// ShowSuduKu(data)
		count := 0
		cantUpdateCount := 0
		possibleCount := 9
		// var possibleVals = [9]int{}
		for i, row := range curData {
			for j, col := range row {
				if col == 0 {
					// if i == 6 && j == 8 {
					// 	println("test")
					// }
					possVals := GetPossibleValues(i, j, &curData)

					tmpCount := len(possVals.vals)

					if tmpCount == 0 {
						// fmt.Printf("[%d, %d] no value can fit in this cell.\n", i, j)
						return false
					}

					if tmpCount < possibleCount {
						possibleCount = tmpCount
						minRow = i
						minCol = j
					}

					if possibleCount == 1 {
						curData[i][j] = possVals.vals[0]
						possibleCount = 9
					} else {
						cantUpdateCount = cantUpdateCount + 1
					}

				} else {
					count = count + 1
				}
			}			
		}

		if preCantUpdateCount == cantUpdateCount {
			// fmt.Printf("Try to set value of cell[%d,%d]\n", minRow, minCol)
			possibleVals := GetPossibleValues(minRow, minCol, &curData)
			// fmt.Printf("cell[%d,%d] = %v\n", minRow, minCol, possibleVals)

			for _, v := range possibleVals.vals {
				curData[minRow][minCol] = v
				// fmt.Printf("set cell[%d,%d] = %d\n", minRow, minCol, v)
				if RunSuduku(&curData) == false {
					// fmt.Printf("restore cell[%d,%d] = %d\n", minRow, minCol, v)
					curData[minRow][minCol] = 0
				} else {
					return true
				}
			}

			return false

		}

		preCantUpdateCount = cantUpdateCount

		// fmt.Printf("##############[%d]##############\n", cantUpdateCount)
		if count == 81 || cantUpdateCount == 0 {
			ShowSuduKu(&curData)
			return true
		}

	}
	// ShowSuduKu(data)
}

// ShowSuduKu show it
func ShowSuduKu(data *[9][9]int) {
	for i, row := range data {
		for j, col := range row {
			print(col, ` `)
			if j%3 == 2 && j != 8 {
				print(`|`)
			}
		}

		if i%3 == 2 && i != 8 {
			println()
			println(`-------------------`)
		} else {
			println()
		}
	}
}

// GetPossibleValues get possible values
func GetPossibleValues(rowNum int, colNum int, data *[9][9]int) PossVal {
	p := PossVal{
		i: rowNum,
		j: colNum,
	}
	possibleValues := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	iStart := (rowNum / 3) * 3
	jStart := (colNum / 3) * 3
	// excluding numbers in 3x3 range data that start at data[iStart][jStart]
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// println(iStart+i, jStart+j, data[iStart+i][jStart+j])
			for k := 0; k < 9; k++ {
				if data[iStart+i][jStart+j] == possibleValues[k] {
					possibleValues[k] = 0
				}
			}
		}
	}
	// fmt.Println(possibleValues)
	// excluding number in current row
	for i := 0; i < 9; i++ {
		if i == colNum {
			continue
		}
		for k := 0; k < 9; k++ {
			if data[rowNum][i] == possibleValues[k] {
				possibleValues[k] = 0
			}
		}
	}
	// fmt.Println(possibleValues)
	// excluding number in current column
	for i := 0; i < 9; i++ {
		if i == rowNum {
			continue
		}
		for k := 0; k < 9; k++ {
			if data[i][colNum] == possibleValues[k] {
				possibleValues[k] = 0
			}
		}
	}
	// fmt.Println(rowNum, colNum, possibleValues)
	for k := 0; k < 9; k++ {
		if possibleValues[k] != 0 {
			p.vals = append(p.vals, possibleValues[k])
		}
	}
	return p
}

func main() {
	// Load sudoku.txt
	bytesRead, _ := ioutil.ReadFile("sudoku.txt")
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")
	lines = append(lines, "end")
	suduku := [9][9]int{}
	for rowNum, line := range lines {
		if rowNum%10 == 0 {
			if rowNum == 0 {
				continue
			}
			// println(line)
			fmt.Printf("\n[# solution grid %d #]\n", rowNum/10)
			// run algorithm on each sudoku
			RunSuduku(&suduku)
		} else {
			strs := strings.Split(line, "")
			for i := 0; i < 9; i++ {
				suduku[rowNum%10-1][i], _ = strconv.Atoi(strs[i])
			}
		}
	}

}
