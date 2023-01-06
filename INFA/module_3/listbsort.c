#include <stdio.h>
#include <string.h>
#include <stdlib.h>

typedef struct Elem {
    struct Elem *next;
    char *word;
} type_elem;

int compare(type_elem *elem1, type_elem *elem2) {
    return strlen(elem1->word) > strlen(elem2->word) ? 1 : 0;
}

void swap(type_elem *prev_elem, type_elem *elem1, type_elem *elem2) {
    // "prev_elem->elem1->elem2" меняется на "prev_elem->elem2->elem1"
    //printf("-DO-\n%s->%s\n%s->%s\n%s->%s\n----\n", prev_elem->word, prev_elem->next->word, elem1->word, elem1->next->word, elem2->word, elem2->next->word);
    type_elem *temp = elem2->next;
    prev_elem->next = elem2;
    elem2->next = elem1;
    elem1->next = temp;
    //printf("%s->%s\n%s->%s\n%s->%s\n", prev_elem->word, prev_elem->next->word, elem2->word, elem2->next->word, elem1->word, elem1->next->word);
}


struct Elem *bsort(struct Elem *list) {
    if (list == NULL) { return NULL; }
    if (list->next == NULL) { return list; }
    char was_swap = 0;
    type_elem *prev_prev_elem, *prev_elem, *now_elem, *null_elem = list;
    prev_prev_elem = null_elem;
    prev_elem = null_elem->next;
    now_elem = null_elem->next->next;

    while (1) {
        while (now_elem != NULL) {
            // printf("%s %s %s\n", prev_prev_elem->word, prev_elem->word, now_elem->word);
            if (compare(prev_elem, now_elem) > 0) {
                swap(prev_prev_elem, prev_elem, now_elem);
                was_swap = 1;
                prev_prev_elem = now_elem;
                now_elem = prev_elem->next;
            } else {
                prev_prev_elem = prev_prev_elem->next;
                prev_elem = prev_elem->next;
                now_elem = now_elem->next;
            }
        }
        if (was_swap == 0) {
            break;
        } else {
            was_swap = 0;
            prev_prev_elem = null_elem;
            prev_elem = null_elem->next;
            now_elem = null_elem->next->next;
        }
    }
    if (compare(null_elem, null_elem->next) > 0) {
        type_elem *temp = null_elem->next;
        null_elem->next = null_elem->next->next;
        temp->next = null_elem;
        null_elem = bsort(temp);
    }
    return null_elem;
    
}


int main()
{
    char** arr_words = (char**)malloc(1000 * sizeof(char*));
    for (int i = 0; i < 1000; i++) {
        arr_words[i] = (char*)malloc(1001 * sizeof(char));
    }
    
    int ind = 0;
    do {
        scanf("%s", arr_words[ind]);
        //printf("%s ", arr_words[ind]);
        ind++;
    } while (getchar() != 10);
    
    if (ind == 1 || ind == 0) {
        if (ind == 1) { printf("%s", arr_words[0]); }
        for (int i = 0; i < 1000; i++) {
            free(arr_words[i]);
        }
        free(arr_words);
        return 0;
    }
    
    type_elem arr[ind];
    type_elem *w = arr; // указатель на первый элемент
    
    for (int i = 0; i < ind - 1; i++) {
        arr[i].word = (char*)malloc(1001 * sizeof(char));
        arr[i].next = &arr[i + 1];
        strncpy(arr[i].word, arr_words[i], 1001); 
    }
    
    arr[ind - 1].word = (char*)malloc(1001 * sizeof(char));
    arr[ind - 1].next = NULL;
    strncpy(arr[ind - 1].word, arr_words[ind - 1], 1001);
    
    for (int i = 0; i < 1000; i++) {
        free(arr_words[i]);
    }
    free(arr_words);
    
    w = bsort(w);
    
    for (int i = 0; i < ind - 1; i++) {
        printf("%s ", w->word);
        w = w->next;
    }
    printf("%s ", w->word);
    
    for (int i = 0; i < ind; i++) {
        free(arr[i].word);
    }
    
    return 0;
}
