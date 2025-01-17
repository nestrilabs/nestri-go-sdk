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

- <code title="post /sessions">client.Sessions.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionNewParams">SessionNewParams</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionNewResponse">SessionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /sessions/{id}">client.Sessions.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionGetResponse">SessionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /sessions">client.Sessions.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionListResponse">SessionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /sessions/{id}">client.Sessions.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionDeleteResponse">SessionDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Active

Response Types:

- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionActiveListResponse">SessionActiveListResponse</a>

Methods:

- <code title="get /sessions/active">client.Sessions.Active.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionActiveService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionActiveListResponse">SessionActiveListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Public

Response Types:

- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionActivePublicListResponse">SessionActivePublicListResponse</a>

Methods:

- <code title="get /sessions/active/public">client.Sessions.Active.Public.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionActivePublicService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SessionActivePublicListResponse">SessionActivePublicListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Games

Response Types:

- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#GameNewResponse">GameNewResponse</a>
- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#GameGetResponse">GameGetResponse</a>
- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#GameUpdateResponse">GameUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#GameListResponse">GameListResponse</a>
- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#GameDeleteResponse">GameDeleteResponse</a>

Methods:

- <code title="post /games/{steamID}">client.Games.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#GameService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, steamID <a href="https://pkg.go.dev/builtin#float64">float64</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#GameNewResponse">GameNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /games/{steamID}">client.Games.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#GameService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, steamID <a href="https://pkg.go.dev/builtin#float64">float64</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#GameGetResponse">GameGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /games">client.Games.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#GameService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#GameUpdateParams">GameUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#GameUpdateResponse">GameUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /games">client.Games.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#GameService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#GameListResponse">GameListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /games/{steamID}">client.Games.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#GameService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, steamID <a href="https://pkg.go.dev/builtin#float64">float64</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#GameDeleteResponse">GameDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Sessions

Response Types:

- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#GameSessionListResponse">GameSessionListResponse</a>

Methods:

- <code title="get /games/{steamID}/sessions">client.Games.Sessions.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#GameSessionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, steamID <a href="https://pkg.go.dev/builtin#float64">float64</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#GameSessionListResponse">GameSessionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Users

Response Types:

- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#UserGetResponse">UserGetResponse</a>

Methods:

- <code title="get /users/@me">client.Users.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#UserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#UserGetResponse">UserGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Teams

Response Types:

- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#TeamNewResponse">TeamNewResponse</a>
- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#TeamGetResponse">TeamGetResponse</a>
- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#TeamListResponse">TeamListResponse</a>
- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#TeamDeleteResponse">TeamDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#TeamInviteResponse">TeamInviteResponse</a>

Methods:

- <code title="post /teams">client.Teams.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#TeamService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#TeamNewParams">TeamNewParams</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#TeamNewResponse">TeamNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /teams/{slug}">client.Teams.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#TeamService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, slug <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#TeamGetResponse">TeamGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /teams">client.Teams.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#TeamService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#TeamListResponse">TeamListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /teams/{slug}">client.Teams.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#TeamService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, slug <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#TeamDeleteResponse">TeamDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /teams/{slug}/invite/{email}">client.Teams.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#TeamService.Invite">Invite</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, slug <a href="https://pkg.go.dev/builtin#string">string</a>, email <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#TeamInviteResponse">TeamInviteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Subscriptions

Response Types:

- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SubscriptionNewResponse">SubscriptionNewResponse</a>
- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SubscriptionListResponse">SubscriptionListResponse</a>
- <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SubscriptionDeleteResponse">SubscriptionDeleteResponse</a>

Methods:

- <code title="post /subscriptions">client.Subscriptions.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SubscriptionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SubscriptionNewParams">SubscriptionNewParams</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SubscriptionNewResponse">SubscriptionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /subscriptions">client.Subscriptions.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SubscriptionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SubscriptionListResponse">SubscriptionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /subscriptions/{id}">client.Subscriptions.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SubscriptionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk">nestri</a>.<a href="https://pkg.go.dev/github.com/nestrilabs/nestri-go-sdk#SubscriptionDeleteResponse">SubscriptionDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
