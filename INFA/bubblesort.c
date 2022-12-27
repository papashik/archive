void bubblesort(unsigned long nel, int (*compare)(unsigned long i, unsigned long j), void (*swap)(unsigned long i, unsigned long j)) {
    if (nel != 1 && nel != 0) {
        while (1) {
            char swapped = 0;
            for (int i = 0; i < nel - 1; i++) {
                if (compare(i, i+1) > 0) {
                    swap(i, i+1);
                    swapped = 1;
                }
            }
            if (swapped == 0) {
                break;
            }
            
            swapped = 0;
            for (int i = nel - 2; i >= 0; i--) {
                if (compare(i, i+1) > 0) {
                    swap(i, i+1);
                    swapped = 1;
                }
            }
            if (swapped == 0) {
                break;
            }
        }    
    }
}
