package calculator

func Add(arr []int) int {
    var sum int = 0
    for i := 0; i < len(arr); i++ {
        sum += arr[i]
    }
    return sum
}

func Sub(arr []int) int {
    var substact int = arr[0]
    for _, num := range arr[1:] {
        substact -= num
    }
    return substact
}

func Mul(arr []int) int {
    var mul int = 1
    for _, num := range arr {
        mul *= num
    }
    return mul
}

func Div(arr []int) int {
    var div int = arr[0]
    for _, num := range arr[1:] {
        div /= num
    }
    return div
}
