# hipchat-delay

Make better timing for your hipchat messages. 

# Features

### `-in 5m`

### `-at 12:36`

NB: if the time has already passed, next day will be used. E.g. `-at 01:00` on 1st of july will become `02-06 01:00`.

### `-silence=5m`
 
Will prevent your post to appear in the middle of a lively discussion other topic.

# Installation 

`go get github.com/mgurov/hipchat-delay`

# Configuration

### HIPCHAT_AUTH_TOKEN

A HipChat API token with message posting and reading (for silence period) rights is needed. 
Pass it via environmental variable `HIPCHAT_AUTH_TOKEN` or `--auth` command line flag.    