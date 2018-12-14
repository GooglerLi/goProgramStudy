/*
 最近产品需要导一批数据，需要用到两个文件就交集，在交集中显示“是”，不在交集中显示“否”，
 因为最近在学习go语言，所以就go写了一个小工具处理数据了。
 如果只是简单的求交集，可以用sort，uniq等工具处理。代码比较简单，就不解释了。如下：
*/

package slicestudy

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// SliceCalculate 计算两个切片的交集
func SliceCalculate() {

	//定义两个切片存储文本中的数据
	var doubleRoleSlice []int
	var onlyOneContractAnddoubleRoleSlice []int

	allDoubleRoleUID, err1 := os.Open("C:\\Users\\Administrator\\Desktop\\doublerole\\allsordoublerole.txt")
	if err1 != nil {
		fmt.Println("File reading error", err1)
		return
	}

	defer allDoubleRoleUID.Close()

	onlyOneContractAnddoubleRoleUID, err2 := os.Open("C:\\Users\\Administrator\\Desktop\\doublerole\\onlyOneContractAnddoubleRole.txt")
	if err2 != nil {
		fmt.Println("File reading error", err2)
		return
	}

	defer onlyOneContractAnddoubleRoleUID.Close()

	resultFile, err3 := os.Create("C:\\Users\\Administrator\\Desktop\\doublerole\\result.txt") //创建文件
	if err3 != nil {
		fmt.Println("File create error", err3)
		return
	}

	allDoubleRoleUIDRd := bufio.NewReader(allDoubleRoleUID)
	for {
		line, err := allDoubleRoleUIDRd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}

		line = strings.TrimSpace(line)

		uid, _ := strconv.Atoi(line)
		doubleRoleSlice = append(doubleRoleSlice, uid)
	}

	onlyOneContractAnddoubleRoleUIDRd := bufio.NewReader(onlyOneContractAnddoubleRoleUID)
	for {
		line, err := onlyOneContractAnddoubleRoleUIDRd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}

		line = strings.TrimSpace(line)

		uid, _ := strconv.Atoi(line)
		onlyOneContractAnddoubleRoleSlice = append(onlyOneContractAnddoubleRoleSlice, uid)
	}

	fmt.Printf("doubleRoleSlice len is %d \n", len(doubleRoleSlice))
	fmt.Printf("onlyOneContractAnddoubleRoleSlice len is %d \n", len(onlyOneContractAnddoubleRoleSlice))

	isSame := false
	for _, m := range doubleRoleSlice {
		for _, v := range onlyOneContractAnddoubleRoleSlice {
			if m == v {
				isSame = true
				break
			} else {
				isSame = false
			}
		}

		if isSame {
			resultFileWd := bufio.NewWriter(resultFile)
			_, _ = resultFileWd.WriteString("是\n")
			resultFileWd.Flush()
		} else {
			resultFileWd := bufio.NewWriter(resultFile)
			_, _ = resultFileWd.WriteString("否\n")
			resultFileWd.Flush()
		}
	}

	resultFile.Close()
}
