package todo

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/lib/pq"
	"github.com/julienschmidt/httprouter"
)

func Test_Update(t *testing.T) {

	// handler := controllers.NewBookController()
	router := httprouter.New()
	router.PUT("/todos/:id", Update)

	w := httptest.NewRecorder()
	r, err := http.NewRequest("PUT", "/todos/1", bytes.NewBufferString(`{"title":"car clean", "status": "In Progress"}`))
	if err != nil {
		t.Fatal(err)
	}

	router.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("Wrong status: wanted %v, but got %v", http.StatusOK, w.Code)
	}

}


func Test_Create(t *testing.T) {

	// handler := controllers.NewBookController()
	router := httprouter.New()
	router.POST("/todos", Create)

	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/todos", bytes.NewBufferString(`{"title":"car clean", "status": "In Progress"}`))
	if err != nil {
		t.Fatal(err)
	}

	router.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("Wrong status: wanted %v, but got %v", http.StatusOK, w.Code)
	}

}


func Test_List(t *testing.T) {

	// handler := controllers.NewBookController()
	router := httprouter.New()
	router.GET("/todos", List)

	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/todos", nil)
	if err != nil {
		t.Fatal(err)
	}

	router.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("Wrong status: wanted %v, but got %v", http.StatusOK, w.Code)
	}

}



func Test_Handler(t *testing.T) {

	err:= InitDb()
	if err != nil {
		t.Fatal(err)
	}

	r, err := http.NewRequest("GET", "/todos", nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r.URL)
	//.Println(req.URL)

	w := httptest.NewRecorder()
	List(w, r, httprouter.Params{})

	t.Log(w.Body)
	t.Log(w.Code)

	if w.Code != http.StatusOK {
		t.Errorf("response test failed. wanted %v, but got %v", http.StatusOK, w.Code)
		return
	}

	//TODO: need to add reasponse body test
}