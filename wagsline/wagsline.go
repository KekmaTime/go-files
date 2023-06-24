package wagsline

import "fmt"

func main() {
	var username string = "wagsline"
	var password int = 208493274

	fmt.Println("Authorization: Basic", username+ ":", password)
}
