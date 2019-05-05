package main

import (
	"fmt"
	"net/http"
	//"encoding/json"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)
type rating struct {
	Count int
	Value int
}
type tag struct {
	Count int
	Value string
}
type aliases struct{
	Name string
	Sort_name string
}
type artist struct {
	ID   bson.ObjectId `bson:"_id,omitempty"`
	Rating rating      `bson:"rating,omitempty"`
	Name string        `bson:"name"`
	Aliases []aliases
	Area string        `bson:"area"`
	Tags []tag
}

func main() {
	// Echo instance
	e := echo.New()
	e.Use(middleware.CORS())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	search_name:=c.QueryParam("name")
	//search_alias:=c.QueryParam("alias")
	//search_tag:=c.QueryParam("tag")
	fmt.Printf("name : %v\n",search_name)
	session,_:=mgo.Dial("mongodb://localhost:27017")
	defer session.Close()
	db:=session.DB("database064")
	col:=db.C("collection064")
	var p []artist
	col.Find(bson.M{"name":search_name}).All(&p)
	return c.JSON(http.StatusOK, p)
}