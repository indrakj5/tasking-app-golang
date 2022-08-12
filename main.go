package main

import (
	"net/http"
	"tasking-app-golang/controllers/taskcontroller"
)

func main() {
	http.HandleFunc("/", taskcontroller.Index)
	http.HandleFunc("/task/get_form", taskcontroller.GetForm)
	http.HandleFunc("/task/store", taskcontroller.Store)
	http.HandleFunc("/task/delete", taskcontroller.Delete)
	http.HandleFunc("/task/update_task", taskcontroller.UpdateTask)

	http.ListenAndServe(":8000", nil)
}
