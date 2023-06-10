package main

import (
    "fmt";
    "bufio";
    "os";
    )

type edge struct {
	to int;
	input string;
	edge_list *edge;
}

type vertex struct {
	number, status int;
	output string;
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

func (g *graph) printOriented (writer *bufio.Writer) {
	fmt.Fprintf(writer, "digraph {\n\trankdir = LR\n");
	for i := range g.vertexes {
		fmt.Fprintf(writer, "\t%d [label = \"(%d,%s)\"]\n", i, g.vertexes[i].status, g.vertexes[i].output);
	}
	for i := range g.vertexes {
		for e := g.vertexes[i].edge_list; e != nil; e = e.edge_list {
			fmt.Fprintf(writer, "\t%d -> %d [label = \"%s\"]\n", i, e.to, e.input);
		}
	}
	fmt.Fprintln(writer, "}");
}

func readMatrix[T int | string] (reader *bufio.Reader, n, m int) [][]T {
 	matrix := make([][]T, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]T, m);
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &matrix[i][j]);
		}
	}
	return matrix;
}

func readSlice[T int | string] (reader *bufio.Reader, n int) []T {
 	slice := make([]T, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &slice[i]);
	}
	return slice;
}

func (g *graph) getVertexNumber (status int, output string) int {
	for i := range g.vertexes {
		if (g.vertexes[i].status == status && g.vertexes[i].output == output) {
			return i;
		}
	}
	return -1;
}

func main() {
	var x, y, n, status, output_number, new_edge_status int;
	var new_edge_output string;
	var g graph;
	var new_edge *edge;

	writer := bufio.NewWriter(os.Stdout);
	reader := bufio.NewReader(os.Stdin);

	fmt.Fscan(reader, &x); // size of input alphabet
	input_alphabet := readSlice[string](reader, x);
	fmt.Fscan(reader, &y); // size of output alphabet
	output_alphabet := readSlice[string](reader, y);
	fmt.Fscan(reader, &n); // amount of statuses

	matr_trans := readMatrix[int](reader, n, x);
	matr_output := readMatrix[int](reader, n, x);

	set := make(map[[2]int]bool);
	// set that contains possible pairs [status, output]
	// equivalent to vertexes in graph
	for i := 0; i < n; i++ {
		for j := 0; j < x; j++ {
			set[[2]int{matr_trans[i][j], matr_output[i][j]}] = true;
		}
	}

	g.init(len(set));
	ind := 0;
	for pair := range set {
		status = pair[0];
		output_number = pair[1];
		g.vertexes[ind].status = status;
		g.vertexes[ind].output = output_alphabet[output_number];
		ind++;
	}

	for i := range g.vertexes {
		status = g.vertexes[i].status;
		for input_number, input := range input_alphabet {
			new_edge_status = matr_trans[status][input_number];
			new_edge_output = output_alphabet[matr_output[status][input_number]];
			new_edge = g.add_oriented_edge(i, g.getVertexNumber(new_edge_status, new_edge_output));
			new_edge.input = input;
		}
	}


	g.printOriented(writer);

	writer.Flush();
}
