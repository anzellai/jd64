# jd64


## Summary

Simple command line tool to decode base64'd payload either in normal string or JSON or one of those field(s) in JSON payload.
Default to return original input if error is found.

This is a quick tool to help unpack base64'd message in a variety of formats.


## Usage 

Decode base64 JSON string

```
✦ ❯ echo -n '{"message":"hello"}' |base64
eyJtZXNzYWdlIjoiaGVsbG8ifQ==
```

```
✦ ❯ jd64 -in "eyJtZXNzYWdlIjoiaGVsbG8ifQ=="
{
  "message": "hello"
}
```

Same, but using pipe or from stdin

```
✦ ❯ echo -n '{"message":"hello"}' |base64 |jd64
{
  "message": "hello"
}
```


For a field in base64'd JSON
```
✦ ❯ echo -n '{"msg":"world","payload":"eyJtZXNzYWdlIjoiaGVsbG8ifQ=="}' |jd64
{
  "msg": "world",
  "payload": {
    "message": "hello"
  }
}
```

It is a quick tool, and it does not take care nested JSON, I also expect many bugs for different use cases.
PRs welcomed.
