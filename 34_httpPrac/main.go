package main

import (
	"io"
	"net/http"
	"fmt"
	"os/exec"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "hello, world!\n")
	if err != nil {
		fmt.Println(err)
	}
}

func StartDisplay(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w,
		`
		 <!DOCTYPE html>
		 <html>
		 <body>
		 <h1>My First Heading</h1>
		 <p id="firstheading">My first paragraph.</p>
		 <button onclick="changetext()">Click for next video</button>
		 <script>
		 function changetext() {
		 document.getElementById("firstheading").innerHTML = "fucvk";
		 }
		 </script>
		 </body>
		 </html>
		`)
	// open the firefox browser to our site
	cmd := exec.Command("firefox", "localhost:9000/hello")
	cmd.Run()
	fmt.Println("command fired")
}

func ChooseVideo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "need to flesh this out")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/start", StartDisplay)
	http.HandleFunc("/video/id", ChooseVideo)
	http.ListenAndServe("localhost:9000", nil)
}
