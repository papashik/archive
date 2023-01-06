#include <stdio.h>
#include <stdlib.h>
#include <limits.h>
#include <string.h>

// работа с двойным стеком

typedef struct struct_Stack {
    int cap, top1, top2;
    int *data, *maximum;
} stack_type;

void DoubleStackInit(stack_type *stack, int n) {
    stack->cap = n;
    stack->top1 = 0;
    stack->top2 = n - 1;
    stack->maximum = (int*)calloc(n, sizeof(int));
    stack->data = (int*)calloc(n, sizeof(int));
}

char StackEmpty1(stack_type *stack) {
    return stack->top1 == 0 ? 1 : 0;
}
char StackEmpty2(stack_type *stack) {
    return stack->top2 == stack->cap - 1 ? 1 : 0;
}

int StackReceiveMax1(stack_type *stack) {
    if (StackEmpty1(stack) == 1) {
        return INT_MIN;
    } else {
        return stack->maximum[stack->top1 - 1];
    }
}
int StackReceiveMax2(stack_type *stack) {
    if (StackEmpty2(stack) == 1) {
        return INT_MIN;
    } else {
        return stack->maximum[stack->top2 + 1];
    }
}

int StackReceiveMax(stack_type *stack) {
    int max1 = StackReceiveMax1(stack);
    int max2 = StackReceiveMax2(stack);
    return max1 > max2 ? max1 : max2;
}

char StackPush1(stack_type *stack, int x) {
    if (stack->top2 < stack->top1) {
        // printf("Stack overflow!\n");
        return -1;
    } else {
        stack->data[stack->top1] = x;
        int stack1_max = StackReceiveMax1(stack);
        stack->maximum[stack->top1] = x > stack1_max ? x : stack1_max;
        stack->top1++;
        return 0;
    }
}
char StackPush2(stack_type *stack, int x) {
    if (stack->top2 < stack->top1) {
        // printf("Stack overflow!\n");
        return -1;
    } else {
        stack->data[stack->top2] = x;
        int stack2_max = StackReceiveMax2(stack);
        stack->maximum[stack->top2] = x > stack2_max ? x : stack2_max;
        stack->top2--;
        return 0;
    }
}

int StackPop1(stack_type *stack) {
    if (StackEmpty1(stack) == 1) {
        // printf("Stack 1 is empty!\n");
        return INT_MIN;
    } else {
        stack->top1--;
        return stack->data[stack->top1];
    }
}
int StackPop2(stack_type *stack) {
    if (StackEmpty2(stack) == 1) {
        // printf("Stack 2 is empty!\n");
        return INT_MIN;
    } else {
        stack->top2++;
        return stack->data[stack->top2];
    }
}

void StackPopAndTell1(stack_type *stack) {
    printf("%d\n", StackPop1(stack));
}
void StackPopAndTell2(stack_type *stack) {
    printf("%d\n", StackPop2(stack));
}

// работа с очередью на двойном стеке

void QueueOnStackInit(stack_type *stack, int n) {
    DoubleStackInit(stack, n);
}

char QueueEmpty(stack_type *stack) {
    return StackEmpty1(stack)&StackEmpty2(stack);
}

char Enqueue(stack_type *stack, int x) {
    return StackPush1(stack, x); // 0 - normal, -1 - error
}

int Dequeue(stack_type *stack) {
    if (StackEmpty2(stack) == 1) {
        while (! (StackEmpty1(stack) == 1)) {
            StackPush2(stack, StackPop1(stack));
        }
    }
    return StackPop2(stack);
}

int QueueReceiveMax(stack_type *stack) {
    return StackReceiveMax(stack);
}

void QueuePopAndTell(stack_type *stack) {
    printf("%d\n", Dequeue(stack));
}

int main()
{
    stack_type queue;
    DoubleStackInit(&queue, 100000);
    int x = 0;
    char command[] = "EMPTY";
    while (1) {
        scanf("%s", command);
        if (strcmp(command, "ENQ") == 0) {
            scanf("%d", &x);
            Enqueue(&queue, x);
        } else if (strcmp(command, "DEQ") == 0) {
            QueuePopAndTell(&queue);
        } else if (strcmp(command, "EMPTY") == 0) {
            if (QueueEmpty(&queue) == 1) {
                printf("true\n");
            } else {
                printf("false\n");
            }
        } else if (strcmp(command, "MAX") == 0) {
            printf("%d\n", QueueReceiveMax(&queue));
        } else if (strcmp(command, "END") == 0) {
            free(queue.data);
            free(queue.maximum);
            return 0; // завершение работы программы
        }
    }
}
