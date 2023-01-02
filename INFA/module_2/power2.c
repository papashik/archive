#include <stdio.h>
#include <math.h>
int main()
{
    int n;
    scanf("%d", &n);
    int arr[n];
    long long int s = 0, temp = 0, order = 0, res = 0, max_order = (long long int)pow(2, n);
    for (int i = 0; i < n; i++) {
        scanf("%d", &arr[i]);
    }
    
    for (order = 1; order < max_order; order++) {
        s = 0;
        if ((1&order) != 0) {
            s += arr[0];
        }
        temp = 1;
        for (int i = 1; i < n; i++) {
            temp *= 2;
            if ((temp&order) != 0) {
                s += arr[i];
            }
        }
        if ((s&(s-1)) == 0 && s > 0) {
            res++; 
        }   
    }
    printf("%lld", res);
    
    return 0;
}
