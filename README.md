# Leetcode Record RESTful API

The project builds upon my personal leetcode problems practice history, and reflect it in simple html page.


## Introduction:

This project is initiated because I want to learn more about RESTful API and programming language Go. At the same time, I want to keep track of my leetcode change record by rendering data in simple html file and visualize my progress :)

## Overview:

The data is stored in JSON format. Below is the example:

```
record = {
            "ID": 47,
            "number": 122,
            "title": "String Manupulation",
            "difficulty": "easy",
            "date": "2022-01-10"
        }
```

- **ID**: This is an unique identifier that represents the problem
- **number**: This is a problem number that you can see in leetcode website
- **title**: This is a problem title that you can see in leetcode website
- **difficulty**: This represents problem level. (Easy, Medium, Hard)
- **date**: This is a date of my problem completion



## Request Information:

#### GET Request all problems:

```
curl -i http://localhost:9000/api/v1/problems/
```

#### GET Request problem by ID:

```
curl --header "Content-Type: application/json"    --request GET   http://localhost:9000/api/v1/problems/{id}
```

#### POST Request problem by ID:

```
curl --header "Content-Type: application/json"    --request POST   http://localhost:9000/api/v1/problems/
```

#### DELETE Request problem by ID:

```
curl --header "Content-Type: application/json"    --request DEKETE   http://localhost:9000/api/v1/problems/{id}
```

## Visualizer:

```http://localhost:{portid}}/``` will take you to the table with the history of leetcode problems you solved.


### Database:

- MySQL, all my config is stored in config directory - app.go

## Local Setup for Development:

### clone the repo locally

```
git clone https://github.com/YutaUtah/LeetCodeAPI.git
```

### Download the required Go modules
```
go mod download
```

### Run the server
```
cd cmd
go run main.go
```