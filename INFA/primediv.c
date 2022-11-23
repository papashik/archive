#include <stdio.h>

int main()
{
    int n = 46400;
    int primes[4800];
    primes[0] = 2;
    int prime_last_num = 1;
    for (int i = 3; i < n; i++) {
        int prime_num = 0;
        int prime = primes[0];
        while (1) {
            if (i % prime == 0) {
                break;
            }
            prime_num ++;
            if (prime_num == prime_last_num) {
                primes[prime_last_num] = i;
                prime_last_num++;
                break;
            }
            prime = primes[prime_num];
        }
    }
    printf("%d\n%d", prime_last_num, primes[prime_last_num - 1]);
    
    long int number;
    scanf("%ld", &number);
    for (int i = prime_last_num - 1; i > 0; i--) {
        if (number % primes[i] == 0) {
            printf("%d", primes[i]);
            break;
        }    
    }
    
    return 0;
}
