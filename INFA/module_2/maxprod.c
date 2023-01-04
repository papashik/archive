#include <stdio.h>
#include <math.h>
#include <stdlib.h>

int main(int argc, char **argv)
{
    int n; 
    scanf("%d", &n);
    unsigned int arr_1[n], arr_2[n];
    for (int i = 0; i < n; i++) {
    scanf("%d/%d", &arr_1[i], &arr_2[i]);
    }
    int l = 0, r = 0, start = 0;
    long double prod = 1, maxprod = arr_1[0]/arr_2[0];
    for (int i = 0; i < n;) {
        prod = prod * arr_1[i] / arr_2[i];
        if (prod > maxprod) {
            maxprod = prod;
            l = start;
            r = i;
        }
        i++;
        if (prod < 1) {
            prod = 1;
            start = i;
        }
    }

    printf("%d %d", l, r);
    return 0;
}
