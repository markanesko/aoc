use std::fs::read_to_string;

// fn read_lines_one_by_one(filename: &str) -> Vec<String> {
//     let mut result = Vec::new();

//     for line in read_to_string(filename).unwrap().lines() {
//         result.push(line.to_string())
//     }

//     return result
// }

fn read_and_sort(filename: &str) -> (Vec<i32>, Vec<i32>) {
    let (mut vec1, mut vec2): (Vec<i32>, Vec<i32>)= read_to_string(filename)
        .unwrap()
        .lines()
        .filter_map(|line| {
            let parts : Vec<&str> = line.split_whitespace().collect();

            if parts.len() == 2 {
                let first: i32 = parts[0].parse::<i32>().ok()?;
                let second: i32 = parts[1].parse::<i32>().ok()?;
                Some((first, second))
            } else {
                None
            }
        })
        .unzip();

    vec1.sort(); 
    vec2.sort();

    (vec1, vec2)
}

fn main() {
    let (vec1, vec2) = read_and_sort("input");

    let mut sum: i32 = 0;

    for (index, _) in vec1.iter().enumerate() {
        let distance: i32 = vec1[index] - vec2[index];
        sum += distance.abs();
    }

    println!("sum = {}", sum);

    let mut last_num: i32 = 0;
    let mut last_count: i32 = 0;
    let mut similarity_sum: i32 = 0;

    for l_value in vec1.iter() {
        if *l_value == last_num {
            similarity_sum += last_count;
            continue;
        }

        last_num = *l_value;
        last_count = 0;

        for r_value in vec2.iter() {
            
            if l_value > r_value {
                continue;
            }
            if l_value == r_value {
                last_count += l_value;
            }
            if l_value < r_value {
                break;
            }

        }

        similarity_sum += last_count;
    }

    println!("similarity sum = {}", similarity_sum);
}
