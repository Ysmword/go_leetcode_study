package main

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
