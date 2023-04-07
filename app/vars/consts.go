package vars

const (
	ConfigKeyPrefix = "MOBGI_GAME"
	TemplateExt     = "tmpl"
	HotGame         = "Hot"
	NewGame         = "New"
	NormalGame      = "Normal"
)

type Game struct {
	GameId      string
	Name        string
	GameUrl     string
	Icon        string
	Description string
	Tag         string
	Star        uint8
}
