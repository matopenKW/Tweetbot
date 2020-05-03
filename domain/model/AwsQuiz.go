package model

type AwsQuiz struct {
	SeqNo  string `csv:"No"`
	Quiz   string `csv:"Quiz"`
	Choice string `csv:"Choice"`
	Answer string `csv:"Ansewr"`
}
