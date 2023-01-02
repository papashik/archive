#include <stdio.h>

typedef union Int32 { 
    int x; 
    unsigned char bytes[4]; 
} un_int32;

void distribution_sort(int byte_num, int num_base, un_int32 *arr, int nel) {
    int count[num_base]; // Количество записей по каждому ключу
    for (int i = 0; i < num_base; i++) {
        count[i] = 0;
    }
    for (int j = 0; j < nel; j++) {
        int k = arr[j].bytes[byte_num]; // ключ - байт записи по номеру
        count[k] = count[k] + 1;
    }
    
    for (int i = 1; i < num_base; i++) {
        count[i] = count[i] + count[i - 1]; // теперь в count лежат последние индексы для каждого ключа
    }
    
    un_int32 arr_new[nel];
    
    for (int j = nel - 1; j >= 0; j--) {
        int k = arr[j].bytes[byte_num]; // ключ - байт записи по номеру
        int ind = count[k] - 1;
        count[k] = ind;
        arr_new[ind].x = arr[j].x;
    }
    
    for (int i = 0; i < nel; i++) {
        arr[i] = arr_new[i];
    }
}

int main()
{
    int n;
    scanf("%d", &n);
    un_int32 arr[n];
    char is_less_null = 0; // есть ли числа меньше нуля
    for (int i = 0; i < n; i++) {
        scanf("%d", &arr[i].x);
        if (arr[i].x < 0) { is_less_null = 1; }
    }

    for (int byte_num = 0; byte_num < 4; byte_num++) {
        distribution_sort(byte_num, 256, arr, n);
    }
    
    if (is_less_null == 1) {
        un_int32 arr_new[n]; // здесь расставим в нормальном порядке знаковые и беззнаковые числа
        int ind = 0;
        for (int i = 0; i < n; i++) {
            if (arr[i].x < 0) {
                arr_new[ind] = arr[i];
                ind++;
            }
        }
    
        for (int i = ind; i < n; i++) {
            arr_new[i] = arr[i - ind];
        }
        
        for (int i = 0; i < n; i++) {
            printf("%d ", arr_new[i].x);
        }
    } else {
        for (int i = 0; i < n; i++) {
            printf("%d ", arr[i].x);
        }
    }
    return 0;
}
