why go build
check package <name>

## when connection is successfully worked between golang GORM and MySQL:

"mysql", "yutah:Yutagenta55@(localhost:3306)/bookinfo?charset=utf8&parseTime=True&loc=Local"

mysql> show tables

mysql> select * from books

mysql.server stop mysql.server start

password Yutagenta55
SHOW GRANTS FOR 'yutah'@'localhost'

long story short:

panic: dial tcp 127.0.0.1:3306: connect: connection refused

goroutine 1 [running]:
github.com/YutaUtah/RESTAPIBook/pkg/config.Connect(...)
        /Users/yutahayashi/go/src/github.com/YutaUtah/RESTAPIBook/pkg/config/app.go:17
github.com/YutaUtah/RESTAPIBook/pkg/models.init.0()
        /Users/yutahayashi/go/src/github.com/YutaUtah/RESTAPIBook/pkg/models/book.go:18 +0xe5
exit status 2

this was access issue to yutah user