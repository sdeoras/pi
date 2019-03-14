# pi getter
get arbitrary digits of pi using https://pi.delivery/

## usage
```go
import "github.com/sdeoras/pi"

func main() {
	digits, decimal, err := pi.Get(0, 10)
	
	fmt.Println(digits) // prints [3 1 4 1 5 9 2 6 5]
	fmt.Println(decimal) // prints 1 indicating decimal after 3
}
```