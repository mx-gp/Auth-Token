package helper

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	lumberjack "gopkg.in/natefinch/lumberjack.v2"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var log *logrus.Logger

// HandleError: Error Handler
func HandleError(c *gin.Context, err error, context string, code int) {
	erro, _ := err.(*pq.Error)
	if erro == gorm.ErrRecordNotFound {
		c.JSON(400, gin.H{
			"code":    400,
			"error":   err,
			"message": context,
			"success": false,
		})
	} else {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Errorln(context)
		c.JSON(code, gin.H{
			"code":    code,
			"error":   err,
			"message": context,
			"success": false,
		})
	}
}

// Generate Random string with 6 to 12 characters
func GenerateRandomString() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var b []rune

	rand.Seed(time.Now().UnixNano())

	minLength := 6
	maxLength := 12

	// Generate random integer till 12
	randomInt := minLength + rand.Intn(maxLength-minLength)

	b = make([]rune, randomInt)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

type IP struct {
	Query string
}

// GET IP Address from Third Party Website
func GetIP() string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		HandleError(nil, err, "Can't able to fetch data", 500)
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		HandleError(nil, err, "Unable to read response", 500)
	}

	var ip IP
	json.Unmarshal(body, &ip)

	return ip.Query
}

func init() {
	log = NewLogger()

	log.SetOutput(&lumberjack.Logger{
		Filename:   "authtoken.log",
		MaxSize:    500,
		MaxBackups: 7,
		MaxAge:     3,
		Compress:   true,
	})
}

func NewLogger() *logrus.Logger {
	log := logrus.New()
	log.Level = logrus.DebugLevel
	log.Formatter = &logrus.TextFormatter{}

	return log
}
