Skipper - Backups
=================

[![CircleCI](https://circleci.com/gh/previousnext/skipper-backup.svg?style=svg)](https://circleci.com/gh/previousnext/skipper-backup)

A small utility for chaining backup strategies.

## Usage

```bash
$ drush archive-dump --overwrite --destination=/tmp/drupal.tar.gz
$ backup /tmp/drupal.tar.gz BUCKET "%FREQUENCY%/drupal/%TIMESTAMP%-dev.tar.gz"
```

## Development

Ensure tests pass.

```
make test
```
