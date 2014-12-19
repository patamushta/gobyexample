// Go имеет множество встроенных типов данных, включая
// строки (string), целые числа (integer), дробные числа
// (float), логический тип (boolean) и т.д. Тут представлены
// несколько простых примеров.

package main

import "fmt"

func main() {

    // Строки можно объединять оператором `+`.
    fmt.Println("go" + "lang")

    // Целые и дробные числа.
    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    // Логические значения и операторы аналогичны таковым в других языках.
    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}
