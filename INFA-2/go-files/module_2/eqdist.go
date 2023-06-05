package main

import (
    "fmt";
    )

type doubleQueue struct {
	data []int; // здесь будем хранить номера вершин
	info []int; // а здесь расстояния, которые им нужно поставить
}
func (q *doubleQueue) init() {
	q.data = make([]int, 0);
	q.info = make([]int, 0);
}
func (q *doubleQueue) push(x, y int) {
	q.data = append(q.data, x);
	q.info = append(q.info, y);
}
func (q *doubleQueue) pop() (int, int) {
	x := q.data[0];
	q.data = q.data[1:];
	y := q.info[0];
	q.info = q.info[1:];
	return x, y;
}
func (q *doubleQueue) isEmpty() bool {
	return len(q.data) == 0;
}


type edge struct {
	number int;
	edge_list *edge;
}

type vertex struct {
	mark bool;
	edge_list *edge;
	dist []int; // массив расстояний до каждой из опорных вершин
}

func add_edge_to_one(x, y int, vertexes *[]vertex) {
	new_edge := new(edge);
	new_edge.number = y;
	old_edge := (*vertexes)[x].edge_list;
	(*vertexes)[x].edge_list = new_edge;
	(*vertexes)[x].edge_list.edge_list = old_edge;
}

func add_edge(x, y int, vertexes *[]vertex) {
	add_edge_to_one(x, y, vertexes);
	add_edge_to_one(y, x, vertexes);
}

func min(x, y int) int {
	if (x < y) { return x;
	} else { return y; }
}



func main() {
	var n, m, x, y, k int;
	fmt.Scanf("%d", &n); // vertexes
	fmt.Scanf("%d", &m); // edges

	vertexes := make([]vertex, n);

	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d", &x, &y);
		add_edge(x, y, &vertexes);
	}

	fmt.Scanf("%d", &k); // amount of reference vertexes
	refVertexes := make([]int, k); // numbers of ref vertexes
	for i := 0; i < k; i++ {
		fmt.Scanf("%d", &x);
		refVertexes[i] = x;
	}

	// BFS
	for i, _ := range vertexes {
		vertexes[i].dist = make([]int, k);
	}
	var q doubleQueue;
	q.init();

	for i, refVert := range refVertexes {
		// i = 0..k
		// refVert = 0..n-1
		for j, _ := range vertexes {
			vertexes[j].mark = false;
		}
		vertexes[refVert].mark = true;
		q.push(refVert, 0);
		for !q.isEmpty() {
			vertNum, vertDist := q.pop();
			vert := vertexes[vertNum];
			vert.dist[i] = vertDist;
			edge_list := vert.edge_list;
			for ;edge_list != nil; edge_list = edge_list.edge_list {
				if (!(vertexes[edge_list.number].mark)) {
					vertexes[edge_list.number].mark = true;
					q.push(edge_list.number, vertDist + 1);
				}

			}
		}
	}
	result := make([]int, 0);
	for i, _ := range vertexes {
		dist0 := vertexes[i].dist[0];
		flag := (dist0 != 0);
		for j := 1; (j < k) && flag; j++ {
			flag = ((vertexes[i].dist[j] == dist0) && flag);
		}
		if (flag) {
			result = append(result, i);
		}
	}

	if (len(result) == 0) {
		fmt.Printf("-");
	} else {
		for _, v := range result {
			fmt.Printf("%d ", v);
		}
	}
}
