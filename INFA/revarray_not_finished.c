#include <stdio.h>
#include <stdlib.h>


void* revarray(void *base, size_t nel, size_t width) {
    size_t *memory = (size_t*) malloc(nel*width);
    for (size_t i = 0; i < nel; i++) {
        memory[(nel - i - 1)*width] = *((size_t*)base + i*width);
    }
    
    return memory;
}

int main()
{
    int f[] = {1, 2, 3, 4};
    int *p = f;
    int *p1 = (int*) revarray(p, 4, sizeof(int));
    for (int i = 0; i < 4; i++) {
        printf("%d\n", *(p1+i));
    }
    return 0;
}
