package models

type Question struct {
	ID                  interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	QuestionId          int         `bson:"QuestionId,omitempty" json:"QuestionId,omitempty"`
	QuestionUser        string      `bson:"QuestionUser,omitempty" json:"QuestionUser,omitempty"`
	QuestionAnswer      Answer      `bson:"QuestionAnswer,omitempty" json:"QuestionAnswer,omitempty"`
	QuestionTitle       string      `bson:"QuestionTitle,omitempty" json:"QuestionTitle,omitempty"`
	QuestionDescription string      `bson:"QuestionDescription,omitempty" json:"QuestionDescription,omitempty"`
}

type Answer struct {
	AnswerId          int    `bson:"AnswerId,omitempty" json:"UserId,omitempty"`
	AnswerUser        string `bson:"AnswerUser,omitempty" json:"AnswerUser,omitempty"`
	AnswerDescription string `bson:"AnswerDescription,omitempty" json:"AnswerDescription,omitempty"`
}

type QuestionUpdate struct {
	ModifiedCount int64 `json:"modifiedCount"`
	Result        Question
}

type QuestionDelete struct {
	DeletedCount int64 `json:"deletedCount"`
}
