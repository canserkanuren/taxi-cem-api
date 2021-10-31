package models

type MailBody struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	From      string `json:"from" binding:"required"`
	Subject   string `json:"subject" binding:"required"`
	Content   string `json:"content" binding:"required"`
}
