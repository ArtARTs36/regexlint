# regexlint

regexlint - simple application for regex validation, e.g. in CI pipelines for checking trust domains

## Use cases

### Check go regex

```sh
regexlint go "invalid-regex())"
```

Check regex from yaml file with name `conf.yaml`:
```
headers:
  cors: https:\/\/.*test-app.com
```

```sh
regexlint go conf.yaml headers.cors
```
