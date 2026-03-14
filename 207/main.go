package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges   = make([][]int, numCourses) // 表示学习了A之后能学习的课程
		visited = make([]int, numCourses)   // 表示这个课程是完成可以学习的
		result  = make([]int, 0)
		valid   = true
		dfs     func(u int)
	)

	dfs = func(u int) {
		visited[u] = 1 // 表示搜索中
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 { // 搜索中，成环了
				valid = false
				return
			}
		}
		visited[u] = 2 // 表示这个课程完成能够学习
		result = append(result, u)
	}

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i) // 看下这个课程是否能够被学习
		}
	}
	return valid
}

func canFinish1(numCourses int, prerequisites [][]int) bool {
	can := make(map[int][]int) // 学习完key之后，能够学习的课程
	need := make(map[int]int)  // 课程需要依赖多少课程学习完之后，才能够学习

	for i := 0; i < len(prerequisites); i++ {
		can[prerequisites[i][1]] = append(can[prerequisites[i][1]], prerequisites[i][0])
		need[prerequisites[i][0]]++
	}

	q := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if need[i] == 0 {
			q = append(q, i)
		}
	}

	courses := make([]int, 0)
	for len(q) > 0 {
		course := q[0]
		courses = append(courses, course)
		q = q[1:]
		for _, c := range can[course] {
			need[c]--
			if need[c] == 0 {
				q = append(q, c)
			}
		}
	}
	// 如果学习不完，肯定有环的情况
	return len(courses) == numCourses
}

func main() {
	fmt.Println(canFinish1(2, [][]int{{1, 0}}))

	fmt.Println(canFinish1(2, [][]int{{1, 0}, {0, 1}}))
}
