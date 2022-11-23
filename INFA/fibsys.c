#include <stdio.h>
#include <math.h>

int fyb(int n) {
    if (n == 0 || n == 1) {return n;}
    else {
        return fyb(n - 1) + fyb(n - 2);
    }
}

int fib_raz(long long int number, int n, long long int fibs[n]) { // будет раскладывать number в систему счисления Фибоначчи итеративно
   int result[n];
   for (int i = 0; i < n; i++) {
       result[i] = 0;
   }
   int flag = -1;
   for (int i = n - 1; i >= 0; i--) {
       if (fibs[i] <= number) {
           if (flag == -1) {flag = n - i - 1;}
           result[n - i - 1] = 1;
           number -= fibs[i];
           //printf("%lld\n", fibs[i]);
        }
   }
    for (int i = flag; i < n; i++) {
        printf("%d", result[i]);
    }
    //printf("\n----\n");
   return flag;

}   
int main() {
    int n = 91;
    long long int fibs[n];
    fibs[0] = 1;
    fibs[1] = 2;
    for (int i = 2; i < n ; i++) {
        fibs[i] = fibs[i - 1] + fibs[i - 2];
        //printf("%lld\n", fibs[i]);
    }
    long long int question;
    scanf("%lld", &question);
    fib_raz(question, n, fibs);
    
    //printf("%lld", (long long int)pow(2, 63) - 1);
    
    return 0;
}
