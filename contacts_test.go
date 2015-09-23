package textmagic

import (
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestContacts(t *testing.T) {
	newListName := "New List Go Test"
	newContactPhone := "999000010"
	newContactFirstName := "Golang"
	newContactLastName := "Test"

	time.Sleep(interval)

	params := Params{
		"query": strings.ToLower(newListName),
	}

	listList, _ := client.SearchLists(params)

	if listList != nil && len(listList.Resources) > 0 {
		time.Sleep(interval)
		client.DeleteList(listList.Resources[0].ID)
	}

	time.Sleep(interval)

	params.Set("query", newContactPhone)
	listSearchContact, err := client.SearchContactList(params)

	if listSearchContact != nil && len(listSearchContact.Resources) > 0 {
		time.Sleep(interval)
		client.DeleteContact(listSearchContact.Resources[0].ID)
	}

	time.Sleep(interval)

	params.Del("query")
	params.Set("name", newListName)

	newList, err := client.CreateList(params)

	assert.NotEmpty(t, newList.ID)
	assert.NotEmpty(t, newList.Href)

	time.Sleep(interval)

	params.Del("name")
	params.Set("query", strings.ToLower(newListName))

	listList2, err := client.SearchLists(params)

	assert.Equal(t, 1, len(listList2.Resources))
	assert.Equal(t, newListName, listList2.Resources[0].Name)

	time.Sleep(interval)

	// Update a list name
	params.Del("query")
	params.Set("name", "updated go api test")

	updatedList, _ := client.UpdateList(listList2.Resources[0].ID, params)

	assert.NotEmpty(t, updatedList.ID)
	assert.NotEmpty(t, updatedList.Href)

	time.Sleep(interval)

	// Create a new contact
	newContactData := Params{
		"phone":     newContactPhone,
		"lists":     strconv.Itoa(newList.ID),
		"firstName": newContactFirstName,
		"lastName":  newContactLastName,
	}

	newContact, err := client.CreateContact(newContactData)

	assert.NotEmpty(t, newContact.ID)
	assert.NotEmpty(t, newContact.Href)

	time.Sleep(interval)

	// Get contact list

	listContact, err := client.GetContactList(nil)

	assert.NotEmpty(t, listContact.Page)
	assert.NotEmpty(t, listContact.Limit)
	assert.NotEmpty(t, listContact.PageCount)
	assert.NotEqual(t, len(listContact.Resources), 0)
	assert.NotEmpty(t, listContact.Resources[0].ID)

	time.Sleep(interval)

	// Get a contact by id

	contact, _ := client.GetContact(newContact.ID)
	assert.NotEmpty(t, contact.ID)
	assert.NotEmpty(t, contact.Phone)
	assert.NotEmpty(t, contact.FirstName)
	assert.NotEmpty(t, contact.LastName)

	assert.Equal(t, contact.FirstName, newContactFirstName)
	assert.Equal(t, contact.LastName, newContactLastName)
	assert.Equal(t, contact.Phone, newContactPhone)

	time.Sleep(interval)

	// Update a contact

	updatedPhone := "999000" + strconv.Itoa(random(1000, 9999))
	updatedFirstName := "Updated firstname go"
	updatedLastName := "Updated lastname go"

	updatedContactData := Params{
		"phone":     updatedPhone,
		"firstName": updatedFirstName,
		"lastName":  updatedLastName,
	}

	updatedContactData.Set("lists", newList.ID)

	updatedContactNew, err := client.UpdateContact(contact.ID, updatedContactData)

	assert.NotEmpty(t, updatedContactNew.ID)
	assert.NotEmpty(t, updatedContactNew.Href)

	time.Sleep(interval)

	// Get an updated contact

	updatedContact, err := client.GetContact(updatedContactNew.ID)

	assert.NotEmpty(t, updatedContact.ID)
	assert.NotEmpty(t, updatedContact.Phone)
	assert.NotEmpty(t, updatedContact.FirstName)
	assert.NotEmpty(t, updatedContact.LastName)

	assert.Equal(t, updatedContact.FirstName, updatedFirstName)
	assert.Equal(t, updatedContact.LastName, updatedLastName)
	assert.Equal(t, updatedContact.Phone, updatedPhone)

	time.Sleep(interval)

	// Get lists which contact belongs to

	lists, err := client.GetContactLists(newContact.ID, nil)

	assert.NotEmpty(t, lists.Page)
	assert.NotEmpty(t, lists.Limit)
	assert.NotEmpty(t, lists.PageCount)
	assert.Equal(t, len(lists.Resources), 1)

	list := lists.Resources[0]

	assert.Equal(t, list.ID, newList.ID)
	assert.Equal(t, list.MembersCount, 1)
	assert.Equal(t, list.Shared, false)

	time.Sleep(interval)

	// Create a second list
	secListName := "Sec List Go Test"
	secNewList, _ := client.CreateList(Params{"name": secListName})

	time.Sleep(interval)

	// Assign contacts to the specified list

	updatedList2, err := client.PutContactsIntoList(secNewList.ID, updatedContactNew.ID)

	assert.NotEmpty(t, updatedList2.ID)
	assert.NotEmpty(t, updatedList2.Href)

	//debug(updatedList2.Href, err, updatedList2.ID)

	time.Sleep(interval)

	// Get members count of second list

	secList, err := client.GetList(updatedList2.ID)

	assert.Equal(t, 1, secList.MembersCount)

	time.Sleep(interval)

	// Fetch user contacts by given list id

	listContact2, err := client.GetContactsInList(list.ID, nil)

	assert.NotEmpty(t, listContact2.Page)
	assert.NotEmpty(t, listContact2.Limit)
	assert.NotEmpty(t, listContact2.PageCount)
	assert.Equal(t, len(listContact2.Resources), 1)
	assert.NotEmpty(t, listContact2.Resources[0].ID)

	time.Sleep(interval)

	// Find list by id
	l, err := client.GetList(newList.ID)

	assert.Equal(t, list.Name, l.Name)
	assert.Equal(t, list.MembersCount, l.MembersCount)

	time.Sleep(interval)

	// Delete a contact

	err = client.DeleteContact(updatedContact.ID)
	assert.Nil(t, err)

	time.Sleep(interval)

	// Delete a lists

	err = client.DeleteList(newList.ID)
	assert.Nil(t, err)

	time.Sleep(interval)

	err = client.DeleteList(secList.ID)
	assert.Nil(t, err)
}
