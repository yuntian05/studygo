package main

import "fmt"

func main() {
	//fmt.Println(openLock([]string{"0201","0101","0102","1212","2002"}, "0202"))
	fmt.Println(openLock([]string{"0000"}, "8888"))
}

func openLock(deadends []string, target string) int {
	if len(target) == 0 {
		return -1
	}
	deads := make(map[string]bool)
	visited := make(map[string]bool)
	for _,val := range deadends {
		deads[val] = true
	}

	// 从起点开始广度搜索
	step := 0
	queue := []string{"0000"}

	for len(queue) > 0 {
		size := len(queue)
		// 将当前队列中的所有节点向四周扩散
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			// 判断节点是否到达终点
			if _,ok := deads[cur]; ok {
				continue
			}
			//fmt.Println(cur)
			if cur == target {
				fmt.Println("end=>",cur)
				return step
			}

			// 将一个节点的相邻节点加入队列
			for j := 0; j < 4; j++ {
				up := plusOne(cur, j)
				if _,ok := visited[up]; !ok {
					queue = append(queue, up)
					visited[up] = true
				}
				down := minusOne(cur, j)
				if _,ok := visited[down]; !ok {
					queue = append(queue, down)
					visited[down] = true
				}
			}
		}
		step++
	}
	// 如果穷举完都没找到目标密码，那就是找不到了
	return -1
}

// 将s[j]向上拨动一下
func plusOne(s string, j int) string {
	sByte := []byte(s)
	if sByte[j] == '9' {
		sByte[j] = '0'
	} else {
		sByte[j]++
	}
	return string(sByte)
}

// 将s[j]向下拨动一下
func minusOne(s string, j int) string {
	sByte := []byte(s)
	if sByte[j] == '0' {
		sByte[j] = '9'
	} else {
		sByte[j]--
	}
	return string(sByte)
}

