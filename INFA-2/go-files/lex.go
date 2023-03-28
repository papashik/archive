package main
import (
    "fmt";
    "bufio";
    "os";
    "strings";
    )

type AssocStruct struct {
    keys []string;
    values []int;
    amount, top int;
}

func (assoc* AssocStruct) init(size int) {
    assoc.keys = make([]string, size);
    assoc.values = make([]int, size);
    assoc.amount = 0;
    assoc.top = size;
}

func (assoc* AssocStruct) Assign(s string, x int) {
	ind, exists := assoc.Find(s);
	if (exists) {
		assoc.values[ind] = x;
	} else if (assoc.amount != assoc.top) {
		assoc.keys[assoc.amount] = s;
		assoc.values[assoc.amount] = x;
		assoc.amount++;
	}
}

func (assoc* AssocStruct) Find(s string) (ind int, exists bool) {
	for ind := 0; ind < assoc.amount; ind++ {
		if (assoc.keys[ind] == s) {
			return ind, true;
		}
	}
	return 0, false;
}

func (assoc* AssocStruct) Lookup(s string) (x int, exists bool) {
	x, exists = assoc.Find(s);
	if (exists) {
		x = assoc.values[x];
	}
	return;
}

type AssocArray interface {
    Assign(s string, x int);
    Lookup(s string) (x int, exists bool);
}

func lex(sentence string, array AssocArray) []int {
    words := strings.Split(sentence, " ");
	res_arr := make([]int, len(words));
	max_id := 0;
    for ind, word := range words {
        x, exists := array.Lookup(word);
        if (exists) {
			res_arr[ind] = x;
		} else {
			max_id++;
			array.Assign(word, max_id);
			res_arr[ind] = max_id;
		}
    }
	return res_arr;
}

func main() {
	var assoc AssocStruct;
	// assocArrayObj := AssocArray(assoc);
	assoc.init(100);
    // text, _ := bufio.NewReader(os.Stdin).ReadString('\n');
    scanner := bufio.NewScanner(os.Stdin);
    scanner.Scan();
    text := scanner.Text();

	res_arr := lex(text, &assoc);
	for _, x := range res_arr {
		fmt.Printf("%d ", x);
	}
	fmt.Println();
}
