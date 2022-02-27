package Middleware

import (
    "time"
	"os"
	"io"
	"io/ioutil"
	"bytes"
	"strings"

    log "github.com/Sirupsen/logrus"

    "github.com/gin-gonic/gin"
)

var buf bytes.Buffer

//Log to file App
func LoggerApp() gin.HandlerFunc {
	// set golable logs file path.
	execDirAbsPath, _ := os.Getwd()
	f, err := os.OpenFile(execDirAbsPath+"/logs/app.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//set output of logs to file
	log.SetOutput(f)
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp: true,
	})
	
    return func(c *gin.Context){
        //Start time
        startTime := time.Now()

        //Process request
		buf, _ := ioutil.ReadAll(c.Request.Body)
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
		c.Request.Body = rdr2
        c.Next()

        //End time
        endTime := time.Now()

        //Execution time
        latencyTime := endTime.Sub(startTime)

        //Request method
        reqMethod := c.Request.Method

        //Request routing
        reqUri := c.Request.RequestURI

        //Request body
        reqBody := readBody(rdr1)
		reqBody = strings.Replace(reqBody, "\n", "", -1)

        // status code
        statusCode := c.Writer.Status()

        // request IP
        clientIP := c.ClientIP()

        //Log format
        log.Infof("| %3d | %13v | %15s | %s | %s | %s |",
            statusCode,
            latencyTime,
            clientIP,
            reqMethod,
            reqUri,
            reqBody,
        )
    }
}

func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	s := buf.String()
	return s
}