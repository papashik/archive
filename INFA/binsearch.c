#include <stdio.h>

unsigned long binsearch(unsigned long nel, int (*compare)(unsigned long i)) {
    unsigned long low = 0;
    unsigned long high = nel - 1;
    unsigned long i;
    while (low <= high) {
        i = (low + high) / 2;
        if (compare(i) > 0) { // смещаемся вниз
            high = i - 1;
        } else if (compare(i) < 0) { // смещаемся вверх
            low = i + 1;
        } else { // нашли нужный
            return i;
        }
    }
    return nel;
}

int compare(unsigned long i) {
    int a = 10;
    if (i < a) {
        return -1;
    } else if (i > a) {
        return 1;
    } else {
        return 0;
    }
}
