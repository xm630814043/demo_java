package moudle

import (
	"github.com/olivere/elastic"
	"time"
)

type Tweet struct {
	User     string                `json:"user"`
	Age      int                   `json:"age"`
	Message  string                `json:"message"`
	Retweets int                   `json:"retweets"`
	Image    string                `json:"image,omitempty"`
	Created  time.Time             `json:"created,omitempty"`
	Tags     []string              `json:"tags,omitempty"`
	Location string                `json:"location,omitempty"`
	Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
}





