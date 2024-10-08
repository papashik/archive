package main
import ("fmt"; "math")

func printMatr(matr [][]int) {
    row_len := len(matr[0]);
    col_len := len(matr);
    for i := 0; i < col_len; i++ {
        for j := 0; j < row_len; j++ {
            fmt.Printf("%d ", matr[i][j]);
        }
        fmt.Printf("\n");
    }
    fmt.Printf("-----------\n");
}

func swapRows(matr [][] int, str1, str2 int) {
    // меняет местами строки под номерами str1 и str2
    row_len := len(matr[0]);
    for ind := 0; ind < row_len; ind++ {
        matr[str1][ind], matr[str2][ind] = matr[str2][ind], matr[str1][ind];
    }
}

func multiplyRow(matr [][] int, str, num int) {
    // умножает строку под номером str на число num
    row_len := len(matr[0]);
    for ind := 0; ind < row_len; ind++ {
        matr[str][ind] *= num;
    }
}

func addRowToRow(matr [][] int, dest, src, multiplier int) {
    /* dest - номер строки, к которой прибавится строка
    под номером src, умноженная на multiplier */
    row_len := len(matr[0]);
    for ind := 0; ind < row_len; ind++ {
        matr[dest][ind] += (matr[src][ind] * multiplier);
    }
}

func makeStepView(matr [][] int) {
    //row_len := len(matr[0]);
    col_len := len(matr);

    for work_row := 0; work_row < col_len; work_row++ {
        // Обходим все строки, work_row - номер строки для обнуления остальных
        /* Сначала найдем хотя бы один ненулевой
        первый элемент и поставим эту строку в начало */
        row_num := -1;
        for ind := work_row; ind < col_len; ind++ {
            if (matr[ind][work_row] != 0) {
                row_num = ind;
                break;
            }
        }
        if (row_num == -1) {
            continue; // Идем к следующей колонке, в этой все нули
        }

        // Элемент [row_num][work_row] не равен нулю
        swapRows(matr, work_row, row_num);
        // Теперь элемент [work_row][work_row] не равен нулю

        for row_ind := work_row + 1; row_ind < col_len; row_ind++ {
            if (matr[row_ind][work_row] != 0) {
                multiplier := matr[row_ind][work_row];
                multiplyRow(matr, row_ind, (-1) * matr[work_row][work_row]);
                addRowToRow(matr, row_ind, work_row, multiplier);
            }
        }
    }
}

func findSolutions(matr [][]int) (sol_numerator, sol_denominator []int, err int) {
	// Подается матрица в ступенчатом виде
    row_len := len(matr[0]);
    col_len := len(matr);
    sol_numerator = make([]int, col_len);
    sol_denominator = make([]int, col_len);
    for row_ind := col_len - 1; row_ind >= 0; row_ind-- {
        // Если в матрице есть нулевая строка, то она в конце,
        // и тогда решений ноль или бесконечность
        if (matr[row_ind][row_ind] == 0) {
            if (matr[row_ind][row_len - 1] != 0) {
                // Тогда решений нет
                err = 0;
            } else {
                // Решений бесконечность
                err = -1;
            }
            return;
        }
    }

	num, den, sol_num, sol_den := 0, 0, 0, 0;
    for i := col_len - 1; i >= 0; i-- {
    	num, den = matr[i][row_len - 1], 1;
		for j := i + 1; j < col_len; j++ {
			sol_num = sol_numerator[j] * matr[i][j];
			sol_den = sol_denominator[j];
			num, den = sumFractions(num, den, -sol_num, sol_den);
		}
		den *= matr[i][i];
		num, den = simplifyFraction(num, den);
		sol_numerator[i] = num;
		sol_denominator[i] = den;
    }
    err = 1;
    return;
}

func findNOD(int1, int2 int) int {
    for {
        int1 = int1 % int2;
        if (int1 == 0) { break; }
        int2 = int2 % int1;
        if (int2 == 0) { break; }
    }
    return int1 + int2;
}

func sumFractions(num1, den1, num2, den2 int) (int, int) {
	num1_new := num1 * den2;
	num2_new := num2 * den1;
	den_new  := den1 * den2;
	return simplifyFraction(num1_new + num2_new, den_new);
}

func simplifyFraction(numerator, denominator int) (int, int) {
    if (denominator < 0) {
        numerator, denominator = (-1) * numerator, (-1) * denominator;
    }
    NOD := findNOD(int(math.Abs(float64(numerator))), denominator);

    return numerator / NOD, denominator / NOD;
}

func main() {
    var n int;
    fmt.Scanf("%d", &n);
    matr := make([][]int, n);

    for i := range matr {
        matr[i] = make([]int, (n + 1));
        for j := range matr[i] {
            fmt.Scanf("%d", &matr[i][j]);
        }
    }

    makeStepView(matr);

    sol_num, sol_den, err := findSolutions(matr);

    if (err == -1) {
        fmt.Printf("Infinite solutions\n");
    } else if (err == 0) {
        fmt.Printf("No solution\n");
    } else {
        for i := 0; i < n; i++ {
            num, den := simplifyFraction(sol_num[i], sol_den[i])
            fmt.Printf("%d/%d\n", num, den);
        }
    }
}
