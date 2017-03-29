Skipper - Backups
=================

A small utility for chaining backup strategies.

## Usage

```bash
$ drush archive-dump --overwrite --destination=/tmp/drupal.tar.gz
$ backup /tmp/drupal.tar.gz BUCKET "%FREQUENCY%/drupal/%TIMESTAMP%-dev.tar.gz"
```

