# Machines

# Sessions

Response Types:

- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionNewResponse">SessionNewResponse</a>
- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionGetResponse">SessionGetResponse</a>
- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionDeleteResponse">SessionDeleteResponse</a>

Methods:

- <code title="post /sessions">client.Sessions.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionNewParams">SessionNewParams</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionNewResponse">SessionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /sessions/{id}">client.Sessions.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionGetResponse">SessionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /sessions/{id}">client.Sessions.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionDeleteResponse">SessionDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Active

Response Types:

- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionActiveListResponse">SessionActiveListResponse</a>

Methods:

- <code title="get /sessions/active">client.Sessions.Active.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionActiveService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionActiveListResponse">SessionActiveListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Public

# Games

## Sessions

# Users

Response Types:

- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#UserGetResponse">UserGetResponse</a>

Methods:

- <code title="get /users/@me">client.Users.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#UserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#UserGetResponse">UserGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Teams

# Subscriptions

Response Types:

- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SubscriptionNewResponse">SubscriptionNewResponse</a>
- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SubscriptionListResponse">SubscriptionListResponse</a>
- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SubscriptionDeleteResponse">SubscriptionDeleteResponse</a>

Methods:

- <code title="post /subscriptions">client.Subscriptions.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SubscriptionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SubscriptionNewParams">SubscriptionNewParams</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SubscriptionNewResponse">SubscriptionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /subscriptions">client.Subscriptions.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SubscriptionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SubscriptionListResponse">SubscriptionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /subscriptions/{id}">client.Subscriptions.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SubscriptionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SubscriptionDeleteResponse">SubscriptionDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
