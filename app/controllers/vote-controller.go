package controllers

import (
	"fmt"
	"github.com/cahyo0118/votever-vote-service/models"

	"github.com/revel/revel"
)

type VoteController struct {
	*revel.Controller
}

// type Vote struct {
// 	Name        string `json:"name"`
// 	Description string `json:"description"`
// }

var votes []Vote

func (c VoteController) GetAll() revel.Result {

	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)

	votes := append(
		votes,
		Vote{
			Name:        "tayoo",
			Description: "tayoo",
		},
	)

	fmt.Println(votes)

	return c.RenderJSON(votes)
}

func (c VoteController) GetOne() revel.Result {

	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)

	// v := c.Params.JSON

	// data := make(map[string]interface{})
	// data["error"] = nil
	// stuff := Stuff{Foo: "xyz", Bar: 999}
	// data["stuff"] = stuff
	return c.RenderJSON(jsonData)
}

func (c VoteController) Store() revel.Result {

	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)

	// v := c.Params.JSON

	// data := make(map[string]interface{})
	// data["error"] = nil
	// stuff := Stuff{Foo: "xyz", Bar: 999}
	// data["stuff"] = stuff
	return c.RenderJSON(jsonData)
}

func (c VoteController) Update() revel.Result {

	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)

	// v := c.Params.JSON

	// data := make(map[string]interface{})
	// data["error"] = nil
	// stuff := Stuff{Foo: "xyz", Bar: 999}
	// data["stuff"] = stuff
	return c.RenderJSON(jsonData)
}

func (c VoteController) Delete() revel.Result {

	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)

	// v := c.Params.JSON

	// data := make(map[string]interface{})
	// data["error"] = nil
	// stuff := Stuff{Foo: "xyz", Bar: 999}
	// data["stuff"] = stuff
	return c.RenderJSON(jsonData)
}
