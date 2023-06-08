package main

import (
    "fmt";
    "bufio";
    "os";
    )

type edge struct {
	to int;
	input byte;
	output string;
	edge_list *edge;
}

type vertex struct {
	number int;
	edge_list *edge;
}

type graph struct {
	vertex_amount int;
	vertexes []vertex;
}

func (g *graph) init (n int) {
	g.vertex_amount = n;
	g.vertexes = make([]vertex, n);
	for i := range g.vertexes {
		g.vertexes[i].number = i;
	}
}

func (g *graph) add_oriented_edge (x, y int) *edge {
	new_edge := new(edge);
	new_edge.to = y;
	old_edge := g.vertexes[x].edge_list;
	g.vertexes[x].edge_list = new_edge;
	g.vertexes[x].edge_list.edge_list = old_edge;
	return new_edge;
}

func (g *graph) print_oriented () {
	writer := bufio.NewWriter(os.Stdout);
	fmt.Fprintln(writer, "digraph {");
	fmt.Fprintf(writer, "\trankdir = LR\n");
	for i := range g.vertexes {
		fmt.Fprintf(writer, "\t%d\n", i);
	}
	for i := range g.vertexes {
		for e := g.vertexes[i].edge_list; e != nil; e = e.edge_list {
			fmt.Fprintf(writer, "\t%d -> %d [label = \"%c(%s)\"]\n", i, e.to, e.input, e.output);
		}
	}
	fmt.Fprintln(writer, "}");
	writer.Flush();
}

func main() {
	var n, m, q0 int;
	var g graph;
	var new_edge *edge;
	reader := bufio.NewReader(os.Stdin);

	fmt.Fscan(reader, &n); // amount of statuses
	fmt.Fscan(reader, &m); // size of input alphabet
	fmt.Fscan(reader, &q0); // number of start status

	matrix_of_transitions := make([][]int, n);
	matrix_of_outputs := make([][]string, n);
	for i := 0; i < n; i++ {
		matrix_of_transitions[i] = make([]int, m);
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &matrix_of_transitions[i][j]);
		}
	}
	for i := 0; i < n; i++ {
		matrix_of_outputs[i] = make([]string, m);
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &matrix_of_outputs[i][j]);
		}
	}

	g.init(n);
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			new_edge = g.add_oriented_edge(i, matrix_of_transitions[i][j]);
			new_edge.input = byte(97 + j);
			new_edge.output = matrix_of_outputs[i][j];
		}
	}

	g.print_oriented();
}
