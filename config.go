package mergemap

type Config map[string]FieldConfig

type FieldConfig struct {
	Strategy mergeStrategy
}
