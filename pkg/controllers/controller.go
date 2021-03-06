package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/YutaUtah/LeetCodeAPI/pkg/models"
	"github.com/YutaUtah/LeetCodeAPI/pkg/utils"
	"github.com/gorilla/mux"
)

func GetProblem(w http.ResponseWriter, r *http.Request) {
	// store all problems info to newProblems and send it back to the user
	newProblems := models.GetAllProblems()
	// convert newProblems to json object
	res, err := json.Marshal(newProblems)
	if err != nil {
		log.Fatal(err, "json unable to Marshal in getProblem func")
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetProblemById(w http.ResponseWriter, r *http.Request) {
	// get bars (json object)
	vars := mux.Vars(r)
	// get problemId by acccessing problemId attribute
	problemId := vars["Id"]
	ID, err := strconv.ParseInt(problemId, 0, 0)
	if err != nil {
		log.Fatal(err, "error while parsing")
	}
	problemDetails, _ := models.GetProblemById(ID)
	res, err := json.Marshal(problemDetails)
	if err != nil {
		log.Fatal(err, "json unable to Marshal in GetProblemById func")
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateProblem(w http.ResponseWriter, r *http.Request) {
	CreateProblem := &models.Problem{}
	utils.ParseBody(r, CreateProblem)
	b := CreateProblem.CreateProblem()
	res, err := json.Marshal(b)
	if err != nil {
		log.Fatal(err, "json unable to Marshal in CreateProblem func")
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteProblem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	problemId := vars["Id"]
	ID, err := strconv.ParseInt(problemId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	problem := models.DeleteProblem(ID)
	res, err := json.Marshal(problem)
	if err != nil {
		log.Fatal(err, "json unable to Marshal in DeleteProblem func")
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateProblem(w http.ResponseWriter, r *http.Request) {
	var updateProblem = &models.Problem{}
	utils.ParseBody(r, updateProblem)
	vars := mux.Vars(r)
	problemId := vars["Id"]
	ID, err := strconv.ParseInt(problemId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	problemDetails, db := models.GetProblemById(ID)
	if number := strconv.Itoa(updateProblem.Number); number != "" {
		problemDetails.Number = updateProblem.Number
	}
	if updateProblem.Title != "" {
		problemDetails.Title = updateProblem.Title
	}
	if updateProblem.Difficulty != "" {
		problemDetails.Difficulty = updateProblem.Difficulty
	}
	if updateProblem.Comment != "" {
		problemDetails.Comment = updateProblem.Comment
	}
	if updateProblem.Date.String() != "" {
		problemDetails.Date = updateProblem.Date
	}
	// saving problem details in database
	db.Save(&problemDetails)
	res, err := json.Marshal(problemDetails)
	if err != nil {
		log.Fatal(err, "json unable to Marshal in UpdateProblem func")
	}
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DisplayProblems(w http.ResponseWriter, r *http.Request) {

	tf, err := template.ParseFiles("../templates/index.html")
	if err != nil {
		tf, _ = template.New("index").Parse(`<html><body><h1>NO TEMPLATE.</h1></body></html>`)
	}
	err = tf.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
