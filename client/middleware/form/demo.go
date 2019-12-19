package form

type Demo struct {
	Name string `form:"name" json:"name" binding:"required"`
}


