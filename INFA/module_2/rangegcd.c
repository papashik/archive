// У Вас ошибки в условии - два раза вместо слова "наибольший" 
// используется "наименьший" - 3 и 4 абзац. Поправьте, пожалуйста.


#include <stdio.h>
#include <math.h>
#include <stdlib.h>

void compute_logarithms(int nel, int *arr_lg) {   
    int pow2_i = 2, j = 0, i = 1;
    while (i <= nel) {
        while (j < pow2_i && j < (nel * 2)) {
            arr_lg[j] = i - 1;
            j++;
        }
        i++;
        pow2_i *= 2;
    }
}

int gcd(int a, int b) {
    int new_a = a > 0 ? a : -a;
    int new_b = b > 0 ? b : -b;
    while (1) {
        if (new_a == 0) {
            return new_b;
        } else if (new_b == 0) {
            return new_a;
        } else if (new_a > new_b) {
            new_a = new_a % new_b;
        } else {
            new_b = new_b % new_a;
        }
    }
}


void sparse_table_build(int nel, int log_2_nel, int *arr, int *arr_lg, int **arr_sparse) {
    int i = 0;
    int h = arr[i];
    while (i < nel) {
        arr_sparse[0][i] = h;
        i++;
        h = arr[i];
    }
    int j = 1, pow2_j = 2;
    while (j < log_2_nel) {
        i = 0;
        while (i <= nel - pow2_j) {
            arr_sparse[j][i] = gcd(arr_sparse[j - 1][i + (pow2_j / 2)], arr_sparse[j - 1][i]);
            i++;
        }
        j++;
        pow2_j *= 2;
    }
}

int sparse_table_req(int nel, int log_2_nel, int l, int r, int *arr_lg, int **arr_sparse) {
    int j = arr_lg[r - l + 1];
    return gcd(arr_sparse[j][l], arr_sparse[j][r - (int)pow(2, j) + 1]);
}

int main(int argc, char **argv)
{
    
    int l = 0, r = 0, nel = 0, log_2_nel = 0, m = 0; 
    scanf("%d", &nel);
    int arr_lg[2 * nel];
    
    compute_logarithms(nel, arr_lg);
    
    
    log_2_nel = arr_lg[nel] + 1;
    int arr[nel]; // arr_sparse[nel][log_2_nel];
    /*
    int **arr_sparse = (int**)malloc(nel * log_2_nel * sizeof(int));
    
    
    */
    int **arr_sparse = (int**)calloc(log_2_nel, nel * sizeof(int));
    
    for (int i = 0; i < log_2_nel; i++) {
        arr_sparse[i] = (int*)calloc(nel, sizeof(int));
    }
    
    for (int i = 0; i < nel; i++) {
        scanf("%d", &arr[i]);
    }
    
    sparse_table_build(nel, log_2_nel, arr, arr_lg, arr_sparse);
    
    scanf("%d", &m);
    for (int i = 0; i < m; i++) {
        scanf("%d %d", &l, &r);
        
        printf("%d\n", sparse_table_req(nel, log_2_nel, l, r, arr_lg, arr_sparse));
        
    }
    for (int i = 0; i < log_2_nel; i++) {
        free(arr_sparse[i]);
    }
    free(arr_sparse);
    return 0;
}
