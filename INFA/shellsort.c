void shellsort(unsigned long nel, int (*compare)(unsigned long i, unsigned long j), void (*swap)(unsigned long i, unsigned long j)) {
    if (nel == 0 || nel == 1) { return; }
    
    long int f1 = 0, f2 = 1, temp = 0;
    while (1) {
        temp = f2;
        f2 = f1 + f2;
        f1 = temp;
        if (f2 >= nel) { break; }
    }
    for (long int s = f1; s > 0; ) {
        for (long int i = s; i < nel; i++) {
            for (long int j = i - s; j >= 0 && compare(j, j + s) > 0; j -= s) {
                swap(j, j + s);
            }
        }
        temp = f1;
        f1 = f2 - f1;
        f2 = temp;
        s = f1;
    }
}
