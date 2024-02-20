# regexlint

regexlint - simple application for regex validation, e.g. in CI pipelines for checking trust domains

## Use cases

### 🐘 Check PCRE (PERL, PHP) regex

```sh
regexlint pcre "invalid-regex())"
```

#### Check regex from yaml file with name `conf.yaml`:
```yaml
headers:
  cors: https:\/\/.*test-app.com
```

```sh
regexlint pcre conf.yaml headers.cors
```

#### Check regex from json file with name `conf.json`:
```json
{
  "headers": {
    "cors": "https:\/\/.*test-app.com"
  }
}
```

```sh
regexlint pcre conf.json headers.cors
```

#### Check regex from .txt file by row number
```text
string
https:\/\/.*test-app.com
```

```sh
regexlint pcre file.txt row-1
```

### 🦫 Check Go regex

```sh
regexlint go "invalid-regex())"
```

#### Check regex from yaml file with name `conf.yaml`:
```yaml
headers:
  cors: https:\/\/.*test-app.com
```

```sh
regexlint go conf.yaml headers.cors
```

#### Check regex from json file with name `conf.json`:
```json
{
  "headers": {
    "cors": "https:\/\/.*test-app.com"
  }
}
```

```sh
regexlint go conf.json headers.cors
```

#### Check regex from .txt file by row number
```text
string
https:\/\/.*test-app.com
```

```sh
regexlint go file.txt row-1
```
