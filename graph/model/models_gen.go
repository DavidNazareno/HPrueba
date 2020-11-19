// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Client struct {
	ID     string `json:"id"`
	Ticket string `json:"ticket"`
}

type NewRequest struct {
	Token  string `json:"token"`
	Client int    `json:"client"`
	Status int    `json:"status"`
	Score  int    `json:"score"`
}

type Order struct {
	ID         string      `json:"id"`
	Technician *Technician `json:"technician"`
	Token      string      `json:"token"`
}

type Request struct {
	Token   string  `json:"token"`
	Clients *Client `json:"clients"`
	Status  int     `json:"status"`
	Score   *int    `json:"score"`
}

type Technician struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}