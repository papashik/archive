#include <stdio.h>
#include <stddef.h>
#include <math.h>
int main()
{
    int n = 0, res = 0;
    scanf("%d", &n);
    int arr[n], s = 0, temp = 0, order = 0;
    for (int i = 0; i < n; i++) {
        scanf("%d", &arr[i]);
    }
    
    for (order = 1; order < (int)pow(2, n); order++) {
        s = 0;
        for (int i = 0; i < n; i++) {
            temp = (int)pow(2, i);
            if ((temp&order) != 0) {
                s += arr[i];
            }
        }
        //s = s > 0 ? s : -s;
        if ((s&(s-1)) == 0 && s > 0) {
            res++; 
            //printf("%ld\n", s);
            //printf("%ld\n", order);
        }
        
        
    }
   
    
    printf("%d", res);
    
    return 0;
}
