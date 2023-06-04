package main
import (
    "fmt";
    "bufio";
    "os";
    )

type StackType struct {
    size, top int;
    data []string;
}

func (stack *StackType) init(size int) {
    stack.size = size;
    stack.top = 0;
    stack.data = make([]string, size);
}

func (stack *StackType) push(elem string) (err int) {
    if (stack.size == stack.top) {
        return -1; // Stack overflow, err = -1
    }
    stack.data[stack.top] = elem;
    stack.top++;
    return 0;
}

func (stack *StackType) pop() (res string, err int) {
    if (stack.top == 0) {
        return "", -1; // Stack is empty, err = -1
    }
    stack.top--;
    return stack.data[stack.top], 0;
}

func (stack *StackType) contains(elem string) bool {
    for ind := 0; ind < stack.top; ind++ {
        if (elem == stack.data[ind]) {
            return true;
        }
    }
    return false;
}

func (stack *StackType) countEvaluations(text string) (result int) {
    var AlreadyEvalExprs StackType;
    AlreadyEvalExprs.init(100);
    
    result = 0;
    
    length := len(text);
    for i := length - 1; i >= 0; i-- {
        switch text[i] {
            case '#', '$', '@':
                str1, _ := stack.pop();
                str2, _ := stack.pop();
                expr := "(" + str1 + string(text[i]) + str2 + ")";
                
                if !(AlreadyEvalExprs.contains(expr)) {
                    // Если это выражение еще не было вычислено
                    AlreadyEvalExprs.push(expr);
                    result++;
                }
                stack.push(expr);
            
            case ' ', '\n', '(', ')':
                continue;
                
            default: // letter
                stack.push(string(text[i]));
        }
    }
    
    return;
}

func main() {
    var stack StackType;
    stack.init(100);
    text, _ := bufio.NewReader(os.Stdin).ReadString('\n');
    
    res := stack.countEvaluations(text);
    
    fmt.Printf("%d\n", res);
}
