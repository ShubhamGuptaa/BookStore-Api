// Task: 
// Create a bookinfo application in which user will provide book name and you will return book details
// Create 3 microservices
// 1. Details Microservice  (return details from MongoDB based on Name)
// 2. Rating Microservice   (return book rating from MongoDB based on Name)
// 3. Product Microservice  (return list of all the books)
// Create Restful Services and integrate the
// Use MongoDB   
// ///////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"
	"github.com/ShubhamGuptaa/BookStrore/controllers"
)

func main() {

		r:= httprouter.New()
		uc:= controllers.NewBookController(getSession())
		r.GET("/book/product",uc.GetBooksAll)				     // Get all product details ============== On Task>
		r.GET("/book/ById/:id",uc.GetBook)				    	//  Get details by id
		r.GET("/book/Rating/:rating",uc.GetBookByRating)       //  Get details by ratings  =============== On Task> 
		r.GET("/book/details/:name",uc.GetBookByName)		  // Get details by Name   =================== On Task>
		r.POST("/book",uc.CreateBook)                        //  Create new record
		r.DELETE("/book/:id",uc.DeleteBook)					//  Delete the record by Id
		http.ListenAndServe("localhost:8080",r)			   // localhost listner on port no. 8080
}

func getSession() *mgo.Session {
	s,err:= mgo.Dial("mongodb://localhost:27017")
	if err != nil{
		panic(err)
	}
	return s
}
