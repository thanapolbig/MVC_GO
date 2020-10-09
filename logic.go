package main
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)


func Pong(c *gin.Context)  {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (h *_DataHandler) EndPointDataCleansing(c *gin.Context) {

	var input text
	e := c.BindJSON(&input)
	if e != nil {
		fmt.Println(e)
	}

	result, err := h.SubString(input)
	if err != "" {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, result)
}

func (h *_DataHandler) SubString(input text)(result []return_ans, err string){
	var Stringlen = len(input.String)
	if(Stringlen > 255 || Stringlen < 20){
		return
	}
	var lowercase = strings.ToLower(input.String)
	fmt.Println(lowercase)

	lowercase = strings.ReplaceAll(lowercase, "?", "")
	lowercase = strings.ReplaceAll(lowercase, "!", "")
	lowercase = strings.ReplaceAll(lowercase, ",", "")

	fmt.Println(lowercase)

	var Split = strings.Split(lowercase , " ")
	fmt.Println(Split)

	MapCount := make(map[string]*return_ans)

	for _, word := range Split {

		if valueInMap, ok := MapCount[word]; ok {
			valueInMap.Count++
		} else {
			MapCount[word] = &return_ans{
				String_data:  word,
				Count: 1, //count begin 1
			}
		}
	}
	fmt.Println("inputString : ", input.String)

	for _, val := range MapCount {
		result = append(result, *val)
	}

	return
}
