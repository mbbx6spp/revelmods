# assets Revel module

The assets Revel module will contain a useful set of template functions and
related convenience functions to manage:
* external CDN served assets (for production env)
* internal aggregated assets
* external local assets (for development env)

## Usage

To use this module you should add the following into your Revel app's
`conf/app.conf` file:

```
module.assets=github.com/mbbx6spp/revelmods/modules/assets
```


