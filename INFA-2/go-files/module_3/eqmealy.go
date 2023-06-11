package main

import (
    "fmt";
    "bufio";
    "os";
    )
type my interface { int | string }

////////////////////////////////
// ---------- AUTOMAT ----------
type automat[T my] struct {
	// n - amount of statuses
	// m - size of input alphabet
	n, m int;
	trans [][]int;
	output [][]T;
}

func (a *automat[T]) init (matr_trans *[][]int, matr_output *[][]T) {
	a.n = len((*matr_trans));
	a.m = len((*matr_trans)[0]);
	a.trans = makeMatrix[int](a.n, a.m);
	a.output = makeMatrix[T](a.n, a.m);
	for i := 0; i < a.n; i++ {
		for j := 0; j < a.m; j++ {
			a.trans[i][j] = (*matr_trans)[i][j];
			a.output[i][j] = (*matr_output)[i][j]
		}
	}
}

/////////////////////////////////////
// ---------- CANONIZATION ----------
type indexList struct {
	data []int;
	marks []bool;
	index int;
}

func (l *indexList) init (n int) {
	l.data = make([]int, n);
	l.marks = make([]bool, n);
	l.index = 0;
}

func (l *indexList) setAssoc (v int) {
	l.data[v] = l.index;
	l.index++;
}

func (a *automat[T]) canonize_rec (v int, new_matr_trans *[][]int, new_matr_output *[][]T, list *indexList) {
	list.marks[v] = true;
	for i, to := range a.trans[v] {
		if (!list.marks[to]) {
			list.setAssoc(to);
			a.canonize_rec(to, new_matr_trans, new_matr_output, list);
		}
		(*new_matr_trans)[list.data[v]][i] = list.data[to];
		(*new_matr_output)[list.data[v]][i] = a.output[v][i];
	}
}

func (a *automat[T]) canonize (q0 int) {
	var list indexList;
	list.init(a.n);
	list.setAssoc(q0);

	new_matr_trans := makeMatrix[int](a.n, a.m);
	new_matr_output := makeMatrix[T](a.n, a.m);

	a.canonize_rec(q0, &new_matr_trans, &new_matr_output, &list);
	for i, mark := range list.marks {
		if (mark) { continue; }
		list.setAssoc(i);
		a.canonize_rec(i, &new_matr_trans, &new_matr_output, &list);
	}

	// copying data to automat
	for i := 0; i < a.n; i++ {
		for j := 0; j < a.m; j++ {
			a.trans[i][j] = new_matr_trans[i][j];
			a.output[i][j] = new_matr_output[i][j]
		}
	}
}

/////////////////////////////////////////
// ---------- MATRIX FUNCTIONS ----------

func readMatrix[T my] (reader *bufio.Reader, n, m int) [][]T {
 	matrix := make([][]T, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]T, m);
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &matrix[i][j]);
		}
	}
	return matrix;
}

func makeMatrix[T my] (n, m int) (matrix [][]T){
	matrix = make([][]T, n);
	for i := 0; i < n; i++ {
		matrix[i] = make([]T, m);
	}
	return matrix;
}

///////////////////////////////////////////
// ---------- DISJOINT SET UNION ----------
type dsu struct {
	parent, rank []int;
}

func (d *dsu) init (n int) {
	d.parent = make([]int, n);
	d.rank = make([]int, n);
	for i := 0; i < n; i++ {
		d.parent[i] = i;
	}
}

func (d *dsu) find (v int) int {
	if (v == d.parent[v]) { return v; }
	d.parent[v] = d.find(d.parent[v]);
	return d.parent[v];
}

func (d *dsu) union (a, b int) {
	a = d.find(a);
	b = d.find(b);
	if (a != b) {
		if (d.rank[a] < d.rank[b]) { a, b = b, a; }
		d.parent[b] = a;
		if (d.rank[a] == d.rank[b]) { d.rank[a]++; }
	}
}

/////////////////////////////////////
// ---------- MINIMIZATION ----------

func areRowsEqual[T my] (a, b int, matrix *[][]T) bool {
	for j := 0; j < len((*matrix)[0]); j++ {
		if ((*matrix)[a][j] != (*matrix)[b][j]) { return false; }
	}
	return true;
}

func is_in_array (x int, array *[]int) bool {
	for i := range (*array) {
		if (x == (*array)[i]) { return true; }
	}
	return false;
}

func updateEqv (eqv *[]int) {
	n := len(*eqv);
	for i := n - 2; i >= 0; i-- {
		if (!is_in_array(i, eqv)) {
			for j := 0; j < n; j++ {
				if ((*eqv)[j] > i) { (*eqv)[j] = (*eqv)[j] - 1; }
			}
		}
	}
}

func (a *automat[T]) split1 () (amount int, eqv []int) {
	var d dsu;
	amount = a.n;
	eqv = make([]int, a.n);
	d.init(a.n);
	for q1 := 0; q1 < a.n; q1++ {
		for q2 := q1 + 1; q2 < a.n; q2++ {
			if (d.find(q1) != d.find(q2)) {
				if (areRowsEqual(q1, q2, &a.output)) {
					d.union(q1, q2);
					amount--;
				}
			}
		}
	}

	for i := 0; i < a.n; i++ {
		eqv[i] = d.find(i);
	}
	updateEqv(&eqv);
	return amount, eqv;
}

func (a *automat[T]) split (eqv []int) (int, []int) {
	var d dsu;
	var eq bool;
	var trans1, trans2 int;
	var amount int;
	amount = a.n;
	d.init(a.n);
	for q1 := 0; q1 < a.n; q1++ {
		for q2 := q1 + 1; q2 < a.n; q2++ {
			if (eqv[q1] == eqv[q2] && d.find(q1) != d.find(q2)) {
				eq = true;
				for input := 0; input < a.m; input++ {
					trans1 = a.trans[q1][input];
					trans2 = a.trans[q2][input];
					if (eqv[trans1] != eqv[trans2]) {
						eq = false;
						break;
					}
				}
				if (eq) {
					d.union(q1, q2);
					amount--;
				}
			}
		}
	}
	for i := 0; i < a.n; i++ {
		eqv[i] = d.find(i);
	}
	updateEqv(&eqv);
	return amount, eqv;
}

func (a *automat[T]) minimize () automat[T] {
	var amount, amount_old, new_i int;
	var eqv []int;
	var a_new automat[T];

	amount_old, eqv = a.split1();
	for {
		amount, eqv = a.split(eqv);
		if (amount_old == amount) { break; }
		amount_old = amount;
	}

	new_trans := makeMatrix[int](amount, a.m);
	new_output := makeMatrix[T](amount, a.m);
	bool_list := make([]bool, a.n);
	for i := 0; i < a.n; i++ {
		new_i = eqv[i];
		if (!bool_list[new_i]) {
			bool_list[new_i] = true;
			for input := 0; input < a.m; input++ {
				new_trans[new_i][input] = eqv[a.trans[i][input]];
				new_output[new_i][input] = a.output[i][input];
			}
		}
	}
	a_new.init(&new_trans, &new_output);

	return a_new;
}


func areMatrixesEqual[T my] (matrix1, matrix2 *[][]T) bool {
	n := len((*matrix1));
	m := len((*matrix1)[0]);

	// checking size
	if (n != len((*matrix2)) || m != len((*matrix2)[0])) { return false; }

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if ((*matrix1)[i][j] != (*matrix2)[i][j]) { return false; }
		}
	}
	return true;
}

func readAutomat (reader *bufio.Reader) (automat[string]) {
	var n, m, q0 int;
	var a automat[string];
	fmt.Fscan(reader, &n); // amount of statuses
	fmt.Fscan(reader, &m); // size of input alphabet
	fmt.Fscan(reader, &q0); // number of start status

	matr_trans := readMatrix[int](reader, n, m);
	matr_output := readMatrix[string](reader, n, m);
	a.init(&matr_trans, &matr_output);

	a.canonize(q0);
	// now automat is canonized

	a = a.minimize();
	// now automat is minimized

	a.canonize(0);
	return a;
}

func main() {
	writer := bufio.NewWriter(os.Stdout);
	reader := bufio.NewReader(os.Stdin);

	a1 := readAutomat(reader);
	a2 := readAutomat(reader);

	if (areMatrixesEqual(&a1.trans, &a2.trans) && areMatrixesEqual(&a1.output, &a2.output)) {
		fmt.Fprintln(writer, "EQUAL");
	} else {
		fmt.Fprintln(writer, "NOT EQUAL");
	}

	writer.Flush();
}
