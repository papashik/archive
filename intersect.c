/******************************************************************************

                            Online C Compiler.
                Code, Compile, Run and Debug C program online.
Write your code in this editor and press "Run" button to compile and execute it.

*******************************************************************************/

#include <stdio.h>
#include <math.h>
int main()
{
    int a_len, b_len, temp;
    
    scanf("%d", &a_len);
    int a = 0;
    for (int i = 0; i < a_len; i++) {
        scanf("%d", &temp);
        a += pow(2, temp);
    }
    //printf("%d\n", a);
    //-------------------------------//
    scanf("%d", &b_len);
    int b = 0;
    for (int i = 0; i < b_len; i++) {
        scanf("%d", &temp);
        b += pow(2, temp);
    }
    //printf("%d\n", b);
    
    int intersection = a & b;
    
    for (int i = 0; i < 32; i++) {
        temp = ((int)pow(2, i)) & intersection;
        if (temp) { printf("%d ", i);}
    }
    
    return 0;
}
