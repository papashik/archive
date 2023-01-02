#include <stdio.h>
#include <math.h>
#include <string.h>
#include <limits.h>
#include <stdlib.h>

const int inf = INT_MAX;

int find_max(int *arr, int n, int l, int r) {
    int ans = INT_MIN;
    l += n - 1; // n - половина размера массива
    r += n - 1;
    while (l <= r) {
        if ((l & 1) == 1) {
            ans = ans > arr[l] ? ans : arr[l];
        }
        if ((r & 1) == 0) {
            ans = ans > arr[r] ? ans : arr[r];
        }
        l = (l + 1) / 2;
        r = (r - 1) / 2;
    }
    return ans;
}

void tree_update(int *arr, int n, int ind, int v) {
    ind += n - 1;
    arr[ind] = v;
    while (ind > 1) {
        ind /= 2;
        arr[ind] = arr[ind * 2] > arr[ind * 2 + 1] ? arr[ind * 2] : arr[ind * 2 + 1];
    }
}


int main()
{
    int n;
    scanf("%d", &n);
    
	int new_n = (1 << (int)(log2(n - 1) + 1)); // размер, доведённый до степени двойки
    // int arr[new_n]; // массив с изначальными числами, добитый бесконечно малыми
    int *arr = (int*)malloc(new_n * 2 * sizeof(int));
    arr[0] = 0;
    for (int i = new_n; i < new_n + n; i++) {
        scanf("%d", &arr[i]);
    }
    // ------------- building -------------
    for (int i = new_n + n; i < 2 * new_n; i++) {
        arr[i] = INT_MIN;
    }
    for (int i = new_n - 1; i > 0; i--) {
        arr[i] = arr[2 * i] > arr[2 * i + 1] ? arr[2 * i] : arr[2 * i + 1];
    }
    // ------------- building -------------
    char s[] = "hey";
    int l = 0, r = 0, ind = 0, v = 0, maximum = 0;
    while (1) {
        scanf("%s", s);
        if (strcmp(s, "MAX") == 0) {
            scanf("%d %d", &l, &r);
            maximum = find_max(arr, new_n, l + 1, r + 1);
            printf("%d\n", maximum);
            
        } else if (strcmp(s, "UPD") == 0) {
            scanf("%d %d", &ind, &v);
            tree_update(arr, new_n, ind + 1, v);
            
        } else if (strcmp(s, "END") == 0) {
            break;
        }
    } 
    free(arr);
    return 0;
}
