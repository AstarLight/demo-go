
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
)

func main() {
	r := gin.Default()

	r.GET("/query_user", queryUser)

	r.Run(":8080")

}


// curl "http://127.0.0.1:8080/query_user?username=lijunshi"
func queryUser(c *gin.Context) {
	username := c.Query("username")

	err := bindUser(username)
	if err != nil {

		// %s: Returns the user-safe error string mapped to the error code or the error message if none is specified.
		fmt.Println("====================> %s <====================")
		fmt.Printf("%s\n\n", err)
	
		// %v: Alias for %s.
		fmt.Println("====================> %v <====================")
		fmt.Printf("%v\n\n", err)
	
		// %-v: Output caller details, useful for troubleshooting.
		fmt.Println("====================> %-v <====================")
		fmt.Printf("%-v\n\n", err)
	
		// %+v: Output full error stack details, useful for debugging.
		fmt.Println("====================> %+v <====================")
		fmt.Printf("%+v\n\n", err)
	
		// %#-v: Output caller details, useful for troubleshooting with JSON formatted output.
		fmt.Println("====================> %#-v <====================")
		fmt.Printf("%#-v\n\n", err)
	
		// %#+v: Output full error stack details, useful for debugging with JSON formatted output.
		fmt.Println("====================> %#+v <====================")
		fmt.Printf("%#+v\n\n", err)
	
		// do some business process based on the error type
		if errors.IsCode(err, ErrDecodingJSON) {
			fmt.Println("this is a ErrEncodingFailed error")
		}
	
		if errors.IsCode(err, ErrDatabase) {
			fmt.Println("this is a ErrDatabase error")
		}

		// we can also find the cause error
		fmt.Println(errors.Cause(err))

		coder := errors.ParseCoder(err)

		c.JSON(200, gin.H{
			"code": coder.HTTPStatus(),
			"msg": coder.String(),
		})

		return
	}

	c.JSON(200, gin.H{
		"code": 0,
		"msg": "ok",
	})

}

func bindUser(username string) error {
	if err := getUser(username); err != nil {
	// Step3: Wrap the error with a new error message and a new error code if needed.
	return errors.WrapC(err, ErrDecodingJSON , fmt.Sprintf("encoding user %s failed.", username))
  }

  return nil
}

func getUser(username string) error {
  if err := queryDatabase(username); err != nil {
	// Step2: Wrap the error with a new error message.
	return errors.WrapC(err, ErrDecodingJSON, "get user failed.")
  }

  return nil
}

func queryDatabase(username string) error {
  // Step1. Create error with specified error code.
  return errors.WithCode(ErrDatabase, "user %s not found.", username)
}