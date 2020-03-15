
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

#### Sample request

```sh
curl \
    --location --request POST 'http://localhost:9090/oauth2/token' \
    --header 'Content-Type: application/x-www-form-urlencoded' \
    --header 'Authorization: Basic YWRyaWFub2xhc2VsdmE6MTIzQG11ZGFy' \
    --data-urlencode 'grant_type=password' \
    --data-urlencode 'client_id=a9832dab-598c-11ea-a5a2-0242c0a8a000' \
    --data-urlencode 'client_secret=8zd9ULma6xNN1wbR7h8er7z7qbERULsjCqD2pzT5' \
    --data-urlencode 'scope=user:read user:write user:delete'
```

#### Sample response

```json
{
    "token_type": "Bearer",
    "expires_in": 1584305783,
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODQzMDU3ODMsInByb2ZpbGUiOnsiaWQiOiI4ZDQyZWUzZS01NzE3LTRiNjUtYjBiNi0yMTgzNjFmOTgxYjMiLCJuYW1lIjoiQWRyaWFubyIsImxhc3RfbmFtZSI6Ik1vcmVpcmEgTGEgU2VsdmEiLCJlbWFpbCI6ImFkcmlhbm9sYXNlbHZhQGdtYWlsLmNvbSIsInVzZXJuYW1lIjoiYWRyaWFub2xhc2VsdmEiLCJhY3RpdmF0ZWQiOnRydWUsImJpcnRoZGF5IjoiMTk4Ny0wMi0xMSIsImNyZWF0ZWRfYXQiOiIyMDIwLTAzLTE0IDE2OjA1OjA0IiwidXBkYXRlZF9hdCI6IjIwMjAtMDMtMTUgMTg6NTY6MTkiLCJleHBpcmVzX2F0IjoiIiwiZGVsZXRlZF9hdCI6IiJ9LCJjbGllbnRfaWQiOiJhOTgzMmRhYi01OThjLTExZWEtYTVhMi0wMjQyYzBhOGEwMDAiLCJzY29wZSI6WyJ1c2VyOnJlYWQsdXNlcjp3cml0ZSx1c2VyOmRlbGV0ZSJdfQ.8ilhi4Ju_YRxOqYtTtM2Qbh3mYmHpNzWhd_Lf4jBCss"
}
```

## Client credentials Grant

#### Sample request

>passing client_id and client_secret in the header
```sh
curl \
    --location --request POST 'http://localhost:9090/oauth2/token' \
    --header 'Content-Type: application/x-www-form-urlencoded' \
    --header 'Authorization: Basic YTk4MzJkYWItNTk4Yy0xMWVhLWE1YTItMDI0MmMwYThhMDAwOjh6ZDlVTG1hNnhOTjF3YlI3aDhlcjd6N3FiRVJVTHNqQ3FEMnB6VDU=' \
    --data-urlencode 'grant_type=client_credentials' \
    --data-urlencode 'scope=user:read user:write user:delete'
```

>passing client_id and client_secret in the body
```sh
curl \
    --location --request POST 'http://localhost:9090/oauth2/token' \
    --header 'Content-Type: application/x-www-form-urlencoded' \
    --data-urlencode 'grant_type=client_credentials' \
    --data-urlencode 'client_id=a9832dab-598c-11ea-a5a2-0242c0a8a000' \
    --data-urlencode 'client_secret=8zd9ULma6xNN1wbR7h8er7z7qbERULsjCqD2pzT5' \
    --data-urlencode 'scope=user:read user:write user:delete'
```

#### Sample response

```json
{
    "token_type": "Bearer",
    "expires_in": 1584308456,
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODQzMDg0NTYsInByb2ZpbGUiOnsiaWQiOiI4ZDQyZWUzZS01NzE3LTRiNjUtYjBiNi0yMTgzNjFmOTgxYjMiLCJuYW1lIjoiQWRyaWFubyIsImxhc3RfbmFtZSI6Ik1vcmVpcmEgTGEgU2VsdmEiLCJlbWFpbCI6ImFkcmlhbm9sYXNlbHZhQGdtYWlsLmNvbSIsInVzZXJuYW1lIjoiYWRyaWFub2xhc2VsdmEiLCJhY3RpdmF0ZWQiOnRydWUsImJpcnRoZGF5IjoiMTk4Ny0wMi0xMSIsImNyZWF0ZWRfYXQiOiIyMDIwLTAzLTE1IDE5OjI5OjAxIiwidXBkYXRlZF9hdCI6IjIwMjAtMDMtMTUgMTk6MzU6NTAiLCJleHBpcmVzX2F0IjoiIiwiZGVsZXRlZF9hdCI6IiJ9LCJjbGllbnRfaWQiOiJhOTgzMmRhYi01OThjLTExZWEtYTVhMi0wMjQyYzBhOGEwMDAiLCJzY29wZSI6WyJ1c2VyOnJlYWQsdXNlcjp3cml0ZSx1c2VyOmRlbGV0ZSJdfQ.6rKFJbmM-Uhgnju5SXDjmtnRpiUczhD3rnC0X-gTu-M"
}
```

## Refresh token Grant

#### Sample request

```sh
curl \
    --location --request POST 'http://localhost:9090/oauth2/token' \
    --header 'Content-Type: application/x-www-form-urlencoded' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODQzMTM4NDMsInByb2ZpbGUiOnsiaWQiOiI4ZDQyZWUzZS01NzE3LTRiNjUtYjBiNi0yMTgzNjFmOTgxYjMiLCJuYW1lIjoiQWRyaWFubyIsImxhc3RfbmFtZSI6Ik1vcmVpcmEgTGEgU2VsdmEiLCJlbWFpbCI6ImFkcmlhbm9sYXNlbHZhQGdtYWlsLmNvbSIsInVzZXJuYW1lIjoiYWRyaWFub2xhc2VsdmEiLCJhY3RpdmF0ZWQiOnRydWUsImJpcnRoZGF5IjoiMTk4Ny0wMi0xMSIsImNyZWF0ZWRfYXQiOiIyMDIwLTAzLTE1IDE5OjI5OjAxIiwidXBkYXRlZF9hdCI6IjIwMjAtMDMtMTUgMjE6MTA6MDgiLCJleHBpcmVzX2F0IjoiIiwiZGVsZXRlZF9hdCI6IiJ9LCJjbGllbnRfaWQiOiJhOTgzMmRhYi01OThjLTExZWEtYTVhMi0wMjQyYzBhOGEwMDAiLCJzY29wZSI6WyJ1c2VyOnJlYWQsdXNlcjp3cml0ZSx1c2VyOmRlbGV0ZSJdfQ.Opm1C8SY8k-4w2oZbOa-3Le8ChAKwHadWBd-W37bheA' \
    --data-urlencode 'grant_type=refresh_token' \
    --data-urlencode 'refresh_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODUxNzA2NDMsInByb2ZpbGUiOnsiaWQiOiIiLCJuYW1lIjoiIiwibGFzdF9uYW1lIjoiIiwiZW1haWwiOiIiLCJ1c2VybmFtZSI6IiIsImFjdGl2YXRlZCI6ZmFsc2UsImJpcnRoZGF5IjoiIiwiY3JlYXRlZF9hdCI6IiIsInVwZGF0ZWRfYXQiOiIiLCJleHBpcmVzX2F0IjoiIiwiZGVsZXRlZF9hdCI6IiJ9fQ.qKFvo1sh0jiVYfghkGY1PEGODCXz4WHrx5dSFN57L_M'
```

#### Sample response

```json
{
    "token_type": "Bearer",
    "expires_in": 1584314330,
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODQzMTQzMzAsInByb2ZpbGUiOnsiaWQiOiI4ZDQyZWUzZS01NzE3LTRiNjUtYjBiNi0yMTgzNjFmOTgxYjMiLCJuYW1lIjoiQWRyaWFubyIsImxhc3RfbmFtZSI6Ik1vcmVpcmEgTGEgU2VsdmEiLCJlbWFpbCI6ImFkcmlhbm9sYXNlbHZhQGdtYWlsLmNvbSIsInVzZXJuYW1lIjoiYWRyaWFub2xhc2VsdmEiLCJhY3RpdmF0ZWQiOnRydWUsImJpcnRoZGF5IjoiMTk4Ny0wMi0xMSIsImNyZWF0ZWRfYXQiOiIyMDIwLTAzLTE1IDE5OjI5OjAxIiwidXBkYXRlZF9hdCI6IjIwMjAtMDMtMTUgMTg6MTg6NTAiLCJleHBpcmVzX2F0IjoiIiwiZGVsZXRlZF9hdCI6IiJ9LCJjbGllbnRfaWQiOiJhOTgzMmRhYi01OThjLTExZWEtYTVhMi0wMjQyYzBhOGEwMDAiLCJzY29wZSI6WyJ1c2VyOnJlYWQiLCJ1c2VyOndyaXRlIiwidXNlcjpkZWxldGUiXX0.weGED6dGrkj6ZzN1Mtc1X0i8uC15XmR97RcKW7J2apg",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODUxNzExMzAsInByb2ZpbGUiOnsiaWQiOiIiLCJuYW1lIjoiIiwibGFzdF9uYW1lIjoiIiwiZW1haWwiOiIiLCJ1c2VybmFtZSI6IiIsImFjdGl2YXRlZCI6ZmFsc2UsImJpcnRoZGF5IjoiIiwiY3JlYXRlZF9hdCI6IiIsInVwZGF0ZWRfYXQiOiIiLCJleHBpcmVzX2F0IjoiIiwiZGVsZXRlZF9hdCI6IiJ9fQ.pjBs53aa49Uxqx-tLQIJbIEmepVuNVV4_vQIsktD6aE"
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