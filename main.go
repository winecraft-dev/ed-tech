package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/winecraft-dev/ed-tech/database"
	"github.com/winecraft-dev/ed-tech/sections"
	"github.com/winecraft-dev/ed-tech/students"
)

func main() {
	dbConn, err := database.NewConnection()
	if err != nil {
		log.Fatalln(err)
	}

	secServ := sections.NewService(dbConn)
	studServ := students.NewService(dbConn)

	e := echo.New()
	e.HTTPErrorHandler = handleError

	secg := e.Group("/sections")
	secg.POST("", secServ.CreateSection)
	secg.GET("", secServ.GetSections)
	secg.GET("/:id", secServ.GetSection)

	stug := e.Group("/students")
	stug.POST("", studServ.CreateStudent)
	stug.GET("/:id", studServ.GetStudent)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Problem starting server: %v\n", err)
	}
}

func handleError(err error, c echo.Context) {
	c.Logger().Error(err)

	c.Echo().DefaultHTTPErrorHandler(err, c)
}
