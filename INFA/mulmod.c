
#include <stdio.h>
#include <math.h>
int main()
{
    unsigned long long int a, b, m, res, r;
    scanf("%lld %lld %lld", &a, &b, &m);
    
    int b_binary[64] = {0};
    
    for (int i = 0; i < 64; i++) {
        r = (unsigned long long int)pow(2, i) & b;
        if (r == 0) {
            b_binary[i] = 0;
        } else {
            b_binary[i] = 1;
        }
    }
    
    res = a * b_binary[63];
    
    for (int bit_num = 62; bit_num >= 0; bit_num--) {
        res = ((res * 2) % m + a * b_binary[bit_num]) % m;
    }
    
    //res = (res + a * b) % m;
    printf("%lld", res);

    return 0;
}
