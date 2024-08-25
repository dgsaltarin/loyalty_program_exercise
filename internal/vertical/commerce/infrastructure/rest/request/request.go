package request

type CreateCommerceRequest struct {
	Name     string                `json:"name"`
	Address  string                `json:"address"`
	Branches []CreateBranchRequest `json:"branches"`
}

type CreateBranchRequest struct {
	Name       string `json:"name"`
	Address    string `json:"address"`
	CommerceID string `json:"commerce_id"`
}
