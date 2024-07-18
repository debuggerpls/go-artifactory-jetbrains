# Download Jetbrains IDEs and Plugins using Artifactory as Remote Proxy

For air-gapped environments there is:
https://www.jetbrains.com/help/idea/fully-offline-mode.html#set_up_offline

The problem is that is requests from different URLs, thus you cannot easily use it with Artifactory set as a proxy. This package should replace the `jetbrains-clients-downloader` for usage in combination with Artifactory's Remote Proxy.

## JSON structure
Its an array of products. Each product has following keys:
```
$ curl -s https://data.services.jetbrains.com/products | jq -r '.[0] | to_entries | .[] | "\(.key): \(.value | type)"'
code: string
name: string
releases: array
forSale: boolean
productFamilyName: string
additionalLinks: array
intellijProductCode: string
alternativeCodes: array
salesCode: string
link: string
description: string
tags: array
types: array
categories: array
distributions: object
```

```
$ cat products | jq '.[] | select(.code=="RR") | .distributions'
{
  "linux": {
    "name": "Linux",
    "extension": "tar.gz"
  },
  "macM1": {
    "name": "macOS Apple Silicon",
    "extension": "dmg"
  },
  "mac": {
    "name": "macOS",
    "extension": "dmg"
  },
  "windows": {
    "name": "Windows",
    "extension": "exe"
  },
  "windowsARM64": {
    "name": "Windows ARM64",
    "extension": "exe"
  },
  "linuxARM64": {
    "name": "Linux ARM64",
    "extension": "tar.gz"
  },
  "thirdPartyLibrariesJson": {
    "name": "Used Third-Party Libraries",
    "extension": "third-party-libraries.json"
  }
}
```

```
$ cat products | jq '.[] | select(.code=="RR") | .releases[0] | to_entries | .[] | "\(.key): \(.value | type)"'
"date: string"
"type: string"
"version: string"
"majorVersion: string"
"patches: object"
"downloads: object"
"licenseRequired: boolean"
"build: string"
"whatsnew: string"
"uninstallFeedbackLinks: object"
```

```
$ cat data/products | jq '.[] | select(.code=="RR") | .releases[0].downloads.linux | to_entries | .[] | "\(.key): \(.value | type)"'
"link: string"
"size: number"
"checksumLink: string"
```
