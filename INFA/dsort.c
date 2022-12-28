#include <stdio.h>
#include <string.h>
int main()
{
    int arr[26] = {0};
    char stroka[1000001];
    char res_stroka[1000001];
    scanf("%s", stroka);
    
    int len_stroka = strlen(stroka);
    for (int i = 0; i < len_stroka; i++) {
        arr[stroka[i] - 97] += 1;
    }
    
    int min_ind = 0, max_ind = 0;
    for (int i = 0; i < 26; i++) {
        min_ind = max_ind;
        max_ind += arr[i];
        for (int j = min_ind; j < max_ind; j++) {
            printf("%c", i + 97);
            //res_stroka[j] = i + 97;
        }
    }

    //printf("%s", res_stroka);
    return 0;
}
