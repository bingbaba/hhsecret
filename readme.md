```sh
# login
curl -XPOST "http://localhost:8003/api/user/{username}/login" -d "{\"password\":\"password\"}"

# sign
curl -XPOST "http://localhost:8003/api/user/{username}/sign"

# sign list
curl "http://localhost:8003/api/user/{username}/sign"
```
