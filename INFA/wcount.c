#include <stdio.h>
#include <stdlib.h>

int wcount(char *s) {
    int counter = 0;
    char in_word = 0;
    for (int i = 0; i < 1000; i++ ) {
        if (s[i] == '\n' || s[i] == '\0') { 
            //printf("\nfound\n");
            return (int)in_word == 1 ? counter + 1 : counter;
        } else if (s[i] != ' ' && s[i] != '\t') {
            in_word = 1; 
        } else if ((int)in_word == 1) { 
            in_word = 0;
            counter++;
        }
    }
}

int main()
{

    char str[1000] = "   qwer d  er   ";
    //fgets(str, 1000, stdin);
    printf("%d", wcount(str));

    return 0;
}
