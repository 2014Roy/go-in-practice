package apis

import (
	. "AppApi/project/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "it works")
}

func AddPerson(c *gin.Context) {
	var person Person
	if err := c.Bind(&person); err != nil {
		log.Fatalln(err)
	}
	index, err := person.AddPerson()
	if err != nil {
		log.Fatal(err)
		return
	}
	msg := fmt.Sprintf("insert successful %d name: %s %s", index, person.FirstName, person.LastName)

	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func DeletePerson(c *gin.Context) {
	type IDSturct struct {
		Id string `json:"id"`
	}
	var sId IDSturct
	if err := c.BindJSON(&sId); err == nil {
		fmt.Printf("cid is : %s \n", sId.Id)
		id, err := strconv.Atoi(sId.Id)
		if err != nil {
			log.Fatalln(err)
		}
		var person = Person{Id: id, FirstName: "", LastName: ""}
		index, err := person.DeletePerson()
		if err != nil {
			log.Fatalln(err)
		}

		msg := fmt.Sprintf("delete person %d successful", index)
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}
}

func GetPerson(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)

	if err != nil {
		log.Fatalln(err)
	}
	person, err := SearchPerson(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"person": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"person": person,
	})
}
