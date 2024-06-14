# Конкин Иван - тестовая работа на Golang
## Кальтулятор строково-числовой
### 1. Установка

- git clone;
- cd to app folder;

### Использование

- ```go mod init <module name>``` // default: go mod init test;
- ```go run main.go``` // compile and run program.

### Структура программы
```bash
├── calc
│   ├── engine
│   │   └── engine.go
│   ├── middleware
│   │   └── utilities.go
│   ├── stringAndNumber.go
│   └── stringAndString.go
├── core
│   ├── getInput.go
│   ├── middlware
│   │   ├── fields.go
│   │   └── inspector.go
│   └── validator.go
├── go.mod
├── main.go
└── README.md
```
### Алгоритм программы
1. Получаем данные из вода в консоли.
2. Валидируем полученные данные и разбиваем на три переменные (оператор и два амперсанда).
3. Проверяем, если второй амперсанд вернулся с типом int, то запускаем калькулятор "строка и число".
4. Проверяем, если второй амперсанд вернулся с типом string, то запускаем калькулятор "строка и строка".

```go
input := core.GetInput()
firstOperand, mathOperator, lastOperand := core.InputExpression(input)
if reflect.TypeOf(lastOperand).Kind() == reflect.Int {
    output := calc.StringAndNumber(firstOperand, mathOperator, lastOperand)
    result := middlware.CheckLength(output, 40)
    fmt.Printf(result)
} else {
    output := calc.StringAndString(firstOperand, mathOperator, lastOperand)
    result := middlware.CheckLength(output, 40)
    fmt.Println(result)

}
```