package main

type UnionFind struct {
    parent []int
    rank []int
}

// UnionFindInit function intializes the varibales
func (u *UnionFind) UnionFindInit(n int) {
    for i := 0; i < n + 1; i++ {
        u.parent = append(u.parent, i)
        u.rank = append(u.rank, 1)
    }
}

// Find function finds parent of an index
func (u *UnionFind) Find(x int) int {
    if u.parent[x] != x {
        u.parent[x] = u.Find(u.parent[x])
    }

    return u.parent[x]
}

// Union function connects two cells
func (u *UnionFind) Union(x int, y int) bool {
    p1 := u.Find(x)
    p2 := u.Find(y)

    if p1 == p2 {
        return false
    } else if u.rank[p1] > u.rank[p2] {
        u.parent[p2] = p1
        u.rank[p1] += u.rank[p2]
    } else {
        u.parent[p1] = p2
        u.rank[p1] += u.rank[p1]
    }

    return true
}