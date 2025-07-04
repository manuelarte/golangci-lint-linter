# Rules

## Experimental: GC001: Linters are sorted alphabetically

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

`--fix` option available.

## Experimental: GC002: Settings are sorted alphabetically

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

`--fix` option available.

## Experimental: GC013: Fields are sorted by: default, enabled, disabled, settings, exclusions

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

## Experimental: GC021: Disabled linters have a reason

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
    - asasalint  # to be checked in code reviews.
    - ...
  settings:
    ...
```

</td></tr>

</tbody>
</table>
