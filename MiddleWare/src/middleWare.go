package main

import (
	"log"
	"net/http"
)


func middleWare(handler http.Handler) http.Handler {

	return  http.HandlerFunc( func( response http.ResponseWriter , request *http.Request ) {
		log.Print("Executing before Phase")
		handler.ServeHTTP(response,request)
		log.Println("Executing after  phase ")

	})
}

func httpHandlerFuncMainLogic ( response http.ResponseWriter , request *http.Request){
	log.Println("Executing Main logic ")
	response.Write([]byte("Welcome    VIKRAM !!!!!!!!!!!!!!!!!!!"))

}

func  main ()  {

	log.Print("Starting server ************************************************")
	mainLogicHandler:= http.HandlerFunc(httpHandlerFuncMainLogic)

	http.Handle("/api",middleWare(mainLogicHandler ))

	log.Print("HIT IP:PORT/api ************************************************")
	http.ListenAndServe(":9093",nil)

}

