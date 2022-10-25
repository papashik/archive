/******************************************************************************

                            Online C Compiler.
                Code, Compile, Run and Debug C program online.
Write your code in this editor and press "Run" button to compile and execute it.

*******************************************************************************/

#include <stdio.h>
#include <math.h>
int fyb(int n) {
    if (n == 0 || n == 1) {return n;}
    else {
        return fyb(n - 1) + fyb(n - 2);
    }
}

int fib_raz(long long int n) { //будет раскладывать число в систему счисления Фибоначчи рекурсивно
    
}
long long int fibs[n];
int main() {
    int n = 93;
    
    long long int *p;
    p = fibs;
    *p = 0;
    *(p + 1) = 1;
    for (p = fibs + 2; p < fibs + n; p++) {
        *p = *(p - 1) + *(p - 2);
        printf("%lld\n", *p);
    }
    printf("%lld", fibs[1]);
    //printf("%lld", (long long int)pow(2, 63) - 1);
    
    return 0;
}
