# regexlint

regexlint - simple application for regex validation, e.g. in CI pipelines for checking trust domains from configuration files

**Supported syntax**
* Go
* PCRE (PERL, PHP)

## Use cases

### Check regex from command-line

regex: `invalid-regex())`

command
```sh
regexlint pcre "invalid-regex())"
```

#### Check regex from yaml file with name `conf.yaml`:

_conf.yaml_
```yaml
headers:
  cors: https:\/\/.*test-app.com
```

command
```sh
regexlint go conf.yaml headers.cors
```

#### Check regex from json file with name `conf.json`:

_file.json_
```json
{
  "headers": {
    "cors": "https:\/\/.*test-app.com"
  }
}
```

command:
```sh
regexlint pcre conf.json headers.cors
```

#### Check regex from .txt file by row number

_file.txt_
```text
string
https:\/\/.*test-app.com
```

command
```sh
regexlint pcre file.txt row-1
```

#### Check many regexes from JSON/YAML file

_file.yaml_
```yaml
headers:
  cors1: https:\/\/.*test-app.com
  cors2: https:\/\/.*test-app.com
```

command
```sh
regexlint go conf.yaml headers.cors1,headers.cors2
```

#### Check all rows from .txt file

_file.txt_
```text
string
https:\/\/.*test-app.com
```

command
```sh
regexlint pcre file.txt row-all
```
