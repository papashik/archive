# Golang - основы. Функции (ДЗ №3)
Разработать функции, названия и сигнатуры которых указаны ниже:<br>
1. `MyRange(start, stop, step int) []int` - функция возвращает срез, содержащий целые числа в полуинтервале `[start, stop)` с шагом `step`;<br>
2. `MyPower(a float64, e int) float64` - функция возводит **вещественное** число `a` в **целую** степень `e`;<br>
3. `MyFactorial(n int) int` - функция находит факториал числа `n`;<br>
4. `MyGcd(a, b int) int` - функция возвращает НОД двух чисел;<br>
5. `IsSymmetric(arr []int) bool` - функция проверяет, является ли данный срез симметричным (палиндромом);<br>
6. `MyReverse(arr []int) []int` - функция возвращает развернутый срез;<br>
7. `MyShorten(a, b int) (int, int)` - функция сокращает дробь $\frac{a}{b}$ и возвращает два числа `p` и `q` такие, что $\frac{a}{b}=\frac{p}{q}$, причем дробь $\frac{p}{q}$ несократимая (`p` - целое число, **`q` - натуральное**, число 0 представлять в виде дроби $\frac{0}{1}$);<br>
8. `IsPrime(a int) bool` - функция проверяет, является ли данное число простым (`a > 0`);
### Множество - это набор неповторяющихся элементов
9. `MySet(arr []int) []int` - функция преобразует переданный ей срез в множество (удаляет все повторяющиеся элементы);<br>
10. `IsSet(arr []int) bool` - функция проверяет, является ли переданный ей срез множеством.<br>
