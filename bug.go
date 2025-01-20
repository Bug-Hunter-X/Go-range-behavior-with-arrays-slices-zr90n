func main() {


  x := make([]int, 0, 5)

  x = append(x, 1, 2, 3, 4, 5)


  fmt.Println(x)

  for i := range x {

    fmt.Println(i)

  }
}