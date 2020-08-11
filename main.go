package main

import (
	"net/http"
    //"time"
    "github.com/go-chi/chi"
)
type Storage struct {
    // dict or DB
    S ServStorage
}
type ServStorage struct {
    data []struct {
        id string //`json:"id"`
        url string
        interval uint8
    }
    history []struct {
        id string
        history []struct {
            response string
            duration float32
            start float32
        }
    }
}

func worker(){
    //repeat in interval
    // store start time
//   r, err :=  http.Get();
/*   select {
        case <-r:
            //store response
        case <-err:
            //store null
        case <-time.After('5s'):
            //store null
        case <-stopWorkerChan:
            done <- true
    }
    store duration
*/
}


func (storage *Storage) h1(w http.ResponseWriter, req *http.Request){
    // GET /api/fetcher
    // response [{"id":1,"url":"https://httpbin.org/range/15","interval":60},
    // {"id":2,"url": "https://httpbin.org/delay/10","interval":120}]
    w.Write([]byte("HANDLER list\n"+ req.RequestURI + "\n"))
}
func (storage *Storage) h2(w http.ResponseWriter, req *http.Request){
    // Create/update a new record.
    // Save curl -si 127.0.0.1:8080/api/fetcher -X POST -d '{"url":
    // "https://httpbin.org/range/15","interval":60}'


    // add data to some DB/local storage
    // response "{"id": 1"}
    // if data > 1M 413 Request entity to large
    // if wrong data (not complete) 400 BadRequest

    // if running stop worker, then create new one
    w.Write([]byte("HANDLER add/update\n"+ req.RequestURI + "\n" ))
}
func (storage *Storage) h3(w http.ResponseWriter, req *http.Request){
    // Remove the record.
    // api/fetcher/id
    id := chi.URLParam(req, "id")
    // retrieve ID from request
    w.Write([]byte("HANDLER Delete\n" + req.RequestURI + id + "\n" ))

}
func (storage *Storage) h4(w http.ResponseWriter, req *http.Request){
    // GET /api/fetcher/<ID>/history
    // get list of events for this ID
    w.Write([]byte("HANDLER history\n"+ req.RequestURI + "\n"))
}

func main() {
    storage := &Storage{}
    r := chi.NewRouter()
    r.Route("/api/fetcher", func (r chi.Router) {
        r.Get("/", storage.h1)
        r.Post("/", storage.h2)
        r.Delete("/{id}", storage.h3)
        r.Get("/{id}/history", storage.h4)
    })

    http.ListenAndServe(":8080", r)
}
