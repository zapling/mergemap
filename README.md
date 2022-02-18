# mergemap

mergemap is a Go library to recursively merge JSON maps.

This is a fork of [peterbourgon/mergemap](github.com/peterbourgon/mergemap) which adds support for
different merge strategies.

## Behavior

mergemap performs a simple merge of the **src** map into the **dst** map. That
is, it takes the **src** value when there is a key conflict, this can be customised by configuring
another merge strategy.

When there is a conflicting key that represents a map in both src and dst, then mergemap recursively 
descends into both maps, repeating the same logic. The max recursion depth is set by **mergemap.MaxDepth**.


## Usage

```go
var m1, m2 map[string]interface{}
json.Unmarshal(buf1, &m1)
json.Unmarshal(buf2, &m2)

merged := mergemap.Merge(m1, m2)
```

If you need more customised behavior you can use `MergeWithConfig`

```go

payload1 := `{"my-key": "first-value"}`
payload2 := `{"my-key": "second-value"}`

var m1, m2 map[string]interface{}
json.Unmarshal(payload1, &m1)
json.Unmarshal(payload2, &m2)

config := map[string]interface{}{
    "my-key": mergemap.StrategyFirstValue,
}

merged := mergemap.MergeWithConfig(m1, m2, config)

fmt.Printf("%v", merged) // map[my-key:first-value]
```

There is a few merge strategies already configuerd, but you can easily add or remove ones by modifying the
`mergemap.DefaultMergeStrategies` map.

The currently supported strategies are:
- Last value (default)
- First value
- Minimum value
- Maximum value

See the test file for some pretty straightforward examples.

