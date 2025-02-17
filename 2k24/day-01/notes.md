# notes day-01

## link with reading the file

https://doc.rust-lang.org/rust-by-example/std_misc/file/read_lines.html

```rust
use std::fs::read_to_string;

fn read_lines(filename: &str) -> Vec<String> {
    read_to_string(filename) 
        .unwrap()  // panic on possible file-reading errors
        .lines()  // split the string into an iterator of string slices
        .map(String::from)  // make each slice into a string
        .collect()  // gather them together into a vector
}
```

### Read into more efficient approach with reading file that is explained on the link