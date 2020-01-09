# Golang API

## Golang API with database MySQL

### Export database
```
create  database people;
use people;

create  table peoples (
	id int  not  null auto_increment,
	nama varchar(20),
	gender varchar(1),
	primary  key (id)
);

insert  into peoples (nama, gender) values
("Erwindo", "M"),
("Rifaldo", "M");

select * from peoples;
```

### Runing server
```
$ go run .
```
### Open Postman

#### Get all students
Method `GET` and URL `http://localhost:8080/students` will get all students from database.

#### Insert data student
Method `POST` and URL `http://localhost:8080/student` with the following body:
```
{
	"Nama" : "Student Name",
	"Gender" : "M"
}
```
will insert data to database.

#### Get data student by ID
Method `GET` and URL `http://localhost:8080/student` with the parameter URL `?id=[STUDENT_ID]` will get one student with the inserted paramater ID

*Example:*
```
Method: GET | URL: http://localhost:8080/student?id=10 
```

#### Update data student by ID
Method `PUT` and URL `http://localhost:8080/student` with the parameter URL `?id=[STUDENT_ID]&nama=[STUDENT_NAMA]&gender=[STUDENT_GENDER]` will update data student where student id same with parameter `ID`

*Example*
```
Method: PUT | URL: http://localhost:8080/student?id=12&nama=Prana&gender=M
```

#### Delete data student by ID
Method `DELETE` and URL `http://localhost:8080/student` with the parameter URL `?id=[STUDENT_ID]` will delete data student where student id same with parameter `ID`