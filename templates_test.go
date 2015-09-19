package textmagic

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTemplates(t *testing.T) {
	time.Sleep(interval)
	// Create template

	newTemplateData := Params{
		"name":    "GO TEMPLATE TEST",
		"content": "GO TEMPLATE CONTENT",
	}
	templateNew, _ := client.CreateTemplate(newTemplateData)

	assert.NotEmpty(t, templateNew.ID)
	assert.NotEmpty(t, templateNew.Href)

	time.Sleep(interval)
	// Get template by id

	template, _ := client.GetTemplate(templateNew.ID)

	assert.Equal(t, templateNew.ID, template.ID)
	assert.Equal(t, "GO TEMPLATE TEST", template.Name)
	assert.Equal(t, "GO TEMPLATE CONTENT", template.Content)

	time.Sleep(interval)
	// Get Templates List

	templates, _ := client.GetTemplateList(nil, false)

	assert.NotEmpty(t, templates.Page)
	assert.NotEmpty(t, templates.Limit)
	assert.NotEmpty(t, templates.PageCount)
	assert.NotEmpty(t, len(templates.Resources))

	time.Sleep(interval)
	// Update template

	updatedTemplateData := Params{
		"name":    "GO TEMPLATE UPD",
		"content": "GO TEMPLATE CONTENT UPD",
	}
	updatedTemplateNew, _ := client.UpdateTemplate(templateNew.ID, updatedTemplateData)

	assert.NotEmpty(t, updatedTemplateNew.ID)
	assert.NotEmpty(t, updatedTemplateNew.Href)

	time.Sleep(interval)
	// Get updated template by id

	updatedTemplate, _ := client.GetTemplate(templateNew.ID)

	assert.Equal(t, templateNew.ID, updatedTemplate.ID)
	assert.Equal(t, "GO TEMPLATE UPD", updatedTemplate.Name)
	assert.Equal(t, "GO TEMPLATE CONTENT UPD", updatedTemplate.Content)

	time.Sleep(interval)
	// Delete template

	err := client.DeleteTemplate(template.ID)

	assert.Nil(t, err)
}
