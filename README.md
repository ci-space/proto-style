# protostyle

[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/ci-space/protostyle/master/LICENSE)

**protostyle** - [protolint](https://github.com/yoheimuta/protolint) plugin for lint team style

## Usage

1. Download [archive](https://github.com/ci-space/protostyle/releases/download/v0.1.0/protostyle-darwin-amd64.zip) and extract **protostyle**
```shell
curl -L https://github.com/ci-space/protostyle/releases/download/v0.1.0/protostyle-darwin-amd64.zip -o protostyle.zip && \
  unzip protostyle.zip protostyle
```

2. Run **protolint** with **protostyle**:
```shell
protolint -plugin protostyle lint
```

## Rules

| Rule                                              | Fixable | Description                                                 |
|---------------------------------------------------|---------|-------------------------------------------------------------|
| PROTOSTYLE_LIST_MESSAGES_RESOURCE_NAME_PLURALIZED | ✅       | List request/response must have pluralized resource name    |
| PROTOSTYLE_RPC_WITH_HTTP                          | -       | Method must have http option                                |
| PROTOSTYLE_RPC_WITHOUT_RESOURCE_NAME              | ✅       | Method must not contain resource name                       |
| PROTOSTYLE_FIELD_WITH_BEHAVIOR                    | -       | Field must have behavior option (google.api.field_behavior) |
