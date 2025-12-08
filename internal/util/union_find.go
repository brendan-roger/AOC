package util

type UnionFind struct {
	parent []int
	size   []int
}

func InitUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}

	return &UnionFind{parent: parent, size: size}
}

func (uf *UnionFind) Find(x int) int {

	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}

	return uf.parent[x]
}

func (uf *UnionFind) Union(a, b int) bool {
	ra := uf.Find(a)
	rb := uf.Find(b)

	if ra == rb {
		return false
	}

	if uf.size[ra] < uf.size[rb] {
		ra, rb = rb, ra
	}

	uf.parent[rb] = ra
	uf.size[ra] += uf.size[rb]

	return true
}
