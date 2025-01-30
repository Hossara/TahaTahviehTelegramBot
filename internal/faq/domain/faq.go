package domain

type QuestionID int64

type FrequentQuestion struct {
	QuestionID QuestionID
	Question   string
	Answer     string
}
