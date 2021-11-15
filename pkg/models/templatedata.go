package models

//TemplateData hold data sent from holder to template
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	float32   map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
}
