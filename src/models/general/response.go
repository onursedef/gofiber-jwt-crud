package general

type Response struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Kind    string      `json:"type"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

type ErrorResponse struct {
	Request_url string `json:"request_url"`
	Title       string `json:"title"`
	Detail      string `json:"detail"`
}

type PaginationResponse struct {
	Page    int         `json:"page"`
	PerPage int         `json:"per_page"`
	Total   int         `json:"total"`
	Items   interface{} `json:"items"`
}
