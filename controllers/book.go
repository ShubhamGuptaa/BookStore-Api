package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ShubhamGuptaa/BookStrore/models"
	"github.com/julienschmidt/httprouter"
	// "go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "github.com/gorilla/mux"
)

type BookController struct{
	session *mgo.Session
}

func NewBookController(s *mgo.Session) *BookController {
	return &BookController{s}
}

// ////////////////////////// GetBook By Id => Working ////////////////////
func (uc BookController) GetBook (w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id:= p.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
	}
		oid:=bson.ObjectIdHex(id)

		
		u:= models.Book{}

	// if err := 
	uc.session.DB("BookStore").C("books").FindId(oid).One(&u);
	//  err != nil {
			// w.WriteHeader(404)
			// return
		// }
		uj, err := json.Marshal(u);
		if err != nil {
			fmt.Println(err)
		}
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w,"%s\n",uj)
}

// ///////////////////////// Get All Books => Working ////////////
func (uc BookController) GetBooksAll (w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	type Books []models.Book
		u:= Books{}
	    uc.session.DB("BookStore").C("books").Find(bson.M{}).All(&u);
		uj, err := json.Marshal(u);
		if err != nil {
			fmt.Println(err)
		}
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w,"%s\n",uj)

}
// /////////////////////////////////////


// ///////////////////////// GetBook By name =>  Working////////////
func (uc BookController) GetBookByName (w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name:= p.ByName("name")
	
		type Books []models.Book
		u:= Books{}
	    uc.session.DB("BookStore").C("books").Find(bson.M{"name": &name}).One(&u);
		uj, err := json.Marshal(u);
		if err != nil {
			fmt.Println(err)
		}
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w,"%s\n",uj)

}
// ///////////////////////// GetBook By Ratings =>  ////////////
func (uc BookController) GetBookByRating (w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ratings:= p.ByName("rating")
		
			u:= models.Book{}
	
			uc.session.DB("BookStore").C("books").Find(bson.M{"ratings": &ratings}).One(&u);
			uj, err := json.Marshal(u);
			if err != nil {
				fmt.Println(err)
			}
			w.Header().Set("Content-Type","application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w,"%s\n",uj)
	
	}


// //////////////////////// Creating Books => Working /////////////

func (uc BookController) CreateBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u:= models.Book{}

	json.NewDecoder(r.Body).Decode(&u)
	u.Id = bson.NewObjectId()
	uc.session.DB("BookStore").C("books").Insert(u)

	uj,err:= json.Marshal(u)

	if err != nil{
		fmt.Println(err)
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w,"%s \n",uj)
}

///////////////////////////////////// Deleting Books => Working /////////////
func (uc BookController) DeleteBook(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id:= p.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(404)
		return
	}

	oid:= bson.ObjectIdHex(id)

	if err:= uc.session.DB("BookStore").C("books").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w,"Deleted Book with id: ",oid,"\n")
}

// ///////////////////////////////////
