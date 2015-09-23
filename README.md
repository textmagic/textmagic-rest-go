# ![TextMagic](https://www.textmagic.com/wp-content/themes/textmagic-genesis/images/logo/logo.svg)

# Google Go SMS API Client
This package includes functionality for accessing the the TextMagic SMS gateway in Go via the REST API.

## Installation
    go get -u github.com/textmagic/textmagic-rest-go

## Usage
Create a new client with your username and access token:

```go
c := NewClient(clientUser, clientToken)

// Get the active user
u, err := c.GetUser()

if err != nil {
    // handle error
}

fmt.Printf("My user ID is: %d\n", u.ID)
```

The `Params` type provides a friendly interface for setting request data:

```go
newListID := 123
p := NewParams("lists", newListID)
p.Set("firstName", "John")
p.Set("lastName", "Smith")

c, err := client.UpdateContact(321, p)
```


## License
The gem is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).