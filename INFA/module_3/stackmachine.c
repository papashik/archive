#include <stdio.h>
#include <string.h>
#include <stdlib.h>

typedef struct struct_Stack {
    int cap, top, *data;
} stack_type;

void StackInit(stack_type *stack, int n) {
    stack->cap = n;
    stack->top = 0;
    stack->data = (int*)calloc(n, sizeof(int));
}

char StackEmpty(stack_type *stack) {
    return stack->top == 0 ? 1 : 0;
}

void StackPush(stack_type *stack, int new_num) {
    if (stack->top == stack->cap) {
        printf("Stack overflow!\n");
    } else {
        stack->data[stack->top] = new_num;
        stack->top++;
    }
}

int StackPop(stack_type *stack) {
    if (StackEmpty(stack) == 1) {
        printf("Stack is empty!\n");
        return -1;
    } else {
        stack->top--;
        return stack->data[stack->top];
    }
}

void StackPopAndTell(stack_type *stack) {
    printf("%d\n", StackPop(stack));
}

int main()
{
    stack_type stack;
    int n = 100000, x = 0, y = 0;
    StackInit(&stack, n);
    
    char command[] = "CONST";
    while (1) {
        scanf("%s", command);
        if (strcmp(command, "CONST") == 0) {
            scanf("%d", &x);
            StackPush(&stack, x);
        } else if (strcmp(command, "ADD") == 0) {
            x = StackPop(&stack);
            y = StackPop(&stack);
            StackPush(&stack, x + y);
        } else if (strcmp(command, "SUB") == 0) {
            x = StackPop(&stack);
            y = StackPop(&stack);
            StackPush(&stack, x - y);
        } else if (strcmp(command, "MUL") == 0) {
            x = StackPop(&stack);
            y = StackPop(&stack);
            StackPush(&stack, x * y);
        } else if (strcmp(command, "DIV") == 0) {
            x = StackPop(&stack);
            y = StackPop(&stack);
            StackPush(&stack, x / y);
        } else if (strcmp(command, "MAX") == 0) {
            x = StackPop(&stack);
            y = StackPop(&stack);
            StackPush(&stack, x > y ? x : y);
        } else if (strcmp(command, "MIN") == 0) {
            x = StackPop(&stack);
            y = StackPop(&stack);
            StackPush(&stack, x > y ? y : x);
        } else if (strcmp(command, "NEG") == 0) {
            x = StackPop(&stack);
            StackPush(&stack, -x);
        } else if (strcmp(command, "DUP") == 0) {
            x = StackPop(&stack);
            StackPush(&stack, x);
            StackPush(&stack, x);
        } else if (strcmp(command, "SWAP") == 0) {
            x = StackPop(&stack);
            y = StackPop(&stack);
            StackPush(&stack, x);
            StackPush(&stack, y);
        } else if (strcmp(command, "END") == 0) {
            StackPopAndTell(&stack);
            free(stack.data);
            return 0; // завершение работы программы
        }
    }
}
