package main

import (
    "fmt";
    "sort";
    "math";
    )

// disjoint set union
func make_set(v int, parent, rank *[]int) {
	(*parent)[v] = v;
	(*rank)[v] = 0;
}

func find_set(v int, parent, rank *[]int) int {
	if (v == (*parent)[v]) { return v; }
	(*parent)[v] = find_set((*parent)[v], parent, rank);
	return (*parent)[v];
}

func union_sets(a, b int, parent, rank *[]int) {
	a = find_set(a, parent, rank);
	b = find_set(b, parent, rank);
	if (a != b) {
		if ((*rank)[a] < (*rank)[b]) { a, b = b, a; }
		(*parent)[b] = a;
		if ((*rank)[a] == (*rank)[b]) { (*rank)[a]++; }
	}
}
// --------------------
// edge and edges types for sorting and storage
type edge struct {
	// number1 < number2
	number1, number2, length_square int;
}

type edgesArray []edge;

func (edges edgesArray) Len() int           { return len(edges); }
func (edges edgesArray) Swap(i, j int)      { edges[i], edges[j] = edges[j], edges[i]; }
func (edges edgesArray) Less(i, j int) bool { return edges[i].length_square < edges[j].length_square; }

type vertex struct {
	x, y int; // координаты
}

func find_dist_square(a, b int, vertexes *[]vertex) int {
	// a, b - vertex numbers
	a_x := (*vertexes)[a].x;
	a_y := (*vertexes)[a].y;
	b_x := (*vertexes)[b].x;
	b_y := (*vertexes)[b].y;
	return (a_x - b_x) * (a_x - b_x) + (a_y - b_y) * (a_y - b_y);
}


func main() {
	var n, x, y int;
	fmt.Scanf("%d", &n); // vertex amount

	vertexes := make([]vertex, n);
	edges := make([]edge, (n * (n - 1)) / 2);

	// for DSU
	parent := make([]int, n);
	rank := make([]int, n);

	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d", &x, &y);
		vertexes[i].x, vertexes[i].y = x, y;
		make_set(i, &parent, &rank);
	}

	x = 0;
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			edges[x].number1, edges[x].number2 = i, j;
			edges[x].length_square = find_dist_square(i, j, &vertexes);
			x++;
		}
	}

	sort.Sort(edgesArray(edges));

	// spanning tree algorythm
	new_edges := make([]edge, n - 1);
	len_new_edges := 0; // amount of new edges, bc len(new_edges) = n - 1 (const)
	for i := 0; len_new_edges < n - 1; i++ {
		possible_edge := edges[i];
		vert1, vert2 := possible_edge.number1, possible_edge.number2;
		if (find_set(vert1, &parent, &rank) != find_set(vert2, &parent, &rank)) {
			new_edges[len_new_edges] = possible_edge;
			len_new_edges++;
			union_sets(vert1, vert2, &parent, &rank);
		}
	}

	sum := 0.0;
	for i := range new_edges {
		sum += math.Sqrt(float64(new_edges[i].length_square));
	}
	fmt.Printf("%.2f\n", sum);

}
