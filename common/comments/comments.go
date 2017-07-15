package comments

import "gitlab.com/sunshower.io/updraft/common/observer"


func CommentScanned(comment string) observer.Message {
	return &observer.BaseEvent{
		Topic: observer.COMMENT,
		Body:  comment,
	}
}
