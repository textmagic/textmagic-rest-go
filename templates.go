package textmagic

import "strconv"

const templateURI = "templates"

// NewTemplate represents a new template.
type NewTemplate struct {
	ID   int    `json:"id"`
	Href string `json:"href"`
}

// Template represents a template.
type Template struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Content      string `json:"content"`
	LastModified string `json:"lastModified"`
}

// TemplateList represents a template list.
type TemplateList struct {
	Page      int         `json:"page"`
	Limit     int         `json:"limit"`
	PageCount int         `json:"pageCount"`
	Resources []*Template `json:"resources"`
}

// GetTemplate returns a single message template
// for the given ID.
func (c *Client) GetTemplate(id int) (*Template, error) {
	var t *Template

	return t, c.get(templateURI+"/"+strconv.Itoa(id), nil, nil, &t)
}

// CreateTemplate creates a new template.
//
// The data payload includes:
// - name:		Template name. Required.
// - content:	Template text. May contain tags inside braces. Required.
func (c *Client) CreateTemplate(d Params) (*NewTemplate, error) {
	var t *NewTemplate

	return t, c.post(templateURI, nil, d, &t)
}

// GetTemplateList returns all user message templates.
//
// The parameter payload includes:
// - page:	Fetch specified results page.
// - limit:	How many results on page.
func (c *Client) GetTemplateList(p Params, search bool) (*TemplateList, error) {
	var l *TemplateList

	return l, c.get(templateURI, p, nil, &l)
}

// SearchTemplateList returns all user message templates
// for the given search.
//
// The parameter payload includes:
// - page:		Fetch specified results page.
// - limit:		How many results on page.
// - name: 		Find template by name.
// - content:	Find template by content.
func (c *Client) SearchTemplateList(p Params) (*TemplateList, error) {
	var l *TemplateList

	return l, c.get(templateURI+"/search", p, nil, &l)
}

// UpdateTemplate updates the template with
// the given ID.
//
// The data payload includes:
// - name:		Template name. Required.
// - content:	Template text. May contain tags inside braces. Required.
func (c *Client) UpdateTemplate(id int, d Params) (*NewTemplate, error) {
	var t *NewTemplate

	return t, c.put(templateURI+"/"+strconv.Itoa(id), nil, d, &t)
}

// DeleteTemplate deletes the template with
// the given ID.
func (c *Client) DeleteTemplate(id int) error {
	return c.delete(templateURI+"/"+strconv.Itoa(id), nil, nil, nil)
}
