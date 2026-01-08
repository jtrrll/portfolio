//go:build dev

package main

import "time"

// keep-sorted start

const PAGE_MAX_AGE = 3 * time.Second
const STATIC_ASSET_MAX_AGE = 1 * time.Minute

// keep-sorted end
