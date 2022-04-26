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
git clone https://github.com/Ax7-cmd/golang-restAPI-JWT.git
cd golang-restAPI-JWT
```

* Build mod Vendor
```sh
# make sure you have folder vendor in your root directory "golang-restAPI-JWT/vendor"
# if you dont have folder vendor create new one with this command 
mkdir vendor

# install golang package in your vendor
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
    
   >router code implement
   ```go
    router.POST("/vaonline/rest/json/gettoken", func(c *gin.Context) {
		// using BindJson method to serialize body with struct
		if err := c.BindJSON(&gt_req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			gt_req = structs.GetTokenRequest{}
			return
		}
		if err := validate.Struct(gt_req); err != nil {
			errs := validator.ToErrResponse(err, trans)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errs,
			})
			gt_req = structs.GetTokenRequest{}
			return
		}
		
		gt_res = vaonline.GenerateToken(gt_req)
		c.JSON(http.StatusOK,&gt_res)
		gt_req = structs.GetTokenRequest{}
	})
   ```

**GRPC**
>Implement [grpc](https://grpc.io) function, A high performance, open-source universal RPC framework

1. How to start grpc whith go?
   [GRPC-go-guideline](https://grpc.io/docs/quickstart/go.html)

2. How to enable grpc?
   > Go to server.go, you will saw below sample code, if don`t want enable it, you can comment it.
   ```go
   go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", 10000))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		var opts []grpc.ServerOption
		grpcServer := grpc.NewServer(opts...)
		pb.RegisterRouteGuideServer(grpcServer, mgrpc.NewServer())
		grpcServer.Serve(lis)
	}()
   ```
3. About grpc here
   For now, It`s just a sample code, and export an api to testing. 

4. GRPC-go environment by docker
    If you want a generate grpc code env, you can go to [here](https://hub.docker.com/r/nightlegend/grpc-go/).


---
## Authors

* **Ax7-cmd** - *Initial work* - [Ax7](https://github.com/Ax7-cmd).
* *Backend dev*