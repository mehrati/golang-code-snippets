package main

import ("fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)
type st struct {
	i int
	s string 
}
func main(){
	db,_ := sql.Open("mysql","root:password@/")
	defer db.Close()
	var ver string
	db.QueryRow("SELECT version();").Scan(&ver)
	fmt.Println(ver)
}
