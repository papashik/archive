#include <stdio.h>
#include <stdlib.h>
void revarray(void *base, size_t nel, size_t width) {
    char *base_char = base; // указатель на char
    // и присваиваем этому указателю значение base
    for (size_t i = 0; i < nel/2; i++) {
        for (size_t char_num = 0; char_num < width; char_num++) {
                char *from_address = base_char + char_num + i*width;
                char *to_address = base_char + char_num + (nel - 1 - i) * width;
                char temp_char = *from_address;
                *from_address = *to_address;
                *to_address = temp_char;
        }
    }
}
