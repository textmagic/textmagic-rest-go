package textmagic

import (
	".."
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestContacts(t *testing.T) {
	username := "xxx"
	token := "xxx"

	interval := 500 * time.Millisecond
	client := textmagic.NewClient(username, token)

	newListName := "New List Go Test"
	newContactPhone := "999000010"
	newContactFirstName := "Golang"
	newContactLastName := "Test"

	time.Sleep(interval)

	// Find a list by name and delete it if exists
	listSearchData := map[string]string{
		"query": strings.ToLower(newListName),
	}
	listList, _ := client.GetListList(listSearchData, true)
	if len(listList.Lists) > 0 {
		time.Sleep(interval)
		client.DeleteList(listList.Lists[0].Id)
	}

	time.Sleep(interval)

	// Find a contact by phone and delete it if exists
	contactSearchData := map[string]string{
		"query": newContactPhone,
	}
	listSearchContact, _ := client.GetContactList(contactSearchData, true)
	if len(listSearchContact.Contacts) > 0 {
		time.Sleep(interval)
		client.DeleteContact(listSearchContact.Contacts[0].Id)
	}

	time.Sleep(interval)

	// Create a new List to assign contact
	newListData := map[string]string{
		"name": newListName,
	}

	newList, _ := client.CreateList(newListData)

	assert.NotEmpty(t, newList.Id)
	assert.NotEmpty(t, newList.Href)

	time.Sleep(interval)

	// Find a list by name
	listSearchData2 := map[string]string{
		"query": strings.ToLower(newListName),
	}
	listList2, _ := client.GetListList(listSearchData2, true)

	assert.Equal(t, 1, len(listList2.Lists))
	assert.Equal(t, newListName, listList2.Lists[0].Name)

	time.Sleep(interval)

	// Update a list name
	updatedListName := "updated go api test"
	updatedList, _ := client.UpdateList(
		listList2.Lists[0].Id,
		map[string]string{
			"name": updatedListName,
		},
	)

	assert.NotEmpty(t, updatedList.Id)
	assert.NotEmpty(t, updatedList.Href)

	time.Sleep(interval)

	// Create a new contact
	newContactData := map[string]string{
		"phone":     newContactPhone,
		"lists":     strconv.Itoa(int(newList.Id)),
		"firstName": newContactFirstName,
		"lastName":  newContactLastName,
	}

	newContact, _ := client.CreateContact(newContactData)

	assert.NotEmpty(t, newContact.Id)
	assert.NotEmpty(t, newContact.Href)

	time.Sleep(interval)

	// Get contact list

	listContact, _ := client.GetContactList(map[string]string{}, false)

	assert.NotEmpty(t, listContact.Page)
	assert.NotEmpty(t, listContact.Limit)
	assert.NotEmpty(t, listContact.PageCount)
	assert.NotEqual(t, len(listContact.Contacts), 0)
	assert.NotEmpty(t, listContact.Contacts[0].Id)

	time.Sleep(interval)

	// Get a contact by id

	contact, _ := client.GetContact(newContact.Id)

	assert.NotEmpty(t, contact.Id)
	assert.NotEmpty(t, contact.Phone)
	assert.NotEmpty(t, contact.FirstName)
	assert.NotEmpty(t, contact.LastName)

	assert.Equal(t, contact.FirstName, newContactFirstName)
	assert.Equal(t, contact.LastName, newContactLastName)
	assert.Equal(t, contact.Phone, newContactPhone)

	time.Sleep(interval)

	// Update a contact

	updatedPhone := "9990000321"
	updatedFirstName := "Updated firstname go"
	updatedLastName := "Updated lastname go"

	updatedContactData := map[string]string{
		"phone":     updatedPhone,
		"firstName": updatedFirstName,
		"lastName":  updatedLastName,
		"lists":     strconv.Itoa(int(newList.Id)),
	}

	updatedContactNew, _ := client.UpdateContact(contact.Id, updatedContactData)

	assert.NotEmpty(t, updatedContactNew.Id)
	assert.NotEmpty(t, updatedContactNew.Href)

	time.Sleep(interval)

	// Get an updated contact

	updatedContact, _ := client.GetContact(updatedContactNew.Id)

	assert.NotEmpty(t, updatedContact.Id)
	assert.NotEmpty(t, updatedContact.Phone)
	assert.NotEmpty(t, updatedContact.FirstName)
	assert.NotEmpty(t, updatedContact.LastName)

	assert.Equal(t, updatedContact.FirstName, updatedFirstName)
	assert.Equal(t, updatedContact.LastName, updatedLastName)
	assert.Equal(t, updatedContact.Phone, updatedPhone)

	time.Sleep(interval)

	// Get lists which contact belongs to

	lists, _ := client.GetContactLists(newContact.Id)

	assert.NotEmpty(t, lists.Page)
	assert.NotEmpty(t, lists.Limit)
	assert.NotEmpty(t, lists.PageCount)
	assert.Equal(t, len(lists.Lists), 1)

	list := lists.Lists[0]

	assert.Equal(t, list.Id, newList.Id)
	assert.Equal(t, list.MembersCount, uint32(1))
	assert.Equal(t, list.Shared, false)

	time.Sleep(interval)

	// Create a second list
	secListName := "Sec List Go Test"
	secNewList, _ := client.CreateList(
		map[string]string{
			"name": secListName,
		},
	)

	time.Sleep(interval)

	// Assign contacts to the specified list

	updatedList2, _ := client.PutContactsIntoList(secNewList.Id, strconv.Itoa(int(updatedContactNew.Id)))

	assert.NotEmpty(t, updatedList2.Id)
	assert.NotEmpty(t, updatedList2.Href)

	time.Sleep(interval)

	// Get members count of second list

	secList, _ := client.GetList(secNewList.Id)

	assert.Equal(t, uint32(1), secList.MembersCount)

	time.Sleep(interval)

	// Fetch user contacts by given list id

	listContact2, _ := client.GetContactsInList(list.Id, map[string]string{})

	assert.NotEmpty(t, listContact2.Page)
	assert.NotEmpty(t, listContact2.Limit)
	assert.NotEmpty(t, listContact2.PageCount)
	assert.Equal(t, len(listContact2.Contacts), 1)
	assert.NotEmpty(t, listContact2.Contacts[0].Id)

	time.Sleep(interval)

	// Find list by id
	l, _ := client.GetList(newList.Id)

	assert.Equal(t, list.Name, l.Name)
	assert.Equal(t, list.MembersCount, l.MembersCount)

	time.Sleep(interval)

	// Delete a contact

	r, _ := client.DeleteContact(updatedContact.Id)
	assert.True(t, r)

	time.Sleep(interval)

	// Delete a lists

	r, _ = client.DeleteList(newList.Id)
	assert.True(t, r)

	time.Sleep(interval)

	r, _ = client.DeleteList(secList.Id)
	assert.True(t, r)
}
