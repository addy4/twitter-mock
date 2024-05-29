# 2witter
A networking platform _(like **twitter**)_ for posting feed, news and more...

## What does it do?
- Uses WebSockets for communication.
  - Post feed, news etc.
  - Pull notifications.
- Hydra APIs for OAuth2.0

## How it happens?
- Go to _"https://console.ory.sh/projects/b1687b4b-cf69-4ed0-a130-e23a82c4ec3c/developers"_ for creating the API Key for executing administrative Hydra APIs _(replace ORY_ID URL)_
  - **API Key:** ory_pat_EHfy924***************ert

### Creating a Client in Hydra

To create a client in Hydra using the command-line interface (CLI), you can use the following cURL command:

```bash
curl --location 'https://trusting-tereshkova-12o8uqnuqz.projects.oryapis.com/admin/clients' \
--header 'Authorization: Bearer ory_pat_92_your_api_key' \
--header 'Content-Type: application/json' \
--data '{
    "client_name": "aabhatia",
    "client_secret": "aabhatia",
    "redirect_uris": [],
    "grant_types": [
        "client_credentials"
    ],
    "response_types": [
        "id_token"
    ],
    "scope": "openid",
    "token_endpoint_auth_method": "client_secret_post",
    "access_token_strategy": "jwt"
}'
```

### List Clients in Hydra

To list client in Hydra using the command-line interface (CLI) with a modified cURL command, you can use the following:

```bash
curl --location 'https://trusting-tereshkova-12o8uqnuqz.projects.oryapis.com/admin/clients' \
--header 'Authorization: Bearer ory_hydra_api_key' \
--data '{
    "client_name": "aabhatia",
    "client_secret": "aabhatia",
    "redirect_uris": [],
    "grant_types": [
        "client_credentials"
    ],
    "response_types": [
        "id_token"
    ],
    "scope": "openid",
    "token_endpoint_auth_method": "client_secret_post",
    "access_token_strategy": "jwt"
}'
```

### Generating a Token in Hydra

To generate a token in Hydra using the command-line interface (CLI), you can use the following:

```bash
curl --location 'https://trusting-tereshkova-12o8uqnuqz.projects.oryapis.com/oauth2/token' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'client_id=b39503c6-********' \
--data-urlencode 'client_secret=aabh******' \
--data-urlencode 'grant_type=client_credentials'
```

### WebSocket Communication b/w Users
- Create multiple clients, "user1", "user2" and "aabhatia".
- Generate token for "aabhatia"
- Connect to **ws://localhost:5020/webs** and add header as **Authorization: Bearer _JWT_**
- Use following API to follow
```bash
{
    "action": "follow",
    "follow": {
        "followeeName": "user2"
    }
}
```
- Generate token for "user2"
- Connect to **ws://localhost:5020/webs** and add header as **Authorization: Bearer _JWT_**
- Use following API to post
```bash
{
    "action": "post",
    "post": {
        "content": "Hi, there!"
    }
}
```
- A notification will be receieved by other client sessions connected to WS server (e.g. aabhatia's client session)
- Connect to **ws://localhost:5020/webs** with header as **Authorization: Bearer _JWT_** where JWT is JWT of aabhatia user
- Use following API to get posts by followees
```bash
{
    "action": "posts_by_followees",
    "posts_by_followees": {
        "foo": true
    }
}
```
- _(foo is placeholder for now)_
