package main
import "fmt"

func partition(low, high int, less func(i, j int) bool, swap func(i, j int)) int {
    pivot := high;
    j := low;
    for i := low; i < high; i++ {
        if (less(i, pivot)) {
            swap(i, j);
            j++;
        }
    }
    swap(j, pivot);
    
    return j;
}

func qsort_rec(low, high int, less func(i, j int) bool, swap func(i, j int)) {
    if (low >= high) { 
        return; // Выход из рекурсии
    }
        pivot := partition(low, high, less, swap);
        qsort_rec(low, pivot - 1, less, swap);
        qsort_rec(pivot + 1, high, less, swap);
}


func qsort(n int, less func(i, j int) bool, swap func(i, j int)) {
    qsort_rec(0, n - 1, less, swap);
}


func main() {
    var n int;
    fmt.Scanf("%d", &n);
    
    arr := make([]int, n);
    
    for i := range arr {
        fmt.Scanf("%d", &arr[i]);
    }
    
    qsort(len(arr), func(i, j int) bool { return arr[i] < arr[j]; },
                func(i, j int) { arr[i], arr[j] = arr[j], arr[i]; });
    
    for i := range arr {
        fmt.Printf("%d ", arr[i]);
    }
}
