# vanity-builder

Simple tool that takes a YAML based config and generates a set of static HTML files
based on Go templates that can be used for vanity URL self-hosting.

**Example**

Use `code.vanderkleijn.net` instead of `github.com` for your modules.

**Why?**

- No renaming issues when (for whatever reason) you're moving away from Github to somewhere else;
- Code can still be hosted on Github (or other places);

_In other words:_ similar to gopkg.in but easily self-hostable.

## Development notes

go-import: `<vanityDomain>/<moduleName> git https://github.com/<user or org>/<repo>.git`

go-source: `content="<prefix> <homepage> <directory template> <file template>"`
Where:
  - prefix: `<vanityDomain>/<moduleName>`
    - The import path corresponding to the repo root
    - `<moduleName>` can include major version, e.g. `moduleName` or `moduleName/v2`
  - homepage: `<homepageURL>`
    - The url of the repo's homepage
    - The '_' character can be used to signify the absence of a homepage
  - directory template: `https://github.com/<user or org>/<repo>/tree/<branch>{/dir}`
    - An template used to create the URL for listing the files in the module
  - file template: `https://github.com/<user or org>/<repo>/blob/<branch>{/dir}/{file}#L{line}`
    - An template used to create the url that points to a line in a file