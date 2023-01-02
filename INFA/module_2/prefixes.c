#include <stdio.h>
#include <string.h>

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
    int nel = strlen(s);
    int func[nel], size = 0;
    prefix(s, func, nel);
    
    for (int i = 0; i < nel; i++) {
        size = i + 1;
        if ((func[i] > 0) && size % (size - func[i]) == 0) {
            printf("%d %d\n", size, size / (size - func[i]));
        }
    }
    
    
    return 0;
}
