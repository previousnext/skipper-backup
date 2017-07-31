Skipper - Backups
=================

[![CircleCI](https://circleci.com/gh/previousnext/skipper-backup.svg?style=svg)](https://circleci.com/gh/previousnext/skipper-backup)

A small utility for chaining backup strategies into s3 storage. 

## Usage

```bash
# Create an archive of the app.
$ drush archive-dump --overwrite --destination=/tmp/drupal.tar.gz

# Send archive to s3 bucket with tokenised destination.
$ backup /tmp/drupal.tar.gz BUCKET "%FREQUENCY%/drupal/%TIMESTAMP%-dev.tar.gz"
```

## Credentials

* AWS environment variables `AWS_SECRET_ACCESS_KEY`, `AWS_ACCESS_KEY_ID`
* AWS named profile via environment variable `AWS_PROFILE`

## Placeholders

* `%FREQUENCY%` - 3 possible values: **monthly** on 1st of month, **weekly** on Sunday, **daily** all other times.
* `%TIMESTAMP%` - Current time formatted as **2006-01-02_15-04-05**

## Development

Ensure tests pass.

```
make test
```
