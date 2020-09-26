# Super OS

```cgo
go get -u github.com/Alvarios/s-os
```

Advanced OS functions for server-side processes.

## CRUD

Interacting with local files has never been as easy than with sos.

### About path syntax

Path syntax in SOS can be an absolute path if it starts with a "/". Otherwise,
sos will search for the first parent directory that matches the first one of the
path and open file relatively to this one.

For example, "myProject/config/api.json" will look for the first parent directory
named "myProject" to look for the api.json file. The reason for this is that in
many Go use cases, including testing, the executable will work from its nearest
sub-directory, which won't be the root folder but a child of it.

sos takes 3 assumptions to ensure this works

- the path, if not absolute, is always relative to project root
- project folder has a unique name within the project (no child folders share their name with the root folder)
- the file is executed somewhere within the project 

### Get file

```go
package myPackage

import "github.com/Alvarios/s-os/crud"

func main() {
    var file interface{} //put any type you want here.

    err := soscrud.Get("path/to/my/file", file)
    if err != nil {
        // Handle error here.
    }
}
```

### Create file only if it doesn't exist already

```go
package myPackage

import "github.com/Alvarios/s-os/crud"

func main() {
    file, err := soscrud.Upsert("path/to/my/file")
    if err != nil {
        // Handle error here.
    }
}
```

### Delete a file

```go
package myPackage

import "github.com/Alvarios/s-os/crud"

func main() {
    err := soscrud.Delete("path/to/my/file")
    if err != nil {
        // Handle error here.
    }
}
```

## Copyright
2020 Kushuh - [MIT license](https://github.com/Alvarios/kushuh-go-utils/blob/master/LICENSE)
