package main;

import (
	"fmt";
	"strings";
	"bufio";
	"os";
	)

type my interface { int | string }

///////////////////////////////
// ---------- PARSER ----------

type parser struct {
	data, parsingObject, ident string;
	ident_list, idents []string;
	pos, ident_number, expr_number, formula_number int;
	next byte; // character at 'pos' position
	err bool;
	where_defined map[string]int;
	identmap map[string] map[string]bool;
}

func (p *parser) parse (obj string) (err bool) {
	if (p.err) { return true; }
	oldParsingObject := p.parsingObject;
	p.parsingObject = obj;
	// fmt.Println("parsing", obj);
	switch obj {
		case "program": 			err = p.parseProgram();
		case "definition": 			err = p.parseDefinition();
		case "ident": 			 	err = p.parseIdent();
		case "number":				err = p.parseNumber();
		case "expr": 				err = p.parseExpr();
		case "expr-2": 				err = p.parseExpr2();
		case "ident-list": 			err = p.parseIdentList();
		case "ident-list-2": 		err = p.parseIdentList2();
		case "term":				err = p.parseTerm();
		case "term-2":				err = p.parseTerm2();
		case "factor":				err = p.parseFactor();
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
	p.identmap = make(map[string]map[string]bool);
	p.ident = "";
	p.formula_number = 0;
	p.where_defined = make(map[string]int);
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
	p.parse("definition");
	for _, ident := range p.idents {
		p.where_defined[ident] = p.formula_number;
	}
	p.formula_number++;
	if (p.next != '$') {
		p.parse("program");
	}
	// if p.next == $, no error, return false
	return false;
}

func (p *parser) parseDefinition () (err bool) {
	p.idents = make([]string, 0);
	p.parse("ident-list");
	if (p.goNext() != '=') { return true; }
	p.expr_number = 1;
	p.parse("expr-list");
	if len(p.idents) != p.expr_number { return true; }
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

func (p *parser) parseExpr () (err bool) {
	p.parse("term");
	p.parse("expr-2");
	return false;
}

func (p *parser) parseExpr2 () (err bool) {
	if (p.next == '+' || p.next == '-') {
		p.goNext();
		p.parse("expr");
	}
	return false;
}

func (p *parser) parseIdentList () (err bool) {
	p.parse("ident");
	// new ident
	if _, exists := p.identmap[p.ident]; exists {
		// that means that ident was in another definition
		return true;
	}
	p.identmap[p.ident] = make(map[string]bool);
	p.idents = append(p.idents, p.ident);
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
		p.parse("ident");
		// new ident
		ident := p.idents[p.expr_number - 1];
		p.identmap[ident][p.ident] = true;
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

func (p *parser) parseExprList () (err bool) {
	p.parse("expr");
	p.parse("expr-list-2");
	return false;
}

func (p *parser) parseExprList2 () (err bool) {
	if (p.next == ',') {
		p.expr_number++;
		p.goNext();
		p.parse("expr-list");
	}
	return false;
}

////////////////////////////////////////
// ---------- GRAPH FUNCTIONS ----------

type edge[T my] struct {
	to int;
	edge_list *edge[T];
}

type vertex[T my] struct {
	number int;
	mark bool;
	edge_list *edge[T];
}

type graph[T my] struct {
	vertex_amount int;
	vertexes []vertex[T];
}

func (g *graph[T]) init (n int) {
	g.vertex_amount = n;
	g.vertexes = make([]vertex[T], n);
	for i := range g.vertexes {
		g.vertexes[i].number = i;
	}
}

func (g *graph[T]) addOrientedEdge (x, y int) *edge[T] {
	new_edge := new(edge[T]);
	new_edge.to = y;
	old_edge := g.vertexes[x].edge_list;
	g.vertexes[x].edge_list = new_edge;
	g.vertexes[x].edge_list.edge_list = old_edge;
	return new_edge;
}

func (g *graph[T]) addOrientedUniqueEdge (x, y int) *edge[T] {
	for e := g.vertexes[x].edge_list; e != nil; e = e.edge_list {
		if (e.to == y) { return e; }
	}
	return g.addOrientedEdge(x, y);
}

func (g *graph[T]) markFalse () {
	for i := range g.vertexes {
		g.vertexes[i].mark = false;
	}
}

func (g *graph[T]) isCyclic_rec (v int, inRoute *[]bool) bool {
	g.vertexes[v].mark = true;
	(*inRoute)[v] = true;

	for e := g.vertexes[v].edge_list; e != nil; e = e.edge_list {
		if (!g.vertexes[e.to].mark && g.isCyclic_rec(e.to, inRoute)) {
			return true;
		} else if ((*inRoute)[e.to]) {
			return true;
		}
	}

	(*inRoute)[v] = false;
	return false;
}

func (g *graph[T]) isCyclic () bool {
	inRoute := make([]bool, g.vertex_amount);
	for i := range g.vertexes {
		if (!g.vertexes[i].mark && g.isCyclic_rec(i, &inRoute)) {
			return true;
		}
	}
	return false;
}

func (g *graph[T]) topologicalSort_rec (v int, q *queue[int]) {
	g.vertexes[v].mark = true;
	for e := g.vertexes[v].edge_list; e != nil; e = e.edge_list {
		if (!g.vertexes[e.to].mark) {
			g.topologicalSort_rec(e.to, q);
		}
	}
	q.push(v);
}

func (g *graph[T]) topologicalSort () (q queue[int]) {
	q.init();
	for i := range g.vertexes {
		if (!g.vertexes[i].mark) {
			g.topologicalSort_rec(i, &q);
		}
	}
	return q;
}

//////////////////////////////
// ---------- QUEUE ----------

type queue[T my] struct {
	data []T;
}
func (q *queue[T]) init() {
	q.data = make([]T, 0);
}
func (q *queue[T]) push(x T) {
	q.data = append(q.data, x);
}
func (q *queue[T]) pop() T {
	x := q.data[0];
	q.data = q.data[1:];
	return x;
}
func (q *queue[T]) isEmpty() bool {
	return len(q.data) == 0;
}

func main() {
	var formula_def, formula_dependent int;
	var p parser;
	var g graph[string];
	var text, temp string;
  var ok bool;
	writer := bufio.NewWriter(os.Stdout);

	scanner := bufio.NewScanner(os.Stdin);
	formula_list := make([]string, 0);
    for scanner.Scan() {
    	temp = scanner.Text();
    	formula_list = append(formula_list, temp);
    	text += temp + ";";
	}

	if (p.start(text)) {
		fmt.Println("syntax error");
		return;
	}

	g.init(p.formula_number);
  

	for def, v := range p.identmap {
		formula_def = p.where_defined[def];
		for dependent := range v {
			formula_dependent, ok = p.where_defined[dependent];
      if (!ok) {
        fmt.Println("syntax error");
		    return;
      }
			g.addOrientedUniqueEdge(formula_def, formula_dependent);
		}
	}

	if (g.isCyclic()) {
		fmt.Println("cycle");
		return;
	}

	g.markFalse();
	q := g.topologicalSort();

	for !q.isEmpty() {
		v := q.pop();
		fmt.Fprintln(writer, formula_list[v]);
	}

	writer.Flush();
}
