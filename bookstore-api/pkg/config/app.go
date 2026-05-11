// Connect your Go app to the database and keep that connection ready to use.
package config

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB //db is a reference to the database connection managed by GORM
)

// GORM is a Go library that lets you use Go structs and methods to create, read, update, and delete data in a database without writing SQL every time.
func Connect() {
	d, err := gorm.Open("mysql", "root:xyz%40%23@tcp(localhost:3306)/parship?charset=utf8&parseTime=True&loc=Local") // Calls GORM's function to open a database connection. Returns: d → database object, err → error if something goes wrong
	if err != nil{
		panic(err)
	}
	db = d // Saves the database connection in the global variable db. Now the whole application can reuse it
}

func GetDB() *gorm.DB {
    return db  // Returns the database connection
}


/*
GORM is a Go library that lets you talk to a database using Go code instead of SQL.

gorm.DB = GORM’s database object
*gorm.DB = reference to that object
db = your variable holding it
*/