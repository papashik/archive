#include <stdio.h>
#include <stdlib.h>

typedef struct Elem { 
    struct Elem *prev, *next;
    int v; 
} type_elem;

void DoubleLinkedListInit(type_elem *arr) {
    arr->prev = arr;
    arr->next = arr;
}

char ListEmpty(type_elem *arr) {
    return arr->next == arr ? 1 : 0;
}

void InsertAfter(type_elem *elem, type_elem *new_elem) {
    type_elem *z = elem->next;
    elem->next = new_elem;
    new_elem->prev = elem;
    z->prev = new_elem;
    new_elem->next = z;
}

void InsertBefore(type_elem *elem, type_elem *new_elem) {
    type_elem *z = elem->prev;
    elem->prev = new_elem;
    new_elem->next = elem;
    z->next = new_elem;
    new_elem->prev = z;
}

type_elem *FindFirstWithVBiggerThan(int v, type_elem *arr) {
    type_elem *x = arr->next;
    while ((x != arr) && (v > x->v)) {
        x = x->next;
    }
    return x;
}

int main() {
    
    int n = 0, x = 0;
    scanf("%d", &n);
    type_elem arr[n + 1], new_elem;
    type_elem *insert_before_elem;
    DoubleLinkedListInit(arr);
    for (int i = 0; i < n; i++) {
        scanf("%d", &x);
        arr[i + 1].v = x;
        insert_before_elem = FindFirstWithVBiggerThan(x, arr);
        InsertBefore(insert_before_elem, &arr[i + 1]);
    }
    
    type_elem elem = *arr[0].next;
    for (int i = 0; i < n; i++) {
        printf("%d ", elem.v);
        elem = *elem.next;
    }
    
    return 0;
}
