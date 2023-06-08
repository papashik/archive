package main

import (
    "fmt";
    "bufio";
    "os";
    "strconv";
    )

type list struct {
	data []int;
	index int;
}

func (l *list) init (n int) {
	l.data = make([]int, n);
	l.index = 0;
}

func (l *list) assoc (v int) {
	l.data[v] = l.index;
	l.index++;
}

func dfs (v int, matr_trans, matr_answer *[][]int, matr_output, matr_answer_output *[][]string, bool_list *[]bool, index_map *list) {
	(*bool_list)[v] = true;
	for i, to := range (*matr_trans)[v] {
		if ((*bool_list)[to]) {
			(*matr_answer)[index_map.data[v]][i] = index_map.data[to];
			(*matr_answer_output)[index_map.data[v]][i] = (*matr_output)[v][i];
			continue;
		}
		index_map.assoc(to);
		(*matr_answer)[index_map.data[v]][i] = index_map.data[to];
		(*matr_answer_output)[index_map.data[v]][i] = (*matr_output)[v][i];
		dfs(to, matr_trans, matr_answer, matr_output, matr_answer_output, bool_list, index_map);
	}
}

func main() {
	var n, m, q0 int;
	var index_map list;
	writer := bufio.NewWriter(os.Stdout);
	reader := bufio.NewReader(os.Stdin);

	fmt.Fscan(reader, &n); // amount of statuses
	fmt.Fscan(reader, &m); // size of input alphabet
	fmt.Fscan(reader, &q0); // number of start status

	matrix_of_transitions := make([][]int, n);
	matrix_answer := make([][]int, n);
	matrix_of_outputs := make([][]string, n);
	matrix_answer_output := make([][]string, n);
	for i := 0; i < n; i++ {
		matrix_of_transitions[i] = make([]int, m);
		matrix_answer[i] = make([]int, m);
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &matrix_of_transitions[i][j]);
		}
	}
	for i := 0; i < n; i++ {
		matrix_of_outputs[i] = make([]string, m);
		matrix_answer_output[i] = make([]string, m);
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &matrix_of_outputs[i][j]);
		}
	}

	index_map.init(n);
	bool_list := make([]bool, n);
	index_map.assoc(q0);
	dfs(q0, &matrix_of_transitions, &matrix_answer, &matrix_of_outputs, &matrix_answer_output, &bool_list, &index_map);

	writer.WriteString(strconv.Itoa(n));
	writer.WriteByte('\n');
	writer.WriteString(strconv.Itoa(m));
	writer.WriteByte('\n');
	writer.WriteByte('0');
	writer.WriteByte('\n');
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			writer.WriteString(strconv.Itoa(matrix_answer[i][j]));
			writer.WriteByte(' ');
		}
		writer.WriteByte('\n');
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			writer.WriteString(matrix_answer_output[i][j]);
			writer.WriteByte(' ');
		}
		writer.WriteByte('\n');
	}
    writer.Flush();
}
