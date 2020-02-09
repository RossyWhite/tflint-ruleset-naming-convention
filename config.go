package main

// Config is plugin configuration
type Config struct {
	Rules []*Rule `json:"rules"`
}

// Rule is a configuration of each rule
type Rule struct {
	Resource  string `json:"resource"`
	Attribute string `json:"Attribute"`
	Regex     string `json:"regex"`
}
