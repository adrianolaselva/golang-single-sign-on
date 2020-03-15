
# Single Sign-On Golang

## Project structure

- `/src` Source
    - `/controllers` HTTP handler
    - `/services` some sort of business logic/use case
    - `/repositories` persistent storage/database interface
    - `/models` relational structure mapping :) 
    - `/dtos` data transfer objects
    - `/helpers` common utilities

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

## Password Grant

```sh
curl \
    --location --request POST 'http://localhost:9090/oauth2/token' \
    --header 'Content-Type: application/x-www-form-urlencoded' \
    --header 'Authorization: Basic YWRyaWFub2xhc2VsdmE6MTIzQG11ZGFy' \
    --data-urlencode 'grant_type=password' \
    --data-urlencode 'client_id=a9832dab-598c-11ea-a5a2-0242c0a8a000' \
    --data-urlencode 'client_secret=secret' \
    --data-urlencode 'scope=user:read user:write user:delete'
```

```json
{
    "token_type": "Bearer",
    "expires_in": 1584305783,
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODQzMDU3ODMsInByb2ZpbGUiOnsiaWQiOiI4ZDQyZWUzZS01NzE3LTRiNjUtYjBiNi0yMTgzNjFmOTgxYjMiLCJuYW1lIjoiQWRyaWFubyIsImxhc3RfbmFtZSI6Ik1vcmVpcmEgTGEgU2VsdmEiLCJlbWFpbCI6ImFkcmlhbm9sYXNlbHZhQGdtYWlsLmNvbSIsInVzZXJuYW1lIjoiYWRyaWFub2xhc2VsdmEiLCJhY3RpdmF0ZWQiOnRydWUsImJpcnRoZGF5IjoiMTk4Ny0wMi0xMSIsImNyZWF0ZWRfYXQiOiIyMDIwLTAzLTE0IDE2OjA1OjA0IiwidXBkYXRlZF9hdCI6IjIwMjAtMDMtMTUgMTg6NTY6MTkiLCJleHBpcmVzX2F0IjoiIiwiZGVsZXRlZF9hdCI6IiJ9LCJjbGllbnRfaWQiOiJhOTgzMmRhYi01OThjLTExZWEtYTVhMi0wMjQyYzBhOGEwMDAiLCJzY29wZSI6WyJ1c2VyOnJlYWQsdXNlcjp3cml0ZSx1c2VyOmRlbGV0ZSJdfQ.8ilhi4Ju_YRxOqYtTtM2Qbh3mYmHpNzWhd_Lf4jBCss"
}
```

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