package zap

import "strings"

type Reports struct {
	c *Client
}

// This component is optional and therefore the API will only work if it is installed
func (r Reports) Templates() (map[string]interface{}, error) {
	return r.c.Request("reports/view/templates/", nil)
}

// This component is optional and therefore the API will only work if it is installed
func (r Reports) TemplateDetails(template string) (map[string]interface{}, error) {
	m := map[string]string{
		"template": template,
	}
	return r.c.Request("reports/view/templateDetails/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (r Reports) Generate(title string, template string, theme string, description string, contexts []string,
	sites []string, sections []string, includedconfidences []string, includedrisks []string, reportfilename string, reportfilenamepattern string, reportdir string, display string) (map[string]interface{}, error) {
	m := map[string]string{
		"title":                 title,
		"template":              template,
		"theme":                 theme,
		"description":           description,
		"contexts":              strings.Join(contexts, ","),
		"sites":                 strings.Join(sites, ","),
		"sections":              strings.Join(sections, ","),
		"includedConfidences":   strings.Join(includedconfidences, ","),
		"includedRisks":         strings.Join(includedrisks, ","),
		"reportFileName":        reportfilename,
		"reportFileNamePattern": reportfilenamepattern,
		"reprortDir":            reportdir,
		"display":               display,
	}
	return r.c.Request("reports/action/generate/", m)
}
