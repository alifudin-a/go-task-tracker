package helpers

// Response custom response struct
type Response struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message,omitempty"`
	Body    map[string]interface{} `json:"body,omitempty"`
}
