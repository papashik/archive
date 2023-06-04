package main

import (
    "fmt";
    //"bufio";
    //"os";
    //"strings";
    )

type queue struct {
	data []int;
}
func (q *queue) init() {
	q.data = make([]int, 0);
}
func (q *queue) push(x int) {
	q.data = append(q.data, x);
}
func (q *queue) pop() int {
	x := q.data[0];
	q.data = q.data[1:];
	return x;
}
func (q *queue) isEmpty() bool {
	return len(q.data) == 0;
}



type edge struct {
	number int;
	color string;
	mark bool;
	edge_list *edge;
}

type vertex struct {
	color string;
	mark bool;
	edge_list *edge;
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


func main() {
	var n, m, x, y int;
	fmt.Scanf("%d", &n); // vertexes
	fmt.Scanf("%d", &m); // edges

	vertexes := make([]vertex, n);


	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d", &x, &y);
		add_edge(x, y, &vertexes);
	}

	// BFS
	for i, _ := range vertexes {
		vertexes[i].mark = false;
	}
	var q queue;
	q.init();
	var max_vertex_num, max_edge_num, min_start_vertex int; // max component data
	var vertex_num, edge_num, start_vertex int; // current component data
	for i, _ := range vertexes {
		if (!vertexes[i].mark) {
			// new component
			vertex_num = 0;
			edge_num = 0;
			start_vertex = i;

			vertexes[i].mark = true;
			q.push(i);
			for !q.isEmpty() {
				vert := vertexes[q.pop()];
				vertex_num++;
				edge_list := vert.edge_list;
				for edge_list != nil {
					if (!edge_list.mark) {
						edge_list.mark = true;
						edge_num++;
					}
					if (!(vertexes[edge_list.number].mark)) {
						vertexes[edge_list.number].mark = true;
						q.push(edge_list.number);
					}
					edge_list = edge_list.edge_list;
				}
			}
			// fmt.Printf("%d %d %d\n", start_vertex, vertex_num, edge_num);
			if (vertex_num > max_vertex_num) ||
			   (vertex_num == max_vertex_num && edge_num > max_edge_num) ||
			   (vertex_num == max_vertex_num && edge_num == max_edge_num && start_vertex < min_start_vertex) {
			    max_vertex_num = vertex_num;
			    max_edge_num = edge_num;
			    min_start_vertex = start_vertex;
			}
		}
	}
	// fmt.Printf("--- %d %d %d\n", min_start_vertex, max_vertex_num, max_edge_num);

	for i, _ := range vertexes {
		vertexes[i].mark = false;
	}
	q.push(min_start_vertex);
	for !q.isEmpty() {
		vert := &vertexes[q.pop()];
		vert.color = "red";
		edge_list := vert.edge_list;
		for edge_list != nil {
			edge_list.color = "red";
			if (!(vertexes[edge_list.number].mark)) {
				vertexes[edge_list.number].mark = true;
				q.push(edge_list.number);
			}
			edge_list = edge_list.edge_list;
		}
	}

	fmt.Printf("graph {\n");
	for i, _ := range vertexes {
		color := vertexes[i].color;
		if (color == "") { fmt.Printf("%d\n", i);
		} else { fmt.Printf("%d [color=%s]\n", i, color); }
	}
	for i := 0; i < n; i++ {
		edge_list := vertexes[i].edge_list;
		for edge_list != nil {
			if (edge_list.number > i) {
				color := edge_list.color;
				if (color == "") { fmt.Printf("%d--%d\n", i, edge_list.number);
				} else { fmt.Printf("%d--%d [color=%s]\n", i, edge_list.number, color); }
			}
			edge_list = edge_list.edge_list;
		}
	}
	fmt.Printf("}");
}
