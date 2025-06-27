# Rules

## Experimental: GCI001: Linters are sorted alphabetically

Linters lists are expected to be in alphabetical order, e.g.

<table>
<thead><tr><th>❌ Bad</th><th>✅ Good</th></tr></thead>
<tbody>
<tr><td>

```yaml
linters:
  default: none
  enable:
    - zerologlint
    - asasalint  
```

</td><td>

```yaml
linters:
  default: none
  enable:
    - asasalint  
    - zerologlint
```

</td></tr>

</tbody>
</table>

## Experimental: GCI002: Settings are sorted alphabetically

Linter's settings are expected to be sorted alphabetically, e.g.

<table>
<thead><tr><th>❌ Bad</th><th>✅ Good</th></tr></thead>
<tbody>
<tr><td>

```yaml
linters:
  default: all
  settings:
    wsl:
      strict-append: false
    asasalint:
      exclude:
        - Append
```

</td><td>

```yaml
linters:
  default: all
  settings:
    asasalint:
      exclude:
        - Append
    wsl:
      strict-append: false

```

</td></tr>

</tbody>
</table>

## Experimental: GCI013: Fields are sorted by: default, enabled, disabled, exclusions

The expected order for the `linters` field is:

+ default
+ enable
+ disable
+ settings
+ exclusions

<table>
<thead><tr><th>❌ Bad</th><th>✅ Good</th></tr></thead>
<tbody>
<tr><td>

```yaml
linters:
  disable:
    - ...
  default: all
  settings:
    ...
```

</td><td>

```yaml

linters:
  default: all
  disable:
    - ...
  settings:
    ...
```

</td></tr>

</tbody>
</table>

## Experimental: GCI020: Warn for settings of not enabled linter

TODO

## Experimental: GCI021: Disabled linters have a reason

Linters that are disabled need to have a comment, e.g.:

<table>
<thead><tr><th>❌ Bad</th><th>✅ Good</th></tr></thead>
<tbody>
<tr><td>

```yaml
linters:
  default: all
  disable:
    - asasalint
    - ...
  settings:
    ...
```

</td><td>

```yaml
linters:
  default: all
  disable:
    - asasalint  # to be checked in code review.
    - ...
  settings:
    ...
```

</td></tr>

</tbody>
</table>

## Experimental: GCI022: Linter Suggestions

TODO

## Experimental: GCI041:  Catch flaws in configuration

TODO
