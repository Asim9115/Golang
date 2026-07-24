// package main

// //--------------------------------SOCKET-----------------------
// import (
// 	"log"
// 	"net/http"
	
// 	"time"
// )

// var myChannel = make(chan string, 100)

// func main() {

// 	http.HandleFunc("/socket", HandleSOCKET)

// 	go logs(myChannel)

// 	log.Println("server listening on port :8080")
// 	http.ListenAndServe(":8080", nil)
// }

// func HandleSOCKET(w http.ResponseWriter, r *http.Request) {

// }



// func logs(myChannel chan string) {

// 	myChannel <- "step1"
// 	time.Sleep(2 * time.Second)

// 	myChannel <- "step2"
// 	time.Sleep(2 * time.Second)

// 	myChannel <- "step3"
// 	time.Sleep(2 * time.Second)

// 	myChannel <- "done"

// }



















//---------------------------------------------SSE----------------------------
package main
import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var myChannel = make(chan string, 100)

func main() {
	var wg sync.WaitGroup
	http.HandleFunc("/sse", HandleSSE)

	log.Println("server started on port :8080")

	wg.Add(4)
	go logs(myChannel, &wg)
	go logs(myChannel, &wg)
	go logs(myChannel, &wg)
	go logs(myChannel, &wg)

	go func() {
		wg.Wait()

		// All senders are finished
		close(myChannel)
	}()
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func HandleSSE(w http.ResponseWriter, r *http.Request) {
		// Set SSE headers
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	for message := range myChannel {
		fmt.Fprintf(w, "data: %s\n\n", message)

		// Send data immediately to the client
		flusher.Flush()
	}
}


func logs(myChannel chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	myChannel <- "step1"
	time.Sleep(2 * time.Second)

	myChannel <- "step2"
	time.Sleep(2 * time.Second)

	myChannel <- "step3"
	time.Sleep(2 * time.Second)

	myChannel <- "done"

}