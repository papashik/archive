#include <stdio.h>

void insertsort(int *arr, int nel, int (*compare)(int i, int j));

void merge_sort(int i, int j, int* arr, int* res, int (*compare)(int i, int j)) {
    
    if (j - i < 4) {
        insertsort(arr + i, j - i + 1, compare);
        for (int k = i; k <= j; k++) {
            res[k] = arr[k];
        }
        return;
    }
    
    int mid = (i + j) / 2;
    
    merge_sort(i, mid, arr, res, compare);     
    merge_sort(mid + 1, j, arr, res, compare);

    int left = i;       
    int right = mid + 1;        
    int k;      

    for (k = i; k <= j; k++) {
        if (left == mid + 1) {
            res[k] = arr[right];
            right++;
        } else if (right == j + 1) {
            res[k] = arr[left];
            left++;
        } else if (compare(arr[left], arr[right]) > 0) {
            res[k] = arr[right];
            right++;
        } else {
            res[k] = arr[left];
            left++;
        }
    }

    for (k = i; k <= j; k++) {
        arr[k] = res[k];
    }
}
/*
void insertsort(int *arr, int nel, int (*compare)(int i, int j)) {
    int temp;
    if (nel > 1) {
        for (int i = 1; i < nel; i++) {
            int x = arr[i];
            int j = i;
            while (j > 0 && compare(arr[j-1], x) > 0) {
                arr[j] = arr[j-1];
                j--;
            }
            arr[j] = x;
        }
    }
}
*/
void insertsort(int *arr, int nel, int (*compare)(int i, int j)) {
    int temp;
    if (nel > 1) {
        for (int i = 1; i < nel; i++) {
            int j = i;
            while (j > 0 && compare(arr[j-1], arr[j]) > 0) {
                temp = arr[j];
                arr[j] = arr[j-1];
                arr[j-1] = temp;
                j--;
            }
        }
    }
}

int compare(int i, int j) {
    int abs_i = i > 0 ? i : -i;
    int abs_j = j > 0 ? j : -j;
    if (abs_i > abs_j) {
        return 1;
    } else {
        return -1;
    }
}

int main()
{
    int n = 0, temp = 0;
    scanf("%d", &n);
    int arr[n], arr_res[n];
    for (int i = 0; i < n; i++) {
        scanf("%d", &arr[i]);
    }
    merge_sort(0, n - 1, arr, arr_res, *compare);

    for (int i = 0; i < n; i++) {
        printf ("%d ", arr_res[i]);
    }
    
}
