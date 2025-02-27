// Embeds static files in the package.
package static

import "embed"

//go:embed favicon.ico styles.css
var StaticFs embed.FS
