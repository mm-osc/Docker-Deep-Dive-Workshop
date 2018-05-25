package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handler).Methods("GET")
	http.ListenAndServe(":8445", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<html>
	  <head>
	      <title>Hi DevOps</title>
	          <link rel="stylesheet" href="http://netdna.bootstrapcdn.com/bootstrap/3.1.1/css/bootstrap.min.css"/>
		      <style>
		      body {
			          background-color: lightblue;
			  }
			  </style>
			    </head>
			      <body>
			          <div class="container">
				        <div class="jumbotron">
					        <h1>May The Source Be With You</h1>
						        <p>Click the button below to view #MOSC...</p>
							        <p> <a class="btn btn-primary" href="https://www.mmosc.org">MOSC</a></p>
								        <p></p>
									      </div>
									          </div>
										    </body>
										    </html>`)
}
