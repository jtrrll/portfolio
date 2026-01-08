//go:build !dev

package main

import "time"

// keep-sorted start

const PAGE_MAX_AGE = 5 * time.Minute
const STATIC_ASSET_MAX_AGE = 24 * time.Hour

// keep-sorted end
