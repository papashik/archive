#include <stdio.h>
#include <stdlib.h>

// работа со стеком
typedef struct Task { 
    int low, high; 
} task_type;

typedef struct struct_Stack {
    int cap, top;
    task_type *data;
} stack_type;

void StackInit(stack_type *stack, int n) {
    stack->cap = n;
    stack->top = 0;
    stack->data = (task_type*)calloc(n, sizeof(task_type));
}

char StackEmpty(stack_type *stack) {
    return stack->top == 0 ? 1 : 0;
}

void StackPush(stack_type *stack, task_type *task) {
    if (stack->top == stack->cap) {
        printf("Stack overflow!\n");
    } else {
        stack->data[stack->top] = *task;
        stack->top++;
    }
}

task_type StackPop(stack_type *stack) {
    if (StackEmpty(stack) == 1) {
        printf("Stack is empty!\n");
    } else {
        stack->top--;
        return stack->data[stack->top];
    }
}

void StackPopAndTell(stack_type *stack) {
    task_type x = StackPop(stack);
    printf("low = %d, high = %d\n", x.low, x.high);
}

// работа с массивом
void swap(int *arr, int i, int j) {
    int temp = arr[i];
    arr[i] = arr[j];
    arr[j] = temp;
}

int compare(int a, int b) {
    return a - b;
}

int partition(int *arr, int l, int r) {
    int ind = l; // указывает на место, левее которого элементы меньше опорного 
    for (int j = l; j < r; j++) {
        if (compare(arr[r], arr[j]) > 0) {
            swap(arr, ind, j);
            ind++;
        }
    }
    swap(arr, r, ind);
    return ind;
}

int main()
{
    stack_type stack;
    StackInit(&stack, 10000);
    
    int n;
    scanf("%d", &n);
    int arr[n];
    for (int i = 0; i < n; i++) {
        scanf("%d", &arr[i]);
    }
    
    task_type null_task = {0, n - 1};
    StackPush(&stack, &null_task);
    
    while (! StackEmpty(&stack) ) {
        task_type task = StackPop(&stack);
        int l = task.low, r = task.high;
        if (l < r) {
            // алгорит разделения относительно последнего - partition
            int q = partition(arr, l, r);
            task_type task_l = {l, q - 1};
            task_type task_r = {q + 1, r};
            StackPush(&stack, &task_l);
            StackPush(&stack, &task_r);
        } else {
            continue;
        }
    }
    
    for (int i = 0; i < n; i++) {
        printf("%d ", arr[i]);
    }
    
    free(stack.data);
    return 0;
}
