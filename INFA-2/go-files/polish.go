package main
import (
    "fmt";
    "bufio";
    "os";
    )

type StackType struct {
    size, top int;
    data []int;
}

func (stack *StackType) init(size int) {
    stack.size = size;
    stack.top = 0;
    stack.data = make([]int, size);
}

func (stack *StackType) push(elem int) (err int) {
    if (stack.size == stack.top) {
        return -1; // Stack overflow, err = -1
    }
    stack.data[stack.top] = elem;
    stack.top++;
    return 0;
}

func (stack *StackType) pop() (res int, err int) {
    if (stack.top == 0) {
        return 0, -1; // Stack is empty, err = -1
    }
    stack.top--;
    return stack.data[stack.top], 0;
}

func (stack *StackType) evalPolishExpression(text string) {
    length := len(text);
    for i := length - 1; i >= 0; i-- {
        switch text[i] {
            case '+', '-', '*':
                num1, _ := stack.pop();
                num2, _ := stack.pop();
                switch text[i] {
                    case '+':
                        stack.push(num1 + num2);
                    case '-':
                        stack.push(num1 - num2);
                    case '*':
                        stack.push(num1 * num2);
                }
            
            case ' ', '\n', '(', ')':
                continue;
                
            default: // number
                stack.push(int(text[i]) - '0');
        }
    }
}

func main() {
    var stack StackType;
    stack.init(100);
    text, _ := bufio.NewReader(os.Stdin).ReadString('\n');
    
    stack.evalPolishExpression(text);
    
    res, _ := stack.pop();
    fmt.Printf("%d\n", res);
}
