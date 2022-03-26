package filter

import "github.com/core-go/search"

type UserFilter struct {
	*search.Filter
	Id          string            `json:"id" gorm:"column:id;primary_key" bson:"_id" dynamodbav:"id" firestore:"id" avro:"id" validate:"required,max=40" match:"equal"`
	Username    string            `json:"username" gorm:"column:username" bson:"username" dynamodbav:"username" firestore:"username" avro:"username" validate:"required,username,max=100" match:"prefix" q:"prefix"`
	Email       string            `json:"email" gorm:"column:email" bson:"email" dynamodbav:"email" firestore:"email" avro:"email" validate:"email,max=100" match:"prefix" q:"prefix"`
	Phone       string            `json:"phone" gorm:"column:phone" bson:"phone" dynamodbav:"phone" firestore:"phone" avro:"phone" validate:"required,phone,max=18" q:"true"`
	DateOfBirth *search.TimeRange `json:"dateOfBirth" gorm:"column:date_of_birth" bson:"dateOfBirth" dynamodbav:"dateOfBirth" firestore:"dateOfBirth" avro:"dateOfBirth"`
}
