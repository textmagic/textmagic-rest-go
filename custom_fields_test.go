package textmagic

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCustomFields(t *testing.T) {
	time.Sleep(interval)

	l, err := client.GetCustomFieldList(nil)

	// Purge existing fields
	if err == nil && l != nil && len(l.Resources) > 0 {
		for _, r := range l.Resources {
			time.Sleep(interval)

			client.DeleteCustomField(r.ID)
		}
	}

	newCustomFieldName := "Custom Field"

	time.Sleep(interval)

	// Create a new custom field
	newCustomField, err := client.CreateCustomField(newCustomFieldName)

	assert.NotEmpty(t, newCustomField.ID)
	assert.NotEmpty(t, newCustomField.Href)

	time.Sleep(interval)
	// Get custom field

	customField, _ := client.GetCustomField(newCustomField.ID)

	assert.NotEmpty(t, customField.ID)
	assert.NotEmpty(t, customField.Name)
	assert.NotEmpty(t, customField.CreatedAt)
	assert.Equal(t, customField.Name, newCustomFieldName)

	time.Sleep(interval)
	// Get custom field list

	listCustomFields, _ := client.GetCustomFieldList(nil)

	assert.NotEmpty(t, listCustomFields.Page)
	assert.NotEmpty(t, listCustomFields.Limit)
	assert.NotEmpty(t, listCustomFields.PageCount)
	assert.NotEqual(t, len(listCustomFields.Resources), 0)
	assert.NotEmpty(t, listCustomFields.Resources[0].ID)

	time.Sleep(interval)
	// Update custom field

	updatedName := "updated go customfield"

	updatedCustomFieldNew, _ := client.UpdateCustomField(customField.ID, updatedName)

	if updatedCustomFieldNew == nil {
		t.Fatal("nil UpdateCustomField")
	} else {
		assert.NotEmpty(t, updatedCustomFieldNew.ID)
		assert.NotEmpty(t, updatedCustomFieldNew.Href)
		assert.Equal(t, updatedCustomFieldNew.ID, customField.ID)
	}

	time.Sleep(interval)
	// Get updated custom field

	updatedCustomField, _ := client.GetCustomField(newCustomField.ID)

	assert.NotEmpty(t, updatedCustomField.ID)
	assert.NotEmpty(t, updatedCustomField.Name)
	assert.NotEmpty(t, updatedCustomField.CreatedAt)
	assert.Equal(t, updatedCustomField.Name, updatedName)

	newContactPhone := "999000010"
	// Find a contact by phone and delete it if exists
	contactSearchData := Params{"query": newContactPhone}
	listSearchContact, _ := client.SearchContactList(contactSearchData)

	var cid int
	newListName := "new_list_custom_fields"

	if listSearchContact != nil && len(listSearchContact.Resources) > 0 {
		time.Sleep(interval)
		cid = listSearchContact.Resources[0].ID
		client.DeleteContact(cid)
	}

	time.Sleep(interval)

	newList, _ := client.CreateList(Params{"name": newListName})

	time.Sleep(interval)

	newContact, _ := client.CreateContact(Params{
		"lists": strconv.Itoa(newList.ID),
		"phone": newContactPhone,
	})
	cid = newContact.ID

	time.Sleep(interval)
	// Update contact's custom field value.

	updatedValue := "gopher"
	updateData := Params{
		"contactId": strconv.Itoa(cid),
		"value":     updatedValue,
	}

	updatedContact, _ := client.UpdateCustomFieldValue(customField.ID, updateData)

	assert.NotEmpty(t, updatedContact.ID)
	assert.NotEmpty(t, updatedContact.Href)

	time.Sleep(interval)
	// Get update custom field value

	contact, _ := client.GetContact(updatedContact.ID)

	contactCustomFields := contact.CustomFields

	for _, field := range contactCustomFields {
		if field.ID == customField.ID {
			assert.Equal(t, updatedValue, field.Value)
		} else {
			t.Fail()
		}
	}

	time.Sleep(interval)
	// Delete a list

	client.DeleteList(newList.ID)

	time.Sleep(interval)
	// Delete a contact

	client.DeleteContact(cid)

	time.Sleep(interval)
	// Delete a custom field

	err = client.DeleteCustomField(newCustomField.ID)

	assert.Nil(t, err)
}
