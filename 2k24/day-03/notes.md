# notes

```rust
fn read_lines_one_by_one(filename: &str) -> Vec<Vec<char>> {
    let mut result = Vec::new();

    for line in read_to_string(filename).unwrap().lines() {
        result.push(line.to_string().chars().collect()) 
    }

    return result
}
```

You cannot index rust string since it's UTF-8 so it means that there is no guarantee that every element is of the same size. If that is what you want (i do in this) you can change `String` to `Vec<char>` and with that you allow yourself the option to use array indexing logic.  (`line.to_string().chars().collect()`)