package textmagic

import (
	".."
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestCustomFields(t *testing.T) {
	username := "xxx"
	token := "xxx"

	client := textmagic.NewClient(username, token)

	newCustomFieldName := "Test Go Custom Field"

	time.Sleep(500 * time.Millisecond)
	// Create a new custom field

	newCustomFieldData := map[string]string{
		"name": newCustomFieldName,
	}

	newCustomField, _ := client.CreateCustomField(newCustomFieldData)

	assert.NotEmpty(t, newCustomField.Id)
	assert.NotEmpty(t, newCustomField.Href)

	time.Sleep(500 * time.Millisecond)
	// Get custom field

	customField, _ := client.GetCustomField(newCustomField.Id)

	assert.NotEmpty(t, customField.Id)
	assert.NotEmpty(t, customField.Name)
	assert.NotEmpty(t, customField.CreatedAt)
	assert.Equal(t, customField.Name, newCustomFieldName)

	time.Sleep(500 * time.Millisecond)
	// Get custom field list

	listCustomFields, _ := client.GetCustomFieldList(map[string]string{})

	assert.NotEmpty(t, listCustomFields.Page)
	assert.NotEmpty(t, listCustomFields.Limit)
	assert.NotEmpty(t, listCustomFields.PageCount)
	assert.NotEqual(t, len(listCustomFields.CustomFields), 0)
	assert.NotEmpty(t, listCustomFields.CustomFields[0].Id)

	time.Sleep(500 * time.Millisecond)
	// Update custom field

	updatedName := "updated go customfield"

	updatedCustomFieldNew, _ := client.UpdateCustomField(customField.Id, map[string]string{
		"name": updatedName,
	})

	assert.NotEmpty(t, updatedCustomFieldNew.Id)
	assert.NotEmpty(t, updatedCustomFieldNew.Href)
	assert.Equal(t, updatedCustomFieldNew.Id, customField.Id)

	time.Sleep(500 * time.Millisecond)
	// Get updated custom field

	updatedCustomField, _ := client.GetCustomField(newCustomField.Id)

	assert.NotEmpty(t, updatedCustomField.Id)
	assert.NotEmpty(t, updatedCustomField.Name)
	assert.NotEmpty(t, updatedCustomField.CreatedAt)
	assert.Equal(t, updatedCustomField.Name, updatedName)

	newContactPhone := "999000010"
	// Find a contact by phone and delete it if exists
	contactSearchData := map[string]string{
		"query": newContactPhone,
	}
	listSearchContact, _ := client.GetContactList(contactSearchData, true)

	var cid uint32
	newListName := "new_list_custom_fields"

	if len(listSearchContact.Contacts) > 0 {
		time.Sleep(500 * time.Millisecond)
		cid = listSearchContact.Contacts[0].Id
		client.DeleteContact(cid)
	}

	time.Sleep(500 * time.Millisecond)

	newList, _ := client.CreateList(map[string]string{
		"name": newListName,
	})

	time.Sleep(500 * time.Millisecond)

	newContact, _ := client.CreateContact(map[string]string{
		"lists": strconv.Itoa(int(newList.Id)),
		"phone": newContactPhone,
	})
	cid = newContact.Id

	time.Sleep(500 * time.Millisecond)
	// Update contact's custom field value.

	updatedValue := "gopher"
	updateData := map[string]string{
		"contactId": strconv.Itoa(int(cid)),
		"value":     updatedValue,
	}

	updatedContact, _ := client.UpdateCustomFieldValue(customField.Id, updateData)

	assert.NotEmpty(t, updatedContact.Id)
	assert.NotEmpty(t, updatedContact.Href)

	time.Sleep(500 * time.Millisecond)
	// Get update custom field value

	contact, _ := client.GetContact(updatedContact.Id)

	contactCustomFields := contact.CustomFields

	for _, field := range contactCustomFields {
		if field.Id == customField.Id {
			assert.Equal(t, updatedValue, field.Value)
		} else {
			t.Fail()
		}
	}

	time.Sleep(500 * time.Millisecond)
	// Delete a list

	client.DeleteList(newList.Id)

	time.Sleep(500 * time.Millisecond)
	// Delete a contact

	client.DeleteContact(cid)

	time.Sleep(500 * time.Millisecond)
	// Delete a custom field

	r, _ := client.DeleteCustomField(newCustomField.Id)

	assert.True(t, r)
}
