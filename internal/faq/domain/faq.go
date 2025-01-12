package domain

type QuestionID uint8

type FrequentQuestion struct {
	QuestionID QuestionID
	Question   string
	Answer     string
}
