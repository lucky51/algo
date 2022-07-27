package leetcode399

// calcEquation 除法求值 ，TODO: 官方题解并查集 ，没看懂，暂时翻译了个 go版本
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	equationSize := len(equations)
	unionFind := NewUnionFind(2 * equationSize)
	idMap := make(map[string]int, 2*equationSize)
	id := 0
	for i := 0; i < equationSize; i++ {
		equation := equations[i]
		var1, var2 := equation[0], equation[1]
		if _, ok := idMap[var1]; !ok {
			idMap[var1] = id
			id++
		}
		if _, ok := idMap[var2]; !ok {
			idMap[var2] = id
			id++
		}
		unionFind.Union(idMap[var1], idMap[var2], values[i])
	}
	// 第2步 做查询
	queriesSize := len(queries)
	res := make([]float64, queriesSize)
	for i := 0; i < queriesSize; i++ {
		var1, var2 := queries[i][0], queries[i][1]
		id1, ok1 := idMap[var1]
		id2, ok2 := idMap[var2]
		if !ok1 || !ok2 {
			res[i] = -1.0
		} else {
			res[i] = unionFind.IsConnected(id1, id2)
		}
	}
	return res
}

type UnionFind struct {
	parent []int
	// 指向父节点的权值
	weight []float64
}

func NewUnionFind(n int) UnionFind {
	uf := UnionFind{
		parent: make([]int, n),
		weight: make([]float64, n),
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.weight[i] = 1.0
	}
	return uf
}

func (this *UnionFind) Union(x, y int, value float64) {
	rootX, rootY := this.Find(x), this.Find(y)
	if rootX == rootY {
		return
	}
	this.parent[rootX] = rootY
	this.weight[rootX] = this.weight[y] * value / this.weight[x]
}

// Find 路径压缩
func (this *UnionFind) Find(x int) int {
	if x != this.parent[x] {
		origin := this.parent[x]
		this.parent[x] = this.Find(this.parent[x])
		this.weight[x] *= this.weight[origin]
	}
	return this.parent[x]
}

// IsConnected
func (this *UnionFind) IsConnected(x, y int) float64 {
	rootX, rootY := this.Find(x), this.Find(y)
	if rootX == rootY {
		return this.weight[x] / this.weight[y]
	} else {
		return -1
	}
}
