//go:build dev

package main

import "time"

// keep-sorted start

const ENABLE_STATIC_ASSET_BROWSING = true
const PAGE_MAX_AGE = 3 * time.Second
const STATIC_ASSET_MAX_AGE = 1 * time.Minute

// keep-sorted end
