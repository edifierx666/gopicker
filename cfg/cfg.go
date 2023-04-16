package cfg

type Cfg struct {
  Name  string `json:"name"`
  Limit int    `json:"limit"`
  Open  bool   `json:"open"`
  Goget bool   `json:"goget"`
}
