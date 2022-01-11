## LEETCODE PROBLEM API

This project is initiated to learn more about RESTful API and programming language Go. At the same time, I want to keep track of my leetcode change record by rendering data in simple html file.

### How to send GET Request:

##### GET Request all problems:

```
curl -i http://localhost:9000/problems/
```

##### GET Request problem by ID:

```
curl --header "Content-Type: application/json"   --request GET   http://localhost:9000/problem/{id}
```

