
# Single Sign-On Golang


## Migrations 

### Create migration

```sh
sql-migrate new create-users-table
```

### Run migrations

```sh
sql-migrate up
```

### Rollback migrations

```sh
sql-migrate down
```

### View status migrations

```sh
sql-migrate status
```



## Example authorization flows

## Authorization Code Grant

#### Sample request

```sh
curl -XGET http://localhost:9099/oauth2/authorize? \
    response_type=code& \
    client_id=ad398u21ijw3s9w3939& \
    redirect_uri=https://YOUR_APP/redirect_uri& \
    state=STATE& \
    scope=customer:write+customer:read
```

#### Sample response

```http
HTTP/1.1 302 Found
Location: https://YOUR_APP/redirect_uri?code=AUTHORIZATION_CODE&state=STATE
```


## Token grant without `openid` scope

#### Sample Request

```sh
curl -XGET http://localhost:9099/oauth2/authorize? \
    response_type=token& \
    client_id=ad398u21ijw3s9w3939& \
    redirect_uri=https://YOUR_APP/redirect_uri& \
    state=STATE& \
    scope=customer:write+customer:read
```

#### Sample response

```http
HTTP/1.1 302 Found
Location: https://YOUR_APP/redirect_uri?access_token=ACCESS_TOKEN&token_type=bearer&expires_in=3600&state=STATE
```

#### Examples of Negative Responses

#### Invalid request

```http
HTTP/1.1 302 Found
Location: https://YOUR_APP/redirect_uri?error=invalid_request
```

#### Unauthorized client

```http
HTTP/1.1 302 Found
Location: https://YOUR_APP/redirect_uri?error=unauthorized_client
```

#### Invalid scope

```http
HTTP/1.1 302 Found
Location: https://YOUR_APP/redirect_uri?error=invalid_scope
```

#### Internal server error

```http
HTTP/1.1 302 Found
Location: https://YOUR_APP/redirect_uri?error=server_error
```

#### Invalid request with error description

```http
HTTP/1.1 302 Found
Location: https://YOUR_APP/redirect_uri?error=invalid_request&error_description=Timeout
```