# protostyle

[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/ci-space/protostyle/master/LICENSE)

protostyle - [protolint](https://github.com/yoheimuta/protolint) plugin for lint team style

## Rules

| Rule                                              | Fixable | Description                                                 |
|---------------------------------------------------|---------|-------------------------------------------------------------|
| PROTOSTYLE_LIST_MESSAGES_RESOURCE_NAME_PLURALIZED | ✅       | List request/response must have pluralized resource name    |
| PROTOSTYLE_RPC_WITH_HTTP                          | -       | Method must have http option                                |
| PROTOSTYLE_RPC_WITHOUT_RESOURCE_NAME              | ✅       | Method must not contain resource name                       |
| PROTOSTYLE_FIELD_WITH_BEHAVIOR                    | -       | Field must have behavior option (google.api.field_behavior) |
