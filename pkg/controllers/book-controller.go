package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/YutaUtah/LeetCodeAPI/pkg/models"
	"github.com/YutaUtah/LeetCodeAPI/pkg/utils"
	"github.com/gorilla/mux"
)

var newProblems models.Problem

func GetProblem(w http.ResponseWriter, r *http.Request) {
	// store all problems info to newProblems and send it back to the user
	newProblems := models.GetAllProblems()
	// convert newProblems to json object
	res, _ := json.Marshal(newProblems)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetProblemById(w http.ResponseWriter, r *http.Request) {
	// get bars (json object)
	vars := mux.Vars(r)
	// get problemId by acccessing problemId attribute
	problemId := vars["problemId"]
	ID, err := strconv.ParseInt(problemId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	problemDetails, _ := models.GetProblemById(ID)
	res, _ := json.Marshal(problemDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateProblem(w http.ResponseWriter, r *http.Request) {
	CreateProblem := &models.Problem{}
	utils.ParseBody(r, CreateProblem)
	b := CreateProblem.CreateProblem()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteProblem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	problemId := vars["problemId"]
	ID, err := strconv.ParseInt(problemId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	problem := models.DeleteProblem(ID)
	res, _ := json.Marshal(problem)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateProblem(w http.ResponseWriter, r *http.Request) {
	var updateProblem = &models.Problem{}
	utils.ParseBody(r, updateProblem)
	vars := mux.Vars(r)
	problemId := vars["problemId"]
	ID, err := strconv.ParseInt(problemId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	problemDetails, db := models.GetProblemById(ID)
	if updateProblem.Name != "" {
		problemDetails.Name = updateProblem.Name
	}
	if updateProblem.Author != "" {
		problemDetails.Author = updateProblem.Author
	}
	if updateProblem.Publication != "" {
		problemDetails.Publication = updateProblem.Publication
	}
	// saving problem details in database
	db.Save(&problemDetails)
	res, _ := json.Marshal(problemDetails)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
