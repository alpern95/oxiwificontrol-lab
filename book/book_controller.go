package book

import (
	"github.com/emicklei/go-restful"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"gitlab.com/bytecraze/go-restful-api/auth"
	"gitlab.com/bytecraze/go-restful-api/db"
)

type BookController struct {
}

func (controller BookController) AddRouters() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/api/v1/book").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	ws.Route(ws.POST("/").Filter(auth.BearerAuth).To(createBook))
	ws.Route(ws.GET("/").Filter(auth.BearerAuth).To(listBooks))
	ws.Route(ws.GET("/{bookId}").Filter(auth.BearerAuth).To(getBook))
	ws.Route(ws.PUT("/{bookId}").Filter(auth.BearerAuth).To(updateBook))
	ws.Route(ws.DELETE("/{bookId}").Filter(auth.BearerAuth).To(deleteBook))

	return ws
}

func createBook(req *restful.Request, resp *restful.Response) {
	book := Book{}
	err := req.ReadEntity(&book)
	if err != nil {
		resp.WriteHeaderAndEntity(400, "invalid request")
		return
	}

	book.ID = bson.NewObjectId()
	session := db.NewDBSession()
	defer session.Close()
	c := session.DB("").C("book")
	err = c.Insert(book)
	if err != nil {
		resp.WriteError(500, err)
		return
	}

	resp.WriteEntity(book)
}

func listBooks(req *restful.Request, resp *restful.Response) {
	allBooks := make([]Book, 0)
	session := db.NewDBSession()
	defer session.Close()
	c := session.DB("").C("book")
	err := c.Find(bson.M{}).All(&allBooks)
	if err != nil {
		resp.WriteError(500, err)
		return
	}

	resp.WriteEntity(allBooks)
}

func getBook(req *restful.Request, resp *restful.Response) {
	bookId := req.PathParameter("bookId")
	book := Book{}
	session := db.NewDBSession()
	defer session.Close()
	c := session.DB("").C("book")
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(bookId)}).One(&book)
	if err != nil {
		if err == mgo.ErrNotFound {
			resp.WriteError(404, err)
		} else {
			resp.WriteError(500, err)
		}
		return
	}

	resp.WriteEntity(book)
}

func updateBook(req *restful.Request, resp *restful.Response) {
	bookId := req.PathParameter("bookId")
	book := Book{}
	err := req.ReadEntity(&book)
	if err != nil {
		resp.WriteHeaderAndEntity(400, "invald request")
		return
	}

	book.ID = bson.ObjectIdHex(bookId)
	session := db.NewDBSession()
	defer session.Close()
	c := session.DB("").C("book")
	err = c.Update(bson.M{"_id": book.ID}, book)
	if err != nil {
		if err == mgo.ErrNotFound {
			resp.WriteError(404, err)
		} else {
			resp.WriteError(500, err)
		}
		return
	}

	resp.WriteEntity(book)
}

func deleteBook(req *restful.Request, resp *restful.Response) {
	bookId := req.PathParameter("bookId")

	session := db.NewDBSession()
	defer session.Close()
	c := session.DB("").C("book")
	err := c.RemoveId(bookId)
	if err != nil {
		if err == mgo.ErrNotFound {
			resp.WriteError(404, err)
		} else {
			resp.WriteError(500, err)
		}
		return
	}

	resp.WriteHeader(200)
}
