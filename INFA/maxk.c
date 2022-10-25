#include <stdio.h>

int main()
{
    /* ---- That would work if we receive k earlier than array ----
    
    int n, k, summax, sumk = 0;
    scanf("%d %d", &n, &k);
    int arr[n], *p;
    for (p = arr; p < arr + k; p++) {
        scanf("%d", p);
        sumk += *p;
    }
    summax = sumk;
    for (; p < arr + n; p++) { // p = arr + k
        scanf("%d", p);
        sumk = sumk + *p - *(p-k);
        if (sumk > summax) {
            summax = sumk;
        }
    }
    
    printf("%d\n", summax); 
    
    ---- But we have what we have and deserve what we deserve ----*/
    
    int n, k, summax, sumk = 0;
    scanf("%d", &n);
    int arr[n], *p;
    
    for (p = arr; p < arr + n; p++) {
        scanf("%d", p);
    }
    
    printf("Enter k:");
    scanf("%d", &k);
    for (p = arr; p < arr + k; p++) {
        sumk += *p;
    }
    summax = sumk;
    
    for (; p < arr + n; p++) { // p = arr + k
        sumk = sumk + *p - *(p-k);
        if (sumk > summax) {
            summax = sumk;
        }
    }
    
    printf("%d\n", summax); 
    
    
    /*for (i = 0; i < n; i++) {
        printf("%d\n", arr[i]);
    }*/

    return 0;
}
