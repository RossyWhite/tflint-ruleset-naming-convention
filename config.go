package configs

type Config struct {
	Rules []*Rule `json:"rules"`
}

type Rule struct {
	Resource  string `json:"resource"`
	Attribute string `json:"Attribute"`
	Regex     string `json:"regex"`
}
