# How to use it

1. Create YAML file representing a request template. e.g.

**slack_send.yml**
```
name: slack/send
description: Send slack message 
url: https://hooks.slack.com/services/XXXXXXXXX/XXXXXXXXXXX/XXXXXXXXXXXXXX
method: post
body: "{ \"text\": \"This is a test message\" }"
```

2. Install request template:

```
$ reqi install slack_send.yml
```

3. Run your request using template name:

```
$ reqi do slack/send
```

# Build from sources

1. Clone repo
2. `cd` to repo root dir
3. Execute go build
4. Done

```
$ ./reqi version
reqi v0.0.1
```

# TODO

- Improve errors messages
- Support for PUT, DELETE and PATCH methods
- Support for parameters
- Improve documentation/help
- Show, edit and delete request templates
- Support for headers