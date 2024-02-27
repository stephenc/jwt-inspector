# jwt-inspector

Decode and pretty print the environment variable `JWT_TOKEN` for use when debugging

## Example use

Using an example JWT token from https://jwt.io

```shell
$ export JWT_TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
$ docker run --rm -e JWT_TOKEN stephenc/jwt-inspector:latest
{
  "alg": "HS256",
  "typ": "JWT"
}
{
  "iat": 1516239022,
  "name": "John Doe",
  "sub": "1234567890"
}
$
```