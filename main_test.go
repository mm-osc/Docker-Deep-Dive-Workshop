package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(handler)
	hf.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `<html>
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
										    </html>`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}
