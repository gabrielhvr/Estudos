package main 

import ("fmt"
"github.com/labstack/echo/v4"
)

type Car struct{
	Name string
	Model string
	Price float64
}

var cars []Car


func generateCars() {
	cars = append(cars, Car{Name: "Ferrari", Model: "xv", Price : 1000000})
	cars = append(cars, Car{Name: "Honda", Model: "Civic", Price: 20000})
	cars = append (cars, Car{Name: "Toyota", Model: "corolla", Price: 19000})
}


func (c Car) Andar(){
	fmt.Println("O Carro", c.Name, "está andando")
}

func getCars(c echo.Context) error{
	return c.JSON(200, cars)

}

func createCars(c echo.Context) error{
	car := new(Car)
	if err := c.Bind(car); err != nil{
		return err
	}
	cars = append(cars, *car)
	return c.JSON(200, cars)
}

func main() {
	generateCars()

	e := echo.New()
	e.GET("/cars", getCars)
	e.POST("/cars", createCars)
	e.Logger.Fatal(e.Start(":8080"))

}


// Banco de dados
/*
func saveCar(car Car) error{
	db, err := sql.Open("sqlite3", "cars.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare ("INSERT INTO cars (name, prime) VALUES ($1, $2)")
	if err != nil {
		return err
	}

	return nil
}
*/