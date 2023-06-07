package main;

import (
	"fmt";
	"strings";
	"bufio";
	"os";
	)

type parser struct {
	data, parsingObject, ident, curfunc string;
	pos int;
	next byte; // character at 'pos' position
	err bool;
	funcmap map[string] []string;
}

func (p *parser) parse (obj string) (err bool) {
	if (p.err) { return true; }
	oldParsingObject := p.parsingObject;
	p.parsingObject = obj;
	// fmt.Println("parsing", obj);
	switch obj {
		case "program": 			err = p.parseProgram();
		case "function": 			err = p.parseFunction();
		case "ident": 			 	err = p.parseIdent();
		case "number":				err = p.parseNumber();
		case "formal-args-list": 	err = p.parseFormalArgsList();
		case "expr": 				err = p.parseExpr();
		case "expr-2": 				err = p.parseExpr2();
		case "comparison-expr": 	err = p.parseComparisonExpr();
		case "comparison-expr-2": 	err = p.parseComparisonExpr2();
		case "ident-list": 			err = p.parseIdentList();
		case "ident-list-2": 		err = p.parseIdentList2();
		case "arith-expr": 			err = p.parseArithExpr();
		case "arith-expr-2": 		err = p.parseArithExpr2();
		case "comparison-operator": err = p.parseComparisonOperator();
		case "term":				err = p.parseTerm();
		case "term-2":				err = p.parseTerm2();
		case "factor":				err = p.parseFactor();
		case "ident-2":				err = p.parseIdent2();
		case "ident-3":				err = p.parseIdent3();
		case "actual-args-list": 	err = p.parseActualArgsList();
		case "expr-list":			err = p.parseExprList();
		case "expr-list-2":			err = p.parseExprList2();
		default:					fmt.Println("ERROR UNKNOWN");
	}
	if (p.err) { return true; }
	if (!err) {
		p.parsingObject = oldParsingObject;
	} else {
		p.err = true;
		// fmt.Printf("error at %d position while parsing %s\n", p.pos - 1, obj);
	}
	return;
}

func (p *parser) start (data string) (err bool) {
	p.data = data;
	if (p.miniParse()) { return true; }
	// удаляем все пробелы, добавляем конечный символ
	p.data = strings.ReplaceAll(strings.ReplaceAll(data, " ", ""), "\n", "") + "$";
	p.parsingObject = "none";
	p.funcmap = make(map[string] []string);
	p.curfunc = "";
	p.ident = "";
	p.pos = 0;
	p.err = false;
	p.next = p.data[0];
	p.parse("program");
	return p.err;
}

func (p *parser) miniParse () (err bool) {
	var s byte;
	was_ident := false;
	was_space := false;
	for i := range p.data {
		s = p.data[i];
		if (was_ident && was_space && (isDigit(s) || isLetter(s))) {
			return true;
		} else if (isDigit(s) || isLetter(s)) {
			was_ident = true;
			was_space = false;
		} else if (s == ' ' || s == '\n') {
			was_space = true;
		} else {
			was_ident = false;
		}
	}
	return false;
}

func (p *parser) goNext () byte {
	// returns current symbol and going next
	if (p.next == '$') { return p.next; }
	old := p.next;
	p.pos++;
	p.next = p.data[p.pos];
	return old;
}

func (p *parser) parseProgram () (err bool) {
	p.parse("function");
	if (p.next != '$') {
		p.parse("program");
	}
	// if p.next == $, no error, return false
	return false;
}

func (p *parser) parseFunction () (err bool) {
	p.parse("ident");
	p.curfunc = p.ident;
	p.funcmap[p.curfunc] = make([]string, 0);
	if (p.goNext() != '(') { return true; }
	p.parse("formal-args-list");
	if (p.goNext() != ')') { return true; }
	if (p.goNext() != ':') { return true; }
	if (p.goNext() != '=') { return true; }
	p.parse("expr");
	if (p.goNext() != ';') { return true; }
	return false;
}

func isLetter (symbol byte) bool {
	if ((symbol < 'a' || symbol > 'z') && (symbol < 'A' || symbol > 'Z')) {
		return false;
	}
	return true;
}

func isDigit (symbol byte) bool {
	if (symbol < '0' || symbol > '9') { return false; }
	return true;
}

func (p *parser) parseIdent () (err bool) {
	// p.next must be letter
	// and then letters and digits
	start_pos := p.pos;
	if (!isLetter(p.goNext())) { return true; }
	for (isDigit(p.next) || isLetter(p.next)) {
		p.goNext();
	}
	p.ident = string(p.data[start_pos:p.pos]);
	return false;
}

func (p *parser) parseNumber () (err bool) {
	for (isDigit(p.next)) {
		p.goNext();
	}
	return false;
}

func (p *parser) parseFormalArgsList () (err bool) {
	if (isLetter(p.next)) { p.parse("ident-list"); }
	return false;
}

func (p *parser) parseExpr () (err bool) {
	p.parse("comparison-expr");
	p.parse("expr-2");
	return false;
}

func (p *parser) parseExpr2 () (err bool) {
	if (p.next == '?') {
		p.goNext();
		p.parse("comparison-expr");
		if (p.goNext() != ':') { return true; }
		p.parse("expr");
	}
	return false;
}

func (p *parser) parseIdentList () (err bool) {
	p.parse("ident");
	p.parse("ident-list-2");
	return false;
}

func (p *parser) parseIdentList2 () (err bool) {
	if (p.next == ',') {
		p.goNext();
		p.parse("ident-list");
	}
	return false;
}

func (p *parser) parseComparisonExpr () (err bool) {
	p.parse("arith-expr");
	p.parse("comparison-expr-2");
	return false;
}

func (p *parser) parseComparisonExpr2 () (err bool) {
	if (p.next == '<' || p.next == '>' || p.next == '=') {
		p.parse("comparison-operator");
		p.parse("arith-expr");
	}
	return false;
}

func (p *parser) parseArithExpr () (err bool) {
	p.parse("term");
	p.parse("arith-expr-2");
	return false;
}

func (p *parser) parseArithExpr2 () (err bool) {
	if (p.next == '+' || p.next == '-') {
		p.goNext();
		p.parse("arith-expr");
	}
	return false;
}

func (p *parser) parseComparisonOperator () (err bool) {
	prev := p.goNext();
	if ((prev == '<' && p.next == '>') ||
		(prev == '<' && p.next == '=') ||
		(prev == '>' && p.next == '=')) {
		p.goNext();
	}
	return false;
}

func (p *parser) parseTerm () (err bool) {
	p.parse("factor");
	p.parse("term-2");
	return false;
}

func (p *parser) parseTerm2 () (err bool) {
	if (p.next == '*' || p.next == '/') {
		p.goNext();
		p.parse("term");
	}
	return false;
}

func (p *parser) parseFactor () (err bool) {
	if (isLetter(p.next)) {
		p.parse("ident-2");
	} else if (isDigit(p.next)) {
		p.parse("number");
	} else if (p.next == '(') {
		p.goNext();
		p.parse("expr");
		if (p.goNext() != ')') { return true; }
	} else if (p.next == '-') {
		p.goNext();
		p.parse("factor");
	} else {
		return true;
	}

	return false;
}

func (p *parser) parseIdent2 () (err bool) {
	p.parse("ident");
	p.parse("ident-3");
	return false;
}

func (p *parser) parseIdent3 () (err bool) {
	if (p.next == '(') {
		// вот здесь была вызвана функция
		p.funcmap[p.curfunc] = append(p.funcmap[p.curfunc], p.ident);
		p.goNext();
		p.parse("actual-args-list");
		if (p.goNext() != ')') { return true; }
	}
	return false;
}

func (p *parser) parseActualArgsList () (err bool) {
	if (isDigit(p.next) || isLetter(p.next) || p.next == '(' || p.next == '-') {
		p.parse("expr-list");
	}
	return false;
}

func (p *parser) parseExprList () (err bool) {
	p.parse("expr");
	p.parse("expr-list-2");
	return false;
}

func (p *parser) parseExprList2 () (err bool) {
	if (p.next == ',') {
		p.goNext();
		p.parse("expr-list");
	}
	return false;
}

type edge struct {
	to string;
	edge_list *edge;
}

type vertex struct {
	mark bool;
	edge_list *edge;
}

type graph struct {
	vertex_amount int;
	vertexes map[string]*vertex;
}

func (g *graph) init (n int) {
	g.vertex_amount = n;
	g.vertexes = make(map[string]*vertex, n);
}

func (g *graph) add_oriented_edge (x, y string) {
	new_edge := new(edge);
	new_edge.to = y;
	if (g.vertexes[x] == nil) { g.vertexes[x] = new(vertex); }
	if (g.vertexes[y] == nil) { g.vertexes[y] = new(vertex); }
	old_edge := g.vertexes[x].edge_list;
	g.vertexes[x].edge_list = new_edge;
	g.vertexes[x].edge_list.edge_list = old_edge;
}

func (g *graph) add_edge (x, y string) {
	g.add_oriented_edge(x, y);
	g.add_oriented_edge(y, x);
}

func (g *graph) mark_false () {
	for i := range g.vertexes {
		g.vertexes[i].mark = false;
	}
}

func (g *graph) dfs1 (v string, order *[]string) {
	g.vertexes[v].mark = true;
	for e := g.vertexes[v].edge_list; e != nil; e = e.edge_list {
		if (g.vertexes[e.to] != nil && !g.vertexes[e.to].mark) { g.dfs1(e.to, order); }
	}
	(*order) = append((*order), v);
}

func (gt *graph) dfs2 (v string, component *[]string) {
	gt.vertexes[v].mark = true;
	(*component) = append((*component), v);
	for e := gt.vertexes[v].edge_list; e != nil; e = e.edge_list {
		if (gt.vertexes[e.to] != nil && !gt.vertexes[e.to].mark) { gt.dfs2(e.to, component); }
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

func (g *graph) find_strong_components () [][]string {
	var v string;
	gt := g.make_transposed_graph();
	order := make([]string, 0);
	g.mark_false();
	for i := range g.vertexes {
		// first DFS, making vertex order
		if (!g.vertexes[i].mark) { g.dfs1(i, &order); }
	}

	component := make([]string, 0);
	components := make([][]string, 0);
	gt.mark_false();
	for i := range order {
		//second DFS, making components
		v = order[len(order)-i-1];
		if (gt.vertexes[v] == nil) {
			gt.vertexes[v] = new(vertex);
		}
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
	var p parser;
	var g graph;
	var text string;
	g.init(100);
	scanner := bufio.NewScanner(os.Stdin);
    for scanner.Scan() {
    	text += scanner.Text();
	}

	if (p.start(text)) {
		fmt.Println("error");
		return;
	}

	for key, _ := range p.funcmap {
		g.vertexes[key] = new(vertex);
	}

	for k, v := range p.funcmap {
		for _, another_func := range v {
			g.add_oriented_edge(k, another_func);
		}
	}

	comps := g.find_strong_components();

	fmt.Println(len(comps));

}
