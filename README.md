# Getting started

1. Create YAML file representing a request template. e.g.

**slack_send.yml**
```
name: slack/send
description: Send slack message 
url: https://hooks.slack.com/services/XXXXXXXXX/XXXXXXXXXXX/XXXXXXXXXXXXXX
method: post
body: "{ \"text\": \"This is a test message\" }"
headers:
  - name: Authorization
    value: Basic aGVsbG8=
```

2. Install template:

```
$ reqi install slack_send.yml
```

3. Execute request from template:

```
$ reqi do slack/send
```

# Commands

## install {template_file}

Install a template defined in YAML file.

*Example:*

```
$ reqi install slack_send.yml
```
---
## ls

List installed templates.

*Example:*

```
$ reqi ls
+-------------+--------------------+
|    NAME     |    DESCRIPTION     |
+-------------+--------------------+
| slack/send  | Send slack message |
+-------------+--------------------+
```

---

## edit {template}

Open text editor to edit template definition.

*Example:*

```
$ reqi edit slack/send
```

---

## uninstall {template}

Uninstall template.

*Example:*

```
$ reqi uninstall slack/send
```

---

## do {template}

Execute request from template.

*Example:*

```
$ reqi do slack/send -p "message=Hello" -o result.json
```

OPTIONS:

- **-p** (optional): Resolve template parameter
- **-o** (optional): Save result to file

## Parameters

You can use template parameters in fields:
- url
- headers
- body
 
using "{{ param }}" notation, and resolve them using **-p** option (see "do" command above). e.g:

```
name: product/get
description: Get product info
url: "https://api.example.com/products/{{ id }}"
method: get
```

```
$ reqi do product/get -p "id=12"
```

# Installing from sources

1. Clone repo
2. `cd` to repo root dir
3. Execute go build
4. Done

```
$ ./reqi version
reqi v0.0.1
```

# TODO

- Support for DELETE and PATCH methods
