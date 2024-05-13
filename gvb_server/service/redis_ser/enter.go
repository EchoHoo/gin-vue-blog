package redis_ser

const (
	articleDiggPrefix    = "article_digg"
	articleLookPrefix    = "article_look"
	articleCommentPrefix = "article_comment_count"
	commentDiggPrefix    = "comment_digg"
)

func NewDigg() CountDB {
	return CountDB{
		Index: articleDiggPrefix,
	}
}
func NewLook() CountDB {
	return CountDB{
		Index: articleLookPrefix,
	}
}
func NewCommentCount() CountDB {
	return CountDB{
		Index: articleCommentPrefix,
	}
}
func NewCommentDigg() CountDB {
	return CountDB{
		Index: commentDiggPrefix,
	}
}
