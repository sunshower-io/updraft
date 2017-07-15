package comments

import "github.com/sunshower-io/updraft/common/observer"


func CommentScanned(comment string) observer.Message {
	return &observer.BaseEvent{
		Topic: observer.COMMENT,
		Body:  comment,
	}
}
