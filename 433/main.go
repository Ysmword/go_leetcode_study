package main

func minMutation(startGene string, endGene string, bank []string) int {
	bankMap := make(map[string]bool)
	for _, v := range bank {
		bankMap[v] = true
	}
	if !bankMap[endGene] {
		return -1
	}
	queue := []string{startGene}
	count := 0
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			gene := queue[0]
			queue = queue[1:]
			for j := 0; j < len(gene); j++ {
				for _, v := range []byte("ACGT") {
					newGene := gene[:j] + string(v) + gene[j+1:]
					if newGene == endGene {
						return count + 1
					}
					if bankMap[newGene] {
						queue = append(queue, newGene)
						delete(bankMap, newGene)
					}
				}
			}
		}
		count++
	}
	return -1
}
