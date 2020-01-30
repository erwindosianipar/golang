# Learn Golang JWT (JSON Web Token)


## Package

```
$ go get github.com/dgrijalva/jwt-go
```

## API

### Sign Up

| Method | URL |
| ------ | --- |
| POST | http://localhost:8000/signup |

**Request Body**

```
{
	"email" : "YOUR_EMAIL_ADDRESS",
	"password" : "YOUR_PASSWORD"
}
```

### Login

| Method | URL |
| ------ | --- |
| POST | http://localhost:8000/login |

**Request Body**

```
{
	"email" : "YOUR_EMAIL_ADDRESS",
	"password" : "YOUR_PASSWORD"
}
```

**Response Body**

```
{
    "token": "THIS_IS_YOUR_TOKEN"
}
```

### Request URL

| Method | URL |
| ------ | --- |
| GET | http://localhost:8000/user |

*Attach a token with type **bearer** in authorization.*