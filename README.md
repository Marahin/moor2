# Moor
  
[![CircleCI](https://circleci.com/gh/Marahin/moor2.svg?style=svg)](https://circleci.com/gh/Marahin/moor2)  

**M**edium D**oor** - **moor** is a WWW proxy that fetches data from given URL (bypassing CORS and `while(1)` anti-JSON-hijack trap). 

**It was primarily created for fetching data from Medium-based blogs**, where you do not have any kind of API, and instead you must retrieve data like Mediums' frontend does: from `%PATH%/?format=json`. There were two caveats though:  

* CORS did not allow requests from non-medium-dot-com domain,
* it begins with `])}while(1);</x>`

For that **moor** has come to life.

moor2 is different from the original moor by being serverless (prepared to host on zeit.co). This way it should be cheaper and require even less maintenance than before. 

## Use cases → How can I use it?

That's up to you. I use it on [my website](http://marahin.pl) to get my latest post title and link :-).

## Installation

1. Clone this repository
2. Configure [https://zeit.co](zeit.co) project
3. Deploy and have fun

## Configuration

### Token

Just to be extra sure, moor2 requires you to set up a token. If you try to request a proxy for a URL without a token, or with an invalid one, you will get 401 Unauthorized. 

The token is fetched from `TOKEN` environment variable.

### CORS

In order to prevent malicious requests from third parties there is CORS support implemented.

In order to set CORS domain use `ALLOW_ORIGIN` environment variable (`ALLOW_ORIGIN=*`).

### Ignored endpoints

You can set ignored endpoints (ones that will NOT be fetched) in [generic_definitions.go](moor/generic_definitions.go#L13).
  
### Blocked characters amount

Blocked characters amount is the amount of characters that prefix the JSON output. It's default value can be seen in [moor/generic_definitions.go](moor/generic_definitions.go#L8) but you can also overwrite it using `MOOR_BLOCKER_CHARACTERS_AMOUNT` environment variable (as seen in [moor/http_client.go](moor/http_client.go#L18)). Default value is enough to remove Medium's trap.

## Usage

```
GET your.moor.instance.now.sh/api/medium?url=URL_ENCODE(URL_TO_FETCH)&token=token
```

## Contributing

If you wish to contribute please create a pull request or an issue.