# Machines

# Sessions

## Active

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
