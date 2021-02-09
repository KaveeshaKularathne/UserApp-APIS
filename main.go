package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type (
	user struct {
		Name  string `json:"name" `
		Email string `json:"email" `
		pnumber string `json:"pnumber" `
		password string `json:"password" `

	}

)
var (
	users=map[string]*user{}

)

func createUser(c echo.Context) error {
	u := &user{

	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.Name] = u

	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context)error  {
	Name:=c.Param("name")
	Email:=c.Param("email")
	Pnumber:=c.Param("pnumber")
	Password:=c.Param("password")
	return c.String(http.StatusOK,Name)
	return c.String(http.StatusOK,Email)
	return c.String(http.StatusOK,Pnumber)
	return c.String(http.StatusOK,Password)

}
func deleteUser(c echo.Context)error  {

	name, _ := strconv.Atoi(c.Param("name"))
	delete(users, string(name))
	return c.NoContent(http.StatusNoContent)

}
func updateUser(c echo.Context)error  {
	u :=new(user)
	if err := c.Bind(u);err!=nil{
		return err
	}
	name,_ :=strconv.Atoi(c.Param("name"))
	delete(users, string(name))
	return c.NoContent(http.StatusNoContent)

}

	func main()  {
	e:=echo.New()

    //e.GET("/users",getUser)
	e.POST("/users",createUser)

	//e.PUT("/users", updateUser)
	//e.DELETE("/users", deleteUser)



	e.Logger.Fatal(e.Start(":8000"))


}

