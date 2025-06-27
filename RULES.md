# Rules

## Style Linters

### Experimental: GCI001: Linters are sorted alphabetically

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

### Experimental: GCI002: Settings are sorted alphabetically

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

### Experimental: GCI013: Fields are sorted by, default, enabled, disabled

The expected order for the `linters` field is:

+ default
+ enable
+ disable
+ settings

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

## Documentation & Metadata

### Experimental: GCI020: Warn for settings of not enabled linter

### Experimental: GCI021: Disabled linters have a reason

### Experimental: GCI022: Linter Suggestions

    + gci with imports in standard, and then the rest.
    + suggestions to use revive and var-naming and filename.
    + replace interface{} to any.

## Miscellaneous

### Experimental: GCI041:  Catch flaws in configuration
