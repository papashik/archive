#include <stdio.h>
typedef struct Date {
    int Day, Month, Year;
} Date;

void date_sort(Date *arr, int nel) {
    int days_count[31] = {0};
    int months_count[12] = {0};
    int years_count[61] = {0}; // 1970 - 2030
    
    int days_indexes[31] = {0};
    int months_indexes[12] = {0};
    int years_indexes[61] = {0};
    
    // -------- amount -------------
    for (int i = 0; i < nel; i++) {
        days_count[arr[i].Day - 1] += 1;
        months_count[arr[i].Month - 1] += 1;
        years_count[arr[i].Year - 1970] += 1;
    }
    
    // -------- indexes ------------
    for (int i = 1; i < 31; i++) {
        days_indexes[i] = days_indexes[i - 1] + days_count[i - 1];
    }
    for (int i = 1; i < 12; i++) {
        months_indexes[i] = months_indexes[i - 1] + months_count[i - 1];
    }
    for (int i = 1; i < 61; i++) {
        years_indexes[i] = years_indexes[i - 1] + years_count[i - 1];
    }
    
    
    Date temp_arr[nel];
    int index = 0, day = 0, month = 0, year = 0;
    
    // --------------- days ---------
    for (int i = 0; i < nel; i++) {
        day = arr[i].Day - 1;
        index = days_indexes[day]; // индекс, по которому будем заносить эту запись в новый массив
        temp_arr[index] = arr[i];
        days_indexes[day] += 1;
    }
    
    for (int i = 0; i < nel; i++) {
        arr[i] = temp_arr[i];
    }
    // -------------- months ---------------
    for (int i = 0; i < nel; i++) {
        month = arr[i].Month - 1;
        index = months_indexes[month];
        temp_arr[index] = arr[i];
        months_indexes[month] += 1;
    }
    
    for (int i = 0; i < nel; i++) {
        arr[i] = temp_arr[i];
    }
    // -------------- years ---------------
    for (int i = 0; i < nel; i++) {
        year = arr[i].Year - 1970;
        index = years_indexes[year];
        temp_arr[index] = arr[i];
        years_indexes[year] += 1;
    }
    
    for (int i = 0; i < nel; i++) {
        arr[i] = temp_arr[i];
    }
}

int main()
{
    int n = 0;
    scanf("%d", &n);
    Date arr[n];
    for (int i = 0; i < n; i++) {
        scanf("%d %d %d", &arr[i].Year, &arr[i].Month, &arr[i].Day);
    }
    
    date_sort(arr, n);
    
    for (int i = 0; i < n; i++) {
        printf("%04d %02d %02d\n", arr[i].Year, arr[i].Month, arr[i].Day); 
    }
    return 0;
}
