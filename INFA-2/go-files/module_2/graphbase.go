package main

import (
    "fmt";
    "sort";
    )

type edge struct {
	to int;
	edge_list *edge;
}

type vertex struct {
	number int;
	mark bool;
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

func (g *graph) add_oriented_edge (x, y int) {
	new_edge := new(edge);
	new_edge.to = y;
	old_edge := g.vertexes[x].edge_list;
	g.vertexes[x].edge_list = new_edge;
	g.vertexes[x].edge_list.edge_list = old_edge;
}

func (g *graph) add_edge (x, y int) {
	g.add_oriented_edge(x, y);
	g.add_oriented_edge(y, x);
}

func (g *graph) mark_false () {
	for i := range g.vertexes {
		g.vertexes[i].mark = false;
	}
}

func (g *graph) dfs1 (v int, order *[]int) {
	g.vertexes[v].mark = true;
	for e := g.vertexes[v].edge_list; e != nil; e = e.edge_list {
		if (!g.vertexes[e.to].mark) { g.dfs1(e.to, order); }
	}
	(*order) = append((*order), v);
}

func (gt *graph) dfs2 (v int, component *[]int) {
	gt.vertexes[v].mark = true;
	(*component) = append((*component), v);
	for e := gt.vertexes[v].edge_list; e != nil; e = e.edge_list {
		if (!gt.vertexes[e.to].mark) { gt.dfs2(e.to, component); }
	}
}

func (g *graph) make_transposed_graph () (gt graph) {
	gt.init(g.vertex_amount);
	for i := range g.vertexes {
		for e := g.vertexes[i].edge_list; e != nil; e = e.edge_list {
			gt.add_oriented_edge(e.to, i);
		}
	}
	return;
}

func min_in_array (array *[]int) (min int) {
	min = (*array)[0];
	for i := range (*array) {
		if ((*array)[i] < min) { min = (*array)[i]; }
	}
	return;
}

func is_in_array (x int, array *[]int) bool {
	for i := range (*array) {
		if (x == (*array)[i]) { return true; }
	}
	return false;
}

func (g *graph) find_strong_components () [][]int {
	var v int;
	gt := g.make_transposed_graph();
	order := make([]int, 0);
	g.mark_false();
	for i := range g.vertexes {
		// first DFS, making vertex order
		if (!g.vertexes[i].mark) { g.dfs1(i, &order); }
	}

	component := make([]int, 0);
	components := make([][]int, 0);
	gt.mark_false();
	for i := range order {
		//second DFS, making components
		v = order[g.vertex_amount-i-1];
		if (!gt.vertexes[v].mark) {
			gt.dfs2(v, &component);
			// new component
			components = append(components, component);
			component = nil;
		}
	}
	return components;
}

func main() {
	var n, m, x, y, min int;
	var g, gt graph;
	var x_is_base bool;
	fmt.Scanf("%d", &n); // vertexes
	fmt.Scanf("%d", &m); // edges

	g.init(n);

	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d", &x, &y);
		g.add_oriented_edge(x, y);
	}

	strong_comps := g.find_strong_components();

	// possible_base - минимальные вершины из компонент связности
	possible_base := make([]int, len(strong_comps));
	for i := range strong_comps {
		min = min_in_array(&strong_comps[i]);
		possible_base[i] = min;
	}

	/*
	Если в gt(g transposed) запустить DFS2 из какой-то
	данной вершины possible_base, то, если мы встретим \
	какую-либо другую вершину из possible_base,
	то данная не является базой
	*/
	base := make([]int, 0);
	component := make([]int, 0);
	gt = g.make_transposed_graph();
	for i := range possible_base {
		// Берем данную вершину x и запускаем DFS2 по gt
		x = possible_base[i];
		gt.mark_false();
		gt.dfs2(x, &component);
		component = component[1:]; // убираем данную
		// В component лежат все встреченные нами вершины
		// Если это одна из possible_base, то x - не база
		x_is_base = true;
		for j := range component {
			if is_in_array(component[j], &possible_base) {
				x_is_base = false;
				break;
			}
		}
		if (x_is_base) { base = append(base, x); }
		component = nil;
	}
	sort.Sort(sort.IntSlice(base));
	for i := range base {
		fmt.Printf("%d ", base[i]);
	}

}
