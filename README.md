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

## Packages

### i18n

The i18n save all the info of i18n system. It provides a completely i18n utils and tools. It contains file operator and
web support. You can use it as an i18n server. It provides some web handler to help i18n value change.

An i18n message contain the language, namespace and code value. The `mocker` provide two method to use i18n tools.

- web application
- sdk code

About the i18n system doc, see [i18n docs](docs/en/i18n/README.md).

## Project layout

### Base layout

About the package settings, the `mocker` project is in standard golang-pro layout. The
cmd dir save all main file(The `mocker` contain too many applications). All the package
will define in same package path which are same useful.

### dir/cmd