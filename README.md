# [Merchant Api Gateway]:

Its clone form [nightlegend/apigateway](https://github.com/nightlegend/apigateway). Build with golang and gin. It features a simple and better performance, and customize with requirements needed.

<br>

<h1>Design</h1>

![Gopher image](Doc/apigateway.jpg)

<h1>How to run ?</h1>

<h2>Start APIGATEWAY</h2>

* Init workdir
```sh
git clone git@github.com:Nomina-VIp/merchant-api-gateway.git
cd merchant-api-gateway
```

* Copy .env.example to .env
```sh
cp .env.example .env
# change default config .env with your local config 
```

* Start APIGATEWAY
```sh
# start with default
go run server.go
# -env: current for enable/disable debug model.
go run server.go -env development
```


If running normally, you can access <a href="http://localhost:8080">http://localhost:8080</a>

**Application details**

---

1. Server starting from server.go
   ```go
    package main
    import (
        "fmt"
        "net"

        mgrpc "merchant-api-gateway/Core/Grpc"
        log "github.com/Sirupsen/logrus"
        pb "merchant-api-gateway/Core/Grpc/Services"
        
        "merchant-api-gateway/Config"
        "merchant-api-gateway/Core/Router"
        "google.golang.org/grpc"
    )

    // Api server start from here. router is define your api router and public it.
    func main() {
        // GRPC
        // Here will enable grpc server, if you don`t want it, you can disable it
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
        app_env := config.GoDotEnvVariable("APP_ENV")

        // HTPP
        // start api server, *env is what`s environment will running, currentlly this only for enable or disable debug modle
        // After may be use it load different varible.
        router.Start(app_env)
    }
   ```
2. Project code structure

    | folder        | content                                   |
    | ------------- |:-------------                             |
    | conf          | put some application configure to here    |
    | core          | put core sources to here(api, grpc, etc)  |
    | database      | here database connection and globar var   |
    | middleware    | put middleware code to here, like cors    |
    | logs          | here will save console log                |
    | vendor        | here is save third party                  |
    | doc           | here is save some document about project  |

3. Router define
   
   Router, It`s like your application gate, help you dispatch all request to target service.
   >go to apigateway/core/router/router.go, you can define your router.

    ```go
    func Start(env string) {
        // enable debug/release mode
        switch env {
        case "development":
            gin.SetMode(gin.DebugMode)
        default:
            gin.SetMode(gin.ReleaseMode)
            fmt.Printf("Start prod mode...\nServer listen on: %v", LisAddr)
        }

        router := gin.New()
        router.Use(gin.Logger())
        router.Use(gin.Recovery())
        router.Use(middleware.CORSMiddleware())
        router.Use(middleware.LoggerApp())
        //No Permission Validation
        public.APIRouter(router)

        router.Run(LisAddr)
    }
    ```
4. Example Api
   > gettoken api: http://localhost:8080/vaonline/rest/json/gettoken

   ```text
    request:

    POST /login HTTP/1.1
    Host: localhost:8080
    Content-Type: application/json
    Cache-Control: no-cache
    Postman-Token: a70f71a7-72b9-4106-9bcd-fd2b65be1e87

    {
        "merchantId": "001",
        "merchantRefCode": "JS008sKs",
        "secureCode": "e524eea49aa125b01cbfe7f7c3bebd7b257b64abd3dc95288c3740a20c04ffc9"
    }
    
    response:
    {
        "merchantId": "001",
        "merchantRefCode":"JS008sKs",
        "token": "RTkwQjk2QzVGQUM4NDIwQzYxMDVCNDI4QUFCNTNGRkEwRkJCNDBEODA4NEIxOUQ1MTc1NjcyMTFGNDBCNUVBOQ==”,
        "responseCode": "200",
        "responseMsg": "Success generate token"
    }
   ```
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