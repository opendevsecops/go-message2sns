[![Codacy Badge](https://api.codacy.com/project/badge/Grade/b07a461e6a9a48fc84226baefff06423)](https://www.codacy.com/app/OpenDevSecOps/go-message2sns?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=opendevsecops/go-message2sns&amp;utm_campaign=Badge_Grade)
[![Follow on Twitter](https://img.shields.io/twitter/follow/opendevsecops.svg?logo=twitter)](https://twitter.com/opendevsecops)

# go-message2sns

Simple utility to send message via AWS SNS.

## Rationale

Although this task can be done with cURL, this command is meant to save you some troubles with the API and future-proof it.

## Getting Started

To install message2sns simply do:

```sh
$ go get github.com/opendevsecops/go-message2sns
```

Once the command is installed in your home go/bin folder execute it like this:

```sh
$ ~/go/bin/go-message2sns --help
```

To send message to slack simply do:

```sh
$ ~/go/bin/go-message2sns --message '{"text": "hello"}' --sns "arn:..."
```
