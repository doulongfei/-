package main

import (
	"fmt"
)

type ValNode struct {
	row int
	col int
	val int
}

func main() {

	//	创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2
	//	输出原始数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
	//	转成稀疏数组
	var sparseArr []ValNode
	valNode := ValNode{11, 11, 0}
	sparseArr = append(sparseArr, valNode)
	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				//	创建一个节点 valNode
				valNode := ValNode{i, j, v2}
				sparseArr = append(sparseArr, valNode)

			}
		}
	}
	//遍历系数矩阵
	for i, v := range sparseArr {
		fmt.Printf("%d:%d %d %d\n", i, v.row, v.col, v.val)
	}

}
