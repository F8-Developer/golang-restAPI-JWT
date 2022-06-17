# [Golang RestAPI with JWT]:

Build with golang and gin with jwt authentication. It features a simple and better performance, and customize with requirements needed.

### Required

 * GO 1.17.5 - [go1.17.5](https://go.dev/doc/devel/release#go1.17).

### Using

- Gin Web Framework 1.3.0 - [Gin-Gionic](https://github.com/gin-gonic/gin)
- MySQL 5.7.26 - [MySQL](https://dev.mysql.com/doc/relnotes/mysql/5.7/en/news-5-7-26.html)
- Go Validator v10 - [go-validator](https://github.com/go-playground/validator)

<br>

<h2>Installation</h2>

* Init workdir
```sh
git clone https://github.com/F8-Developer/golang-restAPI-JWT.git
cd golang-restAPI-JWT
```

* Build mod Vendor
```sh
# make sure you have folder vendor in your root directory "golang-restAPI-JWT/vendor"
# if you dont have folder vendor create new one with this command 
mkdir vendor

# install golang package in folder vendor
go mod vendor
```

* Copy .env.example to .env
```sh
cp .env.example .env
# change default config .env with your local config
```

* Database Note
```sh
# restAPI will automatically migrate when there is no table in you database
```

* Start restAPI 
```sh
# start with default
go run server.go
```

If running normally, you can access <a href="http://localhost:8080">http://localhost:8080</a>

---
<h2>Rest API</h2>

1. Endpoint

    | METHOD | URL                     | INFO                                              |
    | ------ | :-------------          | :-------------                                    |
    | GET    | /                       | index pa                                          |
    | POST   | /register               | for create user                                   |
    | POST   | /login                  | login user and generate jwt token                 |
    | POST   | /secure/category/list   | check all category in database                    |
    | POST   | /secure/product/list    | check all product in database                     |
    | POST   | /secure/product/detail  | check product detail by product id                |
    | POST   | /secure/cart/add        | add product to cart                               |
    | POST   | /secure/cart/list       | check user cart (get user from jwt claims email)  |

2. Example Api
   > register api : http://localhost:8080/register

    ```text
    request:

    POST /register HTTP/1.1
    Host: localhost:8080
    Content-Type: application/json
    Content-Length: 111

    {
        "name": "faishal amrullah",
        "email": "c.faishal.amrullah@gmail.com",
        "password": "faishal123$#"
    }
    
    response:
    {
        "responseCode": 200,
        "responseMsg": "User successfully register"
    }
    ```
    ![Gopher image](Doc/register.png)
    <br>
    <br>
    <br>
    > login api : http://localhost:8080/login

    ```text
    request:

    POST /login HTTP/1.1
    Host: localhost:8080
    Content-Type: application/json
    Content-Length: 81

    {
        "email" : "c.faishal.amrullah@gmail.com",
        "password" : "faishal123$#"
    }
    
    response:
    {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiZmFpc2hhbCBhbXJ1bGxhaCIsImVtYWlsIjoiYy5mYWlzaGFsLmFtcnVsbGFoQGdtYWlsLmNvbSIsImV4cCI6MTY1MTAxMjU5MX0.Sv1hXav7BbwjIEo2aY6MQb8oPD11bB9cq00TQxvyhLk",
        "responseCode": 200,
        "responseMsg": "User successfully login"
    }
    ```
    ![Gopher image](Doc/login.png)
    <br>
    <br>
    <br>
    > category list api : http://localhost:8080/secure/category/list

    ```text
    request:

    POST /secure/category/list HTTP/1.1
    Host: localhost:8080
    Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiZmFpc2hhbCBhbXJ1bGxhaCIsImVtYWlsIjoiYy5mYWlzaGFsLmFtcnVsbGFoQGdtYWlsLmNvbSIsImV4cCI6MTY1MTAxMjU5MX0.Sv1hXav7BbwjIEo2aY6MQb8oPD11bB9cq00TQxvyhLk
    
    response:
    {
        "categories": [
            {
                "ID": 1,
                "Name": "buah",
                "Descriptions": "buah"
            },
            {
                "ID": 2,
                "Name": "sayur",
                "Descriptions": "sayur"
            },
            {
                "ID": 3,
                "Name": "rempah-rempah",
                "Descriptions": "rempah-rempah"
            }
        ],
        "responseCode": 200,
        "responseMsg": "List Categories successful"
    }
    ```
    ![Gopher image](Doc/category-list.png)
    <br>
    <br>
    <br>
    > product list api : http://localhost:8080/secure/product/list

    ```text
    request:

    POST /secure/product/list HTTP/1.1
    Host: localhost:8080
    Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiZmFpc2hhbCBhbXJ1bGxhaCIsImVtYWlsIjoiYy5mYWlzaGFsLmFtcnVsbGFoQGdtYWlsLmNvbSIsImV4cCI6MTY1MTAxMjU5MX0.Sv1hXav7BbwjIEo2aY6MQb8oPD11bB9cq00TQxvyhLk
    
    response:
    {
        "products": [
            {
                "ID": 1,
                "Categories": [
                    {
                        "ID": 1,
                        "Name": "buah",
                        "Descriptions": "buah"
                    }
                ],
                "Name": "Anggur Merah Globe",
                "Descriptions": "Anggur Merah Globe (~0.65 kg)",
                "Quantity": 88,
                "Price": 56200
            },
            {
                "ID": 2,
                "Categories": [
                    {
                        "ID": 1,
                        "Name": "buah",
                        "Descriptions": "buah"
                    }
                ],
                "Name": "Apel Fuji Jingle",
                "Descriptions": "Apel Fuji Jingle (~0.3 kg)",
                "Quantity": 100,
                "Price": 15000
            },
            {
                "ID": 3,
                "Categories": [
                    {
                        "ID": 2,
                        "Name": "sayur",
                        "Descriptions": "sayur"
                    }
                ],
                "Name": "Wortel",
                "Descriptions": "Wortel (~0.1 kg)",
                "Quantity": 100,
                "Price": 1700
            },
            {
                "ID": 4,
                "Categories": [
                    {
                        "ID": 2,
                        "Name": "sayur",
                        "Descriptions": "sayur"
                    }
                ],
                "Name": "Kol Putih",
                "Descriptions": "Kol Putih (~1 kg)",
                "Quantity": 100,
                "Price": 18000
            },
            {
                "ID": 5,
                "Categories": [
                    {
                        "ID": 3,
                        "Name": "rempah-rempah",
                        "Descriptions": "rempah-rempah"
                    }
                ],
                "Name": "Lengkuas",
                "Descriptions": "Lengkuas (~0.2 kg)",
                "Quantity": 100,
                "Price": 8000
            },
            {
                "ID": 6,
                "Categories": [
                    {
                        "ID": 3,
                        "Name": "rempah-rempah",
                        "Descriptions": "rempah-rempah"
                    }
                ],
                "Name": "Jahe",
                "Descriptions": "Jahe (~0.2 kg)",
                "Quantity": 100,
                "Price": 10800
            }
        ],
        "responseCode": 200,
        "responseMsg": "List Products successful"
    }
    ```
    <br>
    <br>
    
    > product detail api : http://localhost:8080/secure/product/detail

    ```text
    request:

    POST /secure/product/detail HTTP/1.1
    Host: localhost:8080
    Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiZmFpc2hhbCBhbXJ1bGxhaCIsImVtYWlsIjoiYy5mYWlzaGFsLmFtcnVsbGFoQGdtYWlsLmNvbSIsImV4cCI6MTY1MTAxMjU5MX0.Sv1hXav7BbwjIEo2aY6MQb8oPD11bB9cq00TQxvyhLk
    Content-Type: application/json
    Content-Length: 23

    {
        "productID" : 3
    }
    
    response:
    {
        "product": {
            "ID": 3,
            "Categories": [
                {
                    "ID": 2,
                    "Name": "sayur",
                    "Descriptions": "sayur"
                }
            ],
            "Name": "Wortel",
            "Descriptions": "Wortel (~0.1 kg)",
            "Quantity": 100,
            "Price": 1700
        },
        "responseCode": 200,
        "responseMsg": "Product detail successful"
    }
    ```
    ![Gopher image](Doc/product-detail.png)
    <br>
    <br>
    <br>
    > cart add api : http://localhost:8080/secure/cart/add

    ```text
    request:

    POST /secure/cart/add HTTP/1.1
    Host: localhost:8080
    Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiZmFpc2hhbCBhbXJ1bGxhaCIsImVtYWlsIjoiYy5mYWlzaGFsLmFtcnVsbGFoQGdtYWlsLmNvbSIsImV4cCI6MTY1MTAxMjU5MX0.Sv1hXav7BbwjIEo2aY6MQb8oPD11bB9cq00TQxvyhLk
    Content-Type: application/json
    Content-Length: 44

    {
        "productID" : 3,
        "quantity" : 10
    }
    
    response:
    {
        "responseCode": 200,
        "responseMsg": "Product successfully add to cart"
    }
    ```
    ![Gopher image](Doc/cart-add.png)
    <br>
    <br>
    <br>
    > cart list api : http://localhost:8080/secure/cart/list

    ```text
    request:

    POST /secure/cart/list HTTP/1.1
    Host: localhost:8080
    Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiZmFpc2hhbCBhbXJ1bGxhaCIsImVtYWlsIjoiYy5mYWlzaGFsLmFtcnVsbGFoQGdtYWlsLmNvbSIsImV4cCI6MTY1MTAxMjU5MX0.Sv1hXav7BbwjIEo2aY6MQb8oPD11bB9cq00TQxvyhLk
    
    response:
    {
        "cart": [
            {
                "ID": 1,
                "User": {
                    "Name": "faishal amrullah",
                    "Email": "c.faishal.amrullah@gmail.com"
                },
                "Product": {
                    "ID": 3,
                    "Categories": [
                        {
                            "ID": 2,
                            "Name": "sayur",
                            "Descriptions": "sayur"
                        }
                    ],
                    "Name": "Wortel",
                    "Descriptions": "Wortel (~0.1 kg)",
                    "Quantity": 90,
                    "Price": 1700
                },
                "Quantity": 10,
                "Price": 1700,
                "Total": 17000
            }
        ],
        "responseCode": 200,
        "responseMsg": "List Cart successful"
    }
    ```
    ![Gopher image](Doc/cart-list.png)
    <br>
    <br>
    <br>
---
<h2>Validation</h2>

Add validation for /register, /login and /secure/cart/add for better experience

---
## Authors

* **Ax7-cmd** - *Initial work* - [Ax7](https://github.com/Ax7-cmd).
