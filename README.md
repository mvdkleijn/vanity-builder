# vanity-builder

Simple tool that takes a YAML based config and generates a set of static HTML files
that can be used for vanity URL hosting.

## Development notes

go-import: <vanityDomain>/<moduleName> git https://github.com/<user or org>/<repo>.git

go-source: content="<prefix> <homepage> <directory template> <file template>"
where:
  prefix  : <vanityDomain>/<moduleName>
            (import path corresponding to repo root)
  homepage: <homepageURL>
            (url of repos homepage where "_" signifies no homepage url)
  directory template: https://github.com/<user or org>/<repo>/tree/<branch>{/dir}
                      (url template for listing files in pkg)
  file template     : https://github.com/<user or org>/<repo>/blob/<branch>{/dir}/{file}#L{line}
                      (url template for line in file)