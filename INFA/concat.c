#include <stdio.h>
#include <string.h>
#include <stdlib.h>

char* concat (char **s, int n) {
    int len0, len1, common_len;
    for (int i = 0; i < n; i++) {
        common_len += strlen(s[i]);
    }
    len0 = strlen(s[0]); // длина первой строки
    //char dest[common_len];
    char *dest = (char *) malloc(common_len * sizeof(char));
    strncpy(dest, s[0], len0);
    //strncpy(dest, s[0], len0 + 1);
    dest[len0 + 1] = 0;
    
    for (int i = 1; i < n; i++) {
        len1 = strlen(s[i]);
        //dest[len0] = 0;
        strncpy(dest + len0, s[i], len1 + 1);
        //dest[len0 + len1] = 0;
        len0 += len1;
    }
    return dest;
}

int main()
{
    //printf("---");
   /* char str0[] = "hello1world";
    char str1[] = "abcd";
    char str2[] = "efgh";
    char str3[] = "123456";
    char *strs[] = {str0, str1, str2, str3};
    char * a = concat(strs, 4); */
    
    int n;
    // printf("%d", getchar());
    scanf("%d ", &n);


    char* strs[n];
    for (int i = 0; i < n; i++) {
        *(strs + i) = malloc(1000 * sizeof(char));
    }
    strs[0] = "asdasd";
    int counter = 0;
    int number = 0;
    char string[1000] = {0}; 
    while (counter != n) {
        char another_char = getchar();
        string[number] = another_char;
        number++;
        if (another_char == 10) {
            strs[counter] = string;
            for (int i = 0; i < 1000; i++) { string[i] = 0; }
            counter++;
            number = 0;
            
            continue;
        }
    }
    
    for(int i = 0; i < n; i++) {
        printf("%s", strs[i]);
    }
    
    for (int i = 0; i < n; i++) {
        free(strs[i]);
    }
    
    /*for (int i = 0; i < n; i++) {
        char string[1000];
        fgets(string, 1000, stdin);
        strs[i] = string;
        printf("%s", strs[i]);
    }
    
    
    char * a = concat(strs, 4);
    printf("%s", a);
    */
    
    /*
    char *strs[n];
    for (int i = 0; i < n; i++) {
        scanf("%s", strs[i]);
    }
    
    for (int i = 0; i < n; i++) {
        printf("%s", strs[i]);
    }*/
    
    
    /* char *strs = (char *) malloc(n * 1000 * sizeof(char));
    for (int i = 0; i < n; i++) {

        fgets(strs + 1000 * i, 1000, stdin);
        fflush(stdin); // очищаем поток ввода
    } 
    for (int j = 0; j < n; j++) {
        char string = strs[j];
        printf("%c", string);
    }
    free(strs); */
    return 0;
}
