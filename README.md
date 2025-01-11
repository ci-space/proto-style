# proto-style

[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/ci-space/proto-style/master/LICENSE)

proto-style - plugin for [protolint](https://github.com/yoheimuta/protolint)

## Rules

| Rule                                      | Fixable | Description                                              |
|-------------------------------------------|---------|----------------------------------------------------------|
| PROTOSTYLE_LIST_MESSAGES_PLURAL_NAME_RULE | ✅       | List request/response must have pluralized resource name |
| PROTOSTYLE_RPC_WITH_HTTP_RULE             | -       | method must have http option                             |
| PROTOSTYLE_RPC_WITHOUT_RESOURCE_NAME_RULE | ✅       | method must not contain resource name                    |
