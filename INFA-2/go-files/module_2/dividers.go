package main

import (
    "fmt";
    )

type edge struct {
	number1 int;
	number2 int;
}

type vertex struct {
	number int;
	vert_map map [int]int;
	list* edge;
}

func vert_init(vert *vertex, x int) {
	vert.number = x;
	vert.vert_map = make(map[int]int);
	vert.list = nil;
}

func find_vertex(vertexes *[]*vertex, x int) *vertex {
	for _, v := range *vertexes {
		if (v.number == x) {
			return v;
		}
	}
	return nil;
}

func find_edge(edges *[]*edge, x, y int) *edge {
	for _, edge := range *edges {
		if ((edge.number1 == x && edge.number2 == y) || (edge.number1 == y && edge.number2 == x)) {
			return edge;
		}
	}
	return nil;
}

func main() {
	var x int;
	fmt.Scanf("%d", &x);

	var vert vertex;
	vert_init(&vert, x);

	n := x;
	for i := 2; i <= x/2; i++ {
		for n % i == 0 {
			n /= i;
			vert.vert_map[i]++;
		}
	}
	if (len(vert.vert_map) == 0 && x != 1) {
		vert.vert_map[x]++;
	}

	vertexes := make([]*vertex, 0);
	edges := make([]*edge, 0);

	working := make([]*vertex, 0);
	working = append(working, &vert);
	vertexes = append(vertexes, &vert);

	for len(working) != 0 {
		work_vert := working[0];
		working = working[1:];
		parent_number := work_vert.number;

		for k, _ := range work_vert.vert_map {
			if find_vertex(&vertexes, parent_number/k) == nil {
				var new_vert vertex;
				vert_init(&new_vert, parent_number/k);

				for key, value := range work_vert.vert_map {
					new_vert.vert_map[key] = value;
				}
				new_vert.vert_map[k]--;
				if new_vert.vert_map[k] == 0 {
					delete(new_vert.vert_map, k);
				}
				working = append(working, &new_vert);
				vertexes = append(vertexes, &new_vert);
			}

			if find_edge(&edges, parent_number, parent_number/k) == nil {
				var new_edge edge;
				new_edge.number1 = parent_number;
				new_edge.number2 = parent_number/k;
				edges = append(edges, &new_edge);
			}
		}
	}
	fmt.Printf("graph {\n");
	for _, v := range vertexes {
		fmt.Printf("%d\n", v.number);
	}
	for _, edge := range edges {
		fmt.Printf("%d--%d\n", edge.number1, edge.number2);
	}
	fmt.Printf("}");
}
