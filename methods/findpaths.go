package methods

func possiblePath(uniqPath [][]int, result []LemIn) [][]int {
	var res [][]int
	i := 0
	for i < len(uniqPath) {
		if len(uniqPath[i]) > 1 {
			if isUniq(uniqPath[i], result) {
				res = append(res, uniqPath[i])
			}
		} else if i == 0 {
			res = append(res, uniqPath[i])
		}
		i++
	}
	return res
}
func isUniq(list []int, lemin []LemIn) bool {
	var comp []string
	i := 0
	for i < len(list) {
		k := 1
		for k < len(lemin[list[i]].Path)-1 {
			comp = append(comp, lemin[list[i]].Path[k].Name)
			k++
		}
		i++
	}
	i = 0
	for i < len(comp) {
		k := i + 1
		for k < len(comp) {
			if comp[i] == comp[k] {
				return false
			}
			k++
		}
		i++
	}
	return true
}
