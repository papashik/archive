#include <stdio.h>
#include <stdlib.h>
#include <string.h>
// работа с очередью

typedef struct struct_Queue {
    int *data, cap, count, head, tail;
} queue_type;

char Enqueue(queue_type *queue, int x);
int Dequeue(queue_type *queue);


void QueueInit(queue_type *queue, int n) {
    queue->cap = n;
    queue->count = 0;
    queue->head = 0;
    queue->tail = 0;
    queue->data = (int*)calloc(n, sizeof(int));
}

void QueueResize(queue_type *queue, int new_n) {
    int count = queue->count;
    int temp_arr[count];
    for (int i = 0; i < count; i++) {
        temp_arr[i] = Dequeue(queue);
    }
    queue->data = (int*)realloc(queue->data, new_n * sizeof(int));
    queue->cap = new_n;
    queue->count = 0;
    queue->head = 0;
    queue->tail = 0;
    for (int i = 0; i < count; i++) {
        Enqueue(queue, temp_arr[i]);
    }
}


char QueueEmpty(queue_type *queue) {
    return queue->count == 0 ? 1 : 0;
}

char Enqueue(queue_type *queue, int x) {
    if (queue->count == queue->cap) {
        //printf("Queue overflow!\n");
        return -1; // means overflow
    }
    queue->data[queue->tail] = x;
    queue->tail++;
    if (queue->tail == queue->cap) { queue->tail = 0; }
    queue->count++;
    return 0; // normal exit code
}

int Dequeue(queue_type *queue) {
    if (QueueEmpty(queue) == 1) {
        printf("Queue is empty!\n");
        return -1;
    } else {
        int x = queue->data[queue->head];
        queue->head++;
        if (queue->cap == queue->head) { queue->head = 0; }
        queue->count--;
        return x;
    }
}

void QueuePopAndTell(queue_type *queue) {
    int x = Dequeue(queue);
    printf("%d\n", x);
}

int main()
{
    queue_type queue;
    int buff = 4, x = 0;
    QueueInit(&queue, buff);
    buff *= 2;
    
    char command[] = "EMPTY";
    while (1) {
        scanf("%s", command);
        if (strcmp(command, "ENQ") == 0) {
            scanf("%d", &x);
            if (Enqueue(&queue, x) == -1) {
                QueueResize(&queue, buff);
                buff *= 2;
                Enqueue(&queue, x);
            }
        } else if (strcmp(command, "DEQ") == 0) {
            QueuePopAndTell(&queue);
        } else if (strcmp(command, "EMPTY") == 0) {
            x = QueueEmpty(&queue);
            if (x == 1) {
                printf("true\n");
            } else {
                printf("false\n");
            }
        } else if (strcmp(command, "END") == 0) {
            free(queue.data);
            return 0; // завершение работы программы
        }
    }
    
    
}
