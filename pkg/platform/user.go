package platform

import (
	"github.com/Luke-Vear/nettaton/pkg/subnet"
)

// User struct contains all data about a user.
type User struct {
	UserID   string                    `json:"userID"`
	Password string                    `json:"password"`
	Email    string                    `json:"email"`
	Status   string                    `json:"status"`
	Scores   map[string]*QuestionScore `json:"scores"`
}

// QuestionScore tracks correct answers and overall attempts for a question kind.
type QuestionScore struct {
	Attempts int
	Correct  int
}

// NewUser returns a fully initialized *User.
func NewUser() *User {

	scores := make(map[string]*QuestionScore)

	// Loop over all question types and initalize zero values.
	for k := range subnet.QuestionFuncMap {
		scores[k] = &QuestionScore{}
	}

	return &User{Scores: scores}
}
