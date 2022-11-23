#include <stdio.h>

char* check(int n) {
    int arr1[n], arr2[n];
    for (int i = 0; i < n; i++) {
        scanf("%d", &arr1[i]);
    }
    for (int i = 0; i < n; i++) {
        scanf("%d", &arr2[i]);
    }
    
    for (int i = 0; i < n; i++) {
        for (int finder = arr1[i], j = 0; j < n; j++) {
            if (finder == arr2[j]) {
                // printf("%d == %d\n", finder, arr2[j]);
                break;
            }
            if (j == n - 1) {
                return "no";
            }
        }
    }
    return "yes";
}
int main()
{
    int n = 8;
    printf("%s", check(n));
}
