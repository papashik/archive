#include <stdio.h>
#include <string.h>
#include <stdlib.h>
void prefix(char *s, int *func, int nel) {
    // func - массив со значениями префиксной функции
    for (int i = 0; i < nel; i++) {
        func[i] = 0;
    }
    int t = 0;
    for (int i = 1; i < nel; i++) {
        while (t > 0 && s[i] != s[t]) {
            t = func[t - 1];
        }
        if (s[i] == s[t]) {
            t++;
        }
        func[i] = t;
    }
}
int main(int argc, char **argv)
{
    char *s = argv[1];
    char *t = argv[2];
    int s_len = strlen(s);
    int t_len = strlen(t);
    int s_func[s_len];
    prefix(s, s_func, s_len);
    int q = 0;
    for (int k = 0; k < t_len; k++) {
        while ((q > 0) && (s[q] != t[k])) {
            q = s_func[q - 1];
        }
        if (s[q] == t[k]) {
            q++;
        }
        if (q == s_len) {
            printf("%d ", k - s_len + 1);
        }
    }
    
    return 0;
}
