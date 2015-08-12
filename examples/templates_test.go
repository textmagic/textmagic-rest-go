package textmagic

import (
	".."
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTemplates(t *testing.T) {
	username := "xxx"
	token := "xxx"

	interval := 500 * time.Millisecond
	client := textmagic.NewClient(username, token)

	time.Sleep(interval)
	// Create template

	newTemplateData := map[string]string{
		"name":    "GO TEMPLATE TEST",
		"content": "GO TEMPLATE CONTENT",
	}
	templateNew, _ := client.CreateTemplate(newTemplateData)

	assert.NotEmpty(t, templateNew.Id)
	assert.NotEmpty(t, templateNew.Href)

	time.Sleep(interval)
	// Get template by id

	template, _ := client.GetTemplate(templateNew.Id)

	assert.Equal(t, templateNew.Id, template.Id)
	assert.Equal(t, "GO TEMPLATE TEST", template.Name)
	assert.Equal(t, "GO TEMPLATE CONTENT", template.Content)

	time.Sleep(interval)
	// Get Templates List

	templates, _ := client.GetTemplateList(map[string]string{}, false)

	assert.NotEmpty(t, templates.Page)
	assert.NotEmpty(t, templates.Limit)
	assert.NotEmpty(t, templates.PageCount)
	assert.NotEmpty(t, len(templates.Templates))

	time.Sleep(interval)
	// Update template

	updatedTemplateData := map[string]string{
		"name":    "GO TEMPLATE UPD",
		"content": "GO TEMPLATE CONTENT UPD",
	}
	updatedTemplateNew, _ := client.UpdateTemplate(templateNew.Id, updatedTemplateData)

	assert.NotEmpty(t, updatedTemplateNew.Id)
	assert.NotEmpty(t, updatedTemplateNew.Href)

	time.Sleep(interval)
	// Get updated template by id

	updatedTemplate, _ := client.GetTemplate(templateNew.Id)

	assert.Equal(t, templateNew.Id, updatedTemplate.Id)
	assert.Equal(t, "GO TEMPLATE UPD", updatedTemplate.Name)
	assert.Equal(t, "GO TEMPLATE CONTENT UPD", updatedTemplate.Content)

	time.Sleep(interval)
	// Delete template

	r, _ := client.DeleteTemplate(template.Id)

	assert.True(t, r)
}
