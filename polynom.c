#include <stdio.h>

int main()
{
    int n;
    long long int x0, polynom;
    scanf("%d %lld", &n, &x0);
    long long int a[n];
    for (int i = 0; i <= n; i++) {
        scanf("%lld", &a[i]);
    }
    
    polynom = a[0]; // sam polynom v tochke x0
    for (int i = 1; i < n + 1; i++) {
        polynom = (polynom * x0) + a[i];
    }
    printf("%lld\n", polynom);
    
    long long int proizv = a[0] * n; // proizvodnaya v tochke x0
    for (int i = 1, mnozhitel = n - 1; i < n; i++, mnozhitel--) {
        proizv = (proizv * x0) + a[i] * mnozhitel;
    }
    printf("%lld\n", proizv);
    
    return 0;
}
