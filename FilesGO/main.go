package main

import (
	"fmt"
	"encoding/csv"
	"os"
)
func main() {
	f, _ := os.Create("file.txt")
	defer f.Close()

	f.WriteString("Hello, Go!")

	data, _ := os.ReadFile("file.txt")
	fmt.Println(string(data))

	c, _ := os.Create("file.csv")

	writer := csv.NewWriter(c)
	writer.Write([]string{"Name", "Age"})
	writer.Write([]string{"John", "30"})
	writer.Flush()

	c.Seek(0, 0)	

	reader := csv.NewReader(c)
	records, _ := reader.ReadAll()
	fmt.Println(records)

}
