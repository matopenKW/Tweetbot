package handler

import (
	"GetQuiz/domain/model"
	"GetQuiz/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

func AwsQuiz(ctx *gin.Context) {

	sqlCon, err := util.GetConnection()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusHTTPVersionNotSupported, "error")
		return
	}
	defer sqlCon.Close()

	quizList, err := selectQuiz(sqlCon)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusHTTPVersionNotSupported, "error")
		return
	}
	ctx.JSON(http.StatusOK, quizList)
}

func selectQuiz(sqlCon *gorm.DB) (*model.AwsQuiz, error) {

	quizList := []*model.AwsQuiz{}
	err := sqlCon.Table("AWS_QUIZ").Find(&quizList).Error
	if err != nil {
		return nil, err
	}

	return quizList[0], nil

}

// 本日のSeqNoを取得
func getSeqNo() int64 {
	return 1
}
