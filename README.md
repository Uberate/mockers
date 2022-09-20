# Mocker for golang
Status:

[![code-test](https://github.com/Uberate/mockers/actions/workflows/code-test.yml/badge.svg)](https://github.com/Uberate/mockers/actions/workflows/code-test.yml)

![Stars](https://img.shields.io/github/stars/Uberate/mockers?label=Stars)
![Forks](https://img.shields.io/github/forks/Uberate/mockers?label=Forks)
![License](https://img.shields.io/github/license/Uberate/mockers?label=LICENSE)

Provide some tools to quickly create fake data.

## Environment define

- **Need Golang 1.18+**
- If you are using goland, please use **2022.2+**.

## Package define
### PKG
- collection: provide some collection struct to operator data.
- errors: provide a common error method to deal the errors. Support the i18n and http error code.
- event: provide a common event system in the application. It will support the RAFT event model.
- hash: provide some method and interface about the hash.
- i18n: provide a common i18n system.
- utils: the utils tools.