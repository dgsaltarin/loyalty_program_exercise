package response

type CreateCommerceResponse struct {
	ID       string                 `json:"id"`
	Name     string                 `json:"name"`
	Address  string                 `json:"address"`
	Branches []CreateBranchResponse `json:"branches"`
}

type CreateBranchResponse struct {
	ID         string `json:"id"`
	CommerceID string `json:"commerce_id"`
	Name       string `json:"name"`
	Address    string `json:"address"`
}
