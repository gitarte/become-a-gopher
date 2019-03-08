package main

// File -
type File struct {
	Index int    `json:"index"`
	Name  string `json:"name"`
}

// Result -
type Result struct {
	Result []File `json:"result"`
}

// Account -
type Account struct {
	ID       int64  `json:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email"    binding:"required"`
}

// AccountResult -
type AccountResult struct {
	Result []Account `json:"result"`
}
