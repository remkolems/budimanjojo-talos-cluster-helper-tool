package main

import "flag"

var minimal bool

var purge bool

var onlyVersions bool
var version string

var trimRegistry bool
var trimSha256 bool
var trimTag bool

func init() {
	flag.BoolVar(&minimal, "minimal", false, "output minimal json consisting of only the org/repo, e.g. 'siderolabs/amd-ucode' -- flag is mutually exclusive to '-trimRegistry', '-trimSha256', and '-trimTag'.")

	flag.BoolVar(&purge, "purgecache", false, "purge the cache file")

	flag.BoolVar(&onlyVersions, "onlyVersions", false, "only output the versions, e.g. 'v1.5.5' -- flag is mutually exclusive to '-trimRegistry', '-trimSha256', and '-trimTag'.")
	flag.StringVar(&version, "version", "", "only output the version specified, e.g. 'v1.5.5' -- flag is mutually exclusive to '-trimRegistry', '-trimSha256', and '-trimTag'.")

	flag.BoolVar(&trimRegistry, "trimRegistry", false, "trim the sha256 suffix e.g. 'ghcr.io/siderolabs/amd-ucode:v1.2.0@sha256:1234567890abcdef' -> 'ghcr.io/siderolabs/amd-ucode:v1.2.0'")
	flag.BoolVar(&trimSha256, "trimSha256", false, "trim the registry prefix, e.g. 'ghcr.io/siderolabs/extensions/siderolabs/amd-ucode:v1.2.0@sha256:1234567890abcdef' -> 'siderolabs/amd-ucode:v1.2.0@sha256:1234567890abcdef'")
	flag.BoolVar(&trimTag, "trimTag", false, "trim the tag, e.g. 'ghcr.io/siderolabs/extensions/siderolabs/amd-ucode:v1.2.0@sha256:1234567890abcdef' -> 'ghcr.io/siderolabs/extensions/siderolabs/amd-ucode@sha256:1234567890abcdef'")

	flag.Parse()
}
