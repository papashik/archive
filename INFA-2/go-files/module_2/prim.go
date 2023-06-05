package main

import (
    "fmt";
    )

type queue struct {
	data []*vertex;
}
func (q *queue) init() {
	q.data = make([]*vertex, 0);
}
func (q *queue) push(vert *vertex, ind int) {
	vert.index = ind;
	q.data = append(q.data, vert);
}
func (q *queue) pop(vertexes *[]vertex) int {
	// find min
	min, min_ind := q.data[0], 0;
	for i, v := range q.data {
		if (v.key < min.key) {
			min = q.data[i];
			min_ind = i;
		}
	}
	q.data = append(q.data[:min_ind], q.data[min_ind + 1:]...);
	found_v := 0;
	for j := range (*vertexes) {
		if (&((*vertexes)[j]) == min) {
			found_v = j;
		}
	}
	return found_v;
}
func (q *queue) changeKey(index, new_key int) {
	for i, v := range q.data {
		if (v.index == index) {
			q.data[i].key = new_key;
			return;
		}
	}
}
func (q *queue) isEmpty() bool {
	return len(q.data) == 0;
}



type edge struct {
	number, length int;
	edge_list *edge;
}

type vertex struct {
	key, index, value int;
	edge_list *edge;
}

func add_edge_to_one(x, y, length int, vertexes *[]vertex) {
	new_edge := new(edge);
	new_edge.number = y;
	new_edge.length = length;
	old_edge := (*vertexes)[x].edge_list;
	(*vertexes)[x].edge_list = new_edge;
	(*vertexes)[x].edge_list.edge_list = old_edge;
}

func add_edge(x, y, length int, vertexes *[]vertex) {
	add_edge_to_one(x, y, length, vertexes);
	add_edge_to_one(y, x, length, vertexes);
}


func main() {
	var n, m, x, y, length int;
	fmt.Scanf("%d", &n); // vertexes
	fmt.Scanf("%d", &m); // edges

	vertexes := make([]vertex, n);
	new_vertexes := make([]vertex, n);

	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d %d", &x, &y, &length);
		add_edge(x, y, length, &vertexes);
	}

	// PRIM algorythm
	for i := range vertexes {
		vertexes[i].index = -1;
	}
	var q queue;
	q.init();
	v := 0;
	ind := 1;
	for {
		vertexes[v].index = -2;
		for edge_list := vertexes[v].edge_list; edge_list != nil; edge_list = edge_list.edge_list {
			u := edge_list.number;
			if (vertexes[u].index == -1) {
				vertexes[u].key = edge_list.length;
				vertexes[u].value = v;
				q.push(&vertexes[u], ind);
				ind++;
			} else if (vertexes[u].index != -2 && edge_list.length < vertexes[u].key) {
				vertexes[u].value = v;
				q.changeKey(vertexes[u].index, edge_list.length);
			}
		}
		if q.isEmpty() { break; }
		v = q.pop(&vertexes);
		add_edge(v, vertexes[v].value, vertexes[v].key, &new_vertexes);
	}

	sum := 0;
	for i := 0; i < n; i++ {
		edge_list := new_vertexes[i].edge_list;
		for edge_list != nil {
			if (edge_list.number > i) {
				sum += edge_list.length;
			}
			edge_list = edge_list.edge_list;
		}
	}
	fmt.Println(sum);

}
