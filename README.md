# Table Format

Format a Markdown table to an ASCII or Markdown formatted table

## Example

All the following examples will use the following ill-formatted
table as input:

```
left aligned | centered | right aligned
:--|:-:|--:|
|a content| a long content here | 42
somtething bigger|and small here| 1337.45
```

- An ASCII table (`tableformat --output ascii`):

<pre>
┌───────────────────┬─────────────────────┬────────────────┐
│ left aligned      │      centered       │  right aligned │
├───────────────────┼─────────────────────┼────────────────┤
│ a content         │ a long content here │             42 │
├───────────────────┼─────────────────────┼────────────────┤
│ somtething bigger │   and small here    │        1337.45 │
└───────────────────┴─────────────────────┴────────────────┘
</pre>

- A Markdown table (`tableformat --output markdown`):

```
| left aligned      |      centered       |  right aligned |
|:------------------|:-------------------:|---------------:|
| a content         | a long content here |             42 |
| somtething bigger |   and small here    |        1337.45 |
```

## Development

### Install

```
go mod download
go build
```

### Testing

```
go test ./...
```

### Deployment

> :warning: Make sure the last version has been build before
> running the command. And don't forget to bump the VERSION
> file

```
./scripts/release.sh
```
