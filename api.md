# Machines

Response Types:

- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#MachineNewResponse">MachineNewResponse</a>
- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#MachineGetResponse">MachineGetResponse</a>
- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#MachineListResponse">MachineListResponse</a>
- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#MachineDeleteResponse">MachineDeleteResponse</a>

Methods:

- <code title="post /machines/{fingerprint}">client.Machines.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#MachineService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fingerprint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#MachineNewResponse">MachineNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /machines/{fingerprint}">client.Machines.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#MachineService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fingerprint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#MachineGetResponse">MachineGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /machines">client.Machines.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#MachineService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#MachineListResponse">MachineListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /machines/{fingerprint}">client.Machines.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#MachineService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fingerprint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#MachineDeleteResponse">MachineDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Sessions

Response Types:

- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionNewResponse">SessionNewResponse</a>
- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionGetResponse">SessionGetResponse</a>
- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionListResponse">SessionListResponse</a>
- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionDeleteResponse">SessionDeleteResponse</a>

Methods:

- <code title="post /sessions/{id}">client.Sessions.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionNewParams">SessionNewParams</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionNewResponse">SessionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /sessions/{id}">client.Sessions.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionGetResponse">SessionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /sessions">client.Sessions.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionListResponse">SessionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /sessions/{id}">client.Sessions.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionDeleteResponse">SessionDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
