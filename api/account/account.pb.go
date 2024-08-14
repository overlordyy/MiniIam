package account

type Account struct {
	Id       int    `json:"id"`
	UserId   string `json:"userId" binding:"required" db:"user_id"`
	UserName string `json:"userName" binding:"required" db:"user_name"`
	Passwd   string `json:"passwd" binding:"required" db:"passwd"`
	Status   int    `json:"status" binding:"required" db:"status"`
	GroupId  int    `json:"groupId" binding:"required" db:"group_id"`
}
