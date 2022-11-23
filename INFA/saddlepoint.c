#include <stdio.h>
#include <math.h>

int main()
{
    int m, n;
    scanf("%d %d", &m, &n);
    int arr[m][n], temp;
    int arr_max_rows[m], arr_min_cols[n];
    
    for (int i = 0; i < fmax(m, n); i++) {
        arr_max_rows[i] = 0;
        arr_min_cols[i] = 2147483648;
    }
    for (int i = 0; i < m; i++) {
        for (int j = 0; j < n; j++) {
            scanf("%d", &temp);
            arr[i][j] = temp;
            if (temp > arr_max_rows[i]) {
                arr_max_rows[i] = temp;
            }
            if (temp < arr_min_cols[j]) {
                arr_min_cols[j] = temp;
            }
        }
    }
    /* printf("\n--\n");
    for (int i = 0; i < m; i++) {
        printf("%d\n", arr_max_rows[i]);
    }
    printf("\n--\n");
    for (int i = 0; i < n; i++) {
        printf("%d\n", arr_min_cols[i]);
    }
    printf("\n--\n"); */
    for (int i = 0; i < fmin(m, n); i++) {
        int finder = arr_max_rows[i];
        for (int j = 0; j < fmin(m, n); j++) {
            if (finder == arr_min_cols[j]) {
                printf("%d %d", i, j);
                return 0;
            }
        }
    }
    printf("none");
    return 0;
}
