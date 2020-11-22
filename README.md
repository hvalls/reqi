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

# Commands

## install {template_file}

Installs a template defined in {template_file} YAML file.

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

Open text editor to edit {template} template definition.

*Example:*

```
$ reqi edit slack/send
```

---

## uninstall {template}

Uninstalls template.

*Example:*

```
$ reqi uninstall slack/send
```

---

## do {template}

Executes request using template {template}.

*Example:*

```
$ reqi do slack/send
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

- Support for PUT, DELETE and PATCH methods
- Support for parameters
- Improve documentation/help
- Support for headers
- Support for save output (-o option)