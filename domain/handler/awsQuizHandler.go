package handler

import (
	"GetQuiz/domain/model"
	"GetQuiz/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func AwsQuiz(ctx *gin.Context) {

	sqlCon, err := util.GetConnection()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusHTTPVersionNotSupported, "error")
		return
	}
	defer sqlCon.Close()

	quiz, err := selectQuiz(sqlCon)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusHTTPVersionNotSupported, "error")
		return
	}
	ctx.JSON(http.StatusOK, quiz)
}

func selectQuiz(sqlCon *gorm.DB) (*model.AwsQuiz, error) {

	cnt := 0
	err := sqlCon.Table("AWS_QUIZ").Count(&cnt).Error
	if err != nil {
		return nil, err
	}

	rand.Seed(time.Now().UnixNano())
	min := 1
	rnd := rand.Intn(cnt) + min

	quiz := &model.AwsQuiz{}
	err = sqlCon.Table("AWS_QUIZ").Find(&quiz, "seq_no=?", rnd).Error

	return quiz, nil
}

// 本日のSeqNoを取得
func getSeqNo() int64 {
	return 1
}
