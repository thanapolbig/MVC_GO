package main
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)


func Pong(c *gin.Context)  {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func SubString(c *gin.Context) {
	var input text
	e := c.BindJSON(&input)
	if e != nil {
		fmt.Println(e)
	}
	var Stringlen = len(input.String)
	if(Stringlen > 255 || Stringlen < 20){
		c.JSON(400,gin.H{
			"error" : "error len out 255 or 20",
		})
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

	//for i := 0 ;i< len(Split);i++{
	//		result[0] = map[string]int{"foo": 1, "bar": 2}
	//}

	c.JSON(200, gin.H{
		"message": len,
	})
}
