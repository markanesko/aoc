use std::fs::read_to_string;

fn check_increasing(array: &Vec<u8>) -> bool {   
    for i in 0..array.len() - 1 {
        let distance: i16 = (array[i+1] as i16 - array[i] as i16).abs();

        if !(distance >= 1 && distance <= 3) {
            return  false;
        }

        if array[i] >= array[i + 1] {
            return  false;
        }
        
    }
    true
}

fn check_decreasing(array: &Vec<u8>) -> bool {
    for i in 0..array.len() - 1 {
        let distance: i16 = (array[i] as i16 - array[i + 1] as i16).abs();

        if !(distance >= 1 && distance <= 3) {
            return  false;
        }

        if array[i] <= array[i + 1] {
            return  false;
        }
        
    }
    true
}

fn check_safety(array: &Vec<u8>) -> bool {
    if array[0] >= array[1] {
        let safe: bool = check_decreasing(&array);
        
        return  safe;
    }

    let safe = check_increasing(&array);
    return  safe;
}

fn main() {
    
    let mut safe_count: u16 = 0;

    let content = read_to_string("input").unwrap();

    for line in content.lines() {
        let levels: Vec<u8> = line
            .split_whitespace()
            .filter_map(|s: &str| s.parse::<u8>().ok())
            .collect();

        if levels.is_empty() {
            continue;
        }

        let safe: bool = check_safety(&levels);

        if safe {
            safe_count += 1;
            continue;
        }

        // there must be some better way instead of generating n - 1 arrays for each array
        let arrays: Vec<Vec<u8>> = levels
            .iter()
            .enumerate()
            .map(|(i, _)| {
                levels
                    .iter()
                    .enumerate()
                    .filter_map(|(j, &value)| if j != i { Some(value) } else { None })
                    .collect()
            })
            .collect();


        for array in arrays.iter() {
            let safe: bool = check_safety(array);
            if safe {
                safe_count += 1;
                break;
            }
        }

    }

    println!("safe_count = {}", safe_count);

    
}
