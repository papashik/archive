package main
import "fmt"

func add(a, b []int32, p int) []int32 {
    result := make([]int32, len(a));
    
    for i := 0; i < len(result) - 1; i++ {
        summ := a[i] + b[i] + result[i];
        remainder := summ % int32(p);
        result[i] = remainder;
        result[i + 1] = (summ / int32(p));
    }
    
    return result;
}

func main() {
    var p, n int;
    
    fmt.Scanf("%d", &p); // Основание системы счисления
    fmt.Scanf("%d", &n); // Количество элементов массива
    a, b := make([]int32, n), make([]int32, n);
    
    for i := range a {
        fmt.Scanf("%d", &a[i]);
    }
    for i := range b {
        fmt.Scanf("%d", &b[i]);
    }
    
    res := add(a, b, p);
    for _, value := range res {
        fmt.Printf("%d ", value);
    }
}
