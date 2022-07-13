package main

import (
	"net/http"
	"os"
	"example.com/controllers"
	"github.com/joho/godotenv"
)
func main() {
	err:=godotenv.Load(".env")
	if err!=nil{
		panic("Error loading the .env file")
	}
	/*noOfRetry,err:=strconv.Atoi(os.Getenv("NoOfRetry"))
	if err!=nil{
		println("NoOfRetry conversion error")
		return
	}*/
	controllers.RegisterContorllers()
	http.ListenAndServe(os.Getenv("SERVER")+":"+os.Getenv("PORT"),nil)
}
