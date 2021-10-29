package types

type UpdateMessage struct{
	ID int `json:"id"` 
	Todo string `json:"todo"`
	Description string `json:"description"`
}