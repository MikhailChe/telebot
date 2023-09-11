module github.com/mikhailche/telebot

go 1.13

require (
	github.com/goccy/go-yaml v1.9.5
	github.com/spf13/viper v1.13.0
	github.com/stretchr/testify v1.8.0
	gopkg.in/telebot.v3 v3.0.0-00010101000000-000000000000
)

replace gopkg.in/telebot.v3 => ./

exclude gopkg.in/telebot.v3 v3.1.3
