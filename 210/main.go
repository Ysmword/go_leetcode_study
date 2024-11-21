package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		edges   = make([][]int, numCourses) // 表示学习了A之后能学习的课程
		visited = make([]int, numCourses)   // 表示这个课程是完成可以学习的
		result  = make([]int, 0)
		valid   = true
		dfs     func(u int)
	)

	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
		result = append([]int{u}, result...) // 头插法, 表示先学习u
	}

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	if !valid {
		return []int{}
	}
	return result
}

func main() {
	fmt.Println(findOrder(2, [][]int{{1, 0}}))
}
