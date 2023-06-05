package main

import (
    "fmt";
    )

type edge struct {
	number int;
	mark bool;
	edge_list *edge;
}

type vertex struct {
	mark bool;
	edge_list *edge;
	tin, tout int;
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


func dfs (v, p, timer, count int, vertexes *[]vertex) int {
	(*vertexes)[v].mark = true;
	timer++;
	(*vertexes)[v].tin = timer;
	(*vertexes)[v].tout = timer;
	edge_list := (*vertexes)[v].edge_list;
	for ;edge_list != nil; edge_list = edge_list.edge_list {
		to := edge_list.number;
		if (to == p) { continue; }
		if ((*vertexes)[to].mark) {
			(*vertexes)[v].tout = min((*vertexes)[v].tout, (*vertexes)[to].tin);
		} else {
			count = dfs(to, v, timer, count, vertexes);
			(*vertexes)[v].tout = min((*vertexes)[v].tout, (*vertexes)[to].tout);
			if ((*vertexes)[to].tout > (*vertexes)[v].tin) {
				// bridge v--to
				count++;
			}
		}
	}
	return count;
}


func main() {
	var n, m, x, y int;
	fmt.Scanf("%d", &n); // vertexes
	fmt.Scanf("%d", &m); // edges

	vertexes := make([]vertex, n);

	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d", &x, &y);
		add_edge(x, y, &vertexes);
	}

	for i, _ := range vertexes {
		vertexes[i].mark = false;
	}

	count := 0;
	for i, _ := range vertexes {
		if (!vertexes[i].mark) {
			count += dfs(i, -1, 0, 0, &vertexes);
		}
	}
	fmt.Printf("%d\n", count);
}
