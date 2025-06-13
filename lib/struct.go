package lib

type Response struct {
	Success  bool        `json:"succes"`
	Message  string      `json:"message"`
	PageInfo any         `json:"pageInfo,omitempty"`
	Results  interface{} `json:"results,omitempty"`
}
