#include <stdio.h>
void swap(int *arr, int i, int j) {
    int temp = arr[i];
    arr[i] = arr[j];
    arr[j] = temp;
}

int compare(int i, int j) {
    if (i > j) {
        return 1;
    } else {
        return -1;
    }
}

int partition(int *arr, int low, int high, int (*compare)(int i, int j), void (*swap)(int *arr, int i, int j)) {
    int i = low, j = low, temp = 0;
    while (j < high) {
        if (compare(arr[high], arr[j]) > 0) {
            swap(arr, i, j);
            i++;
        }
        j++;
    }
    swap(arr, high, i);
    return i;
}
void selection_sort(int *arr, int low, int high, int (*compare)(int i, int j), void (*swap)(int *arr, int i, int j)) {
    int min_ind = 0;
    for (int ind = low; ind < high; ind++) {
        min_ind = ind;
        for (int j = ind + 1; j <= high; j++) {
            if (compare(arr[min_ind], arr[j]) > 0) {
                min_ind = j;
            }
        }
        swap(arr, min_ind, ind);
    }
}
void quick_sort_rec(int *arr, int low, int high, int m, int (*compare)(int i, int j), void (*swap)(int *arr, int i, int j)) {
    int q = 0;
    if (high - low + 1 < m) {
        selection_sort(arr, low, high, compare, swap);
    } else {
        while (low < high) {
            q = partition(arr, low, high, compare, swap);
            if (q - low < high - q) {
                quick_sort_rec(arr, low, q, m, compare, swap);
                low = q + 1;
            } else {
                quick_sort_rec(arr, q, high, m, compare, swap);
                high = q - 1;
            }
            
            
        }
    }
}


int main()
{
    int n = 0, m = 0;
    scanf("%d", &n);
    scanf("%d", &m);
    int arr[n];
    for (int i = 0; i < n; i++) {
        scanf("%d", &arr[i]);
    }
    
    quick_sort_rec(arr, 0, n - 1, m, *compare, *swap);
    // selection_sort(arr, 0, n - 1, *compare, *swap);
    
    for (int i = 0; i < n; i++) {
        printf ("%d ", arr[i]);
    }
    return 0;
}
