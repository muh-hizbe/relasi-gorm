package models

type Tag struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
}

type TagRequest struct {
	Name string `json:"name" gorm:"not null"`
}

type TagResponseWithPost struct {
	ID    int                        `json:"id"`
	Name  string                     `json:"name"`
	Posts []PostResponseHiddenUserID `json:"posts" gorm:"many2many:post_tags;ForeignKey:ID;joinForeignKey:TagID;References:ID;joinReferences:PostID"`
}

func (TagResponseWithPost) TableName() string {
	return "tags"
}
