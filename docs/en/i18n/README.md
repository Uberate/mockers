# The i18n system

The package `i18n` is a part of `mocker` project.

The i18n package provides some method to operate the i18n message quickly and simple.

## The message

The message object is minimum operating unit. Is saves a message value in a language and namespace. In different
namespace, the code can be same. But in one namespace, the code must different.

## About the message file

The i18n system tools support the CSV file. And the csv file is sort by namespace and code. But I18n build from the csv
file will ignore the order.

### Why the csv file

CSV file is friendly to read and write for human. So, the backup of i18n and build from i18n is low frequency. So, the
i18n build should before the application run. 