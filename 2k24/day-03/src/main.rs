use std::{env::current_exe, fs::read_to_string};
use std::fmt;

static NUMBERS: &[char] = &['1', '2', '3', '4', '5', '6', '7', '8', '9', '0'];
static NUMBERS_AND_SPECIALS: &[char] = &['1', '2', '3', '4', '5', '6', '7', '8', '9', '0', ',', ')'];
// static MUL: &[char] = &['m', 'u', 'l'];
// static SPEC: &[char] = &['(', ')', ','];

static ALL: &[char] = &['1', '2', '3', '4', '5', '6', '7', '8', '9', '0', 'm', 'u', 'l', '(', ')', ','];

#[derive(Debug)]
enum State {
    M, 
    U,
    L,
    Open,
    Number,
    Coma,
    Closed,
    Wrong,
}
impl fmt::Display for State {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match self {
            State::M => write!(f, "M"),
            State::U => write!(f, "U"),
            State::L => write!(f, "L"),
            State::Open => write!(f, "Open"),
            State::Number => write!(f, "Number"),
            State::Coma => write!(f, "Coma"),
            State::Closed => write!(f, "Closed"),
            State::Wrong => write!(f, "Wrong"),
        }
    }
}
fn read_lines_one_by_one(filename: &str) -> Vec<Vec<char>> {
    let mut result = Vec::new();

    for line in read_to_string(filename).unwrap().lines() {
        result.push(line.to_string().chars().collect()) 
    }

    return result
}

// fn valdidate_memory(mem: &Vec<char>) -> u32 {
//     let sum_line_product: u32 = 0;

//     let mut filtered: Vec<char> = Vec::new();

//     let mut i: usize = 0;

//     let mut mul_start :bool = false;

//     while i < mem.len() {
//         if !ALL.contains(&mem[i]) {
//             mul_start = false;
//             i += 1;
//             continue;
//         }

//         if mem[i] != 'm' && !mul_start {
//             i += 1;
//             continue;
//         }

//         if mem[i] == ')' && mul_start {
//             mul_start = false;
//             filtered.push(mem[i]);
//             i += 1;
//             continue;
//         }
        
//         if !NUMBERS.contains(&mem[i]) && mul_start {
//             mul_start = false;
//             i += 1;
//             continue;
//         }

//         if mem[i] == 'm' && !mul_start {
//             if (i + 3) == mem.len() {
//                 i += 3;
//                 continue;
//             }
//             if mem[i + 1] != 'u' || mem[i + 2] != 'l' || mem[i + 3] != '('{
//                 i += 4;
//                 mul_start = false;
//                 continue;
//             }
 
//             mul_start = true;
//             filtered.push(mem[i]);
//             filtered.push(mem[i + 1]);
//             filtered.push(mem[i + 2]);
//             filtered.push(mem[i + 3]);
//             i+=4;
//             continue;
//         }

        

//         filtered.push(mem[i]);
//         i += 1;
//     }

//     let my_string: String = filtered.iter().collect();

//     println!("filtered = {:?}", my_string);

//     sum_line_product
// }

fn valdidate_memory_2(mem: &Vec<char>) -> u32 {
    let sum_line_product: u32 = 0;

    let mut current_state = State::Wrong;

    let mut i: usize = 0;

    let mut filtered: Vec<char> = Vec::new();

    let mut full_filtered: String = String::from("");


    while i < mem.len() {
        // println!("current_state = {:?} and char {}", current_state, mem[i]);

        match  current_state {
            State::M => {
                if mem[i] != 'u' {
                    i += 1;
                    current_state = State::Wrong;
                    filtered.clear();
                    continue;
                }
                filtered.push(mem[i]);
                current_state = State::U;
                i += 1;
            }
            State::U => {
                if mem[i] != 'l' {
                    i += 1;
                    current_state = State::Wrong;
                    filtered.clear();
                    continue;
                }
                filtered.push(mem[i]);
                current_state = State::L;
                i += 1;
            }
            State::L => {
                if mem[i] != '(' {
                    i += 1;
                    current_state = State::Wrong;
                    filtered.clear();
                    continue;
                }
                filtered.push(mem[i]);
                current_state = State::Open;
                i += 1;
            }
            State::Open => {
                if !NUMBERS.contains(&mem[i]) {
                    i += 1;
                    current_state = State::Wrong;
                    filtered.clear();
                    continue;
                }
                filtered.push(mem[i]);
                current_state = State::Number;
                i += 1;
            }
            State::Number => {
                if !NUMBERS_AND_SPECIALS.contains(&mem[i]) {
                    i += 1;
                    current_state = State::Wrong;
                    filtered.clear();
                    continue;
                }
                if mem[i] == ')' {
                    filtered.push(mem[i]);
                    i += 1;
                    current_state = State::Closed;
                    continue;
                }
                if mem[i] == ',' {
                    filtered.push(mem[i]);
                    i += 1;
                    current_state = State::Coma;
                    continue;
                }
                filtered.push(mem[i]);
                current_state = State::Number;
                i += 1;
            }
            State::Coma => {
                if !NUMBERS.contains(&mem[i]) {
                    i += 1;
                    current_state = State::Wrong;
                    filtered.clear();
                    continue;
                }
                filtered.push(mem[i]);
                current_state = State::Number;
                i += 1;
            }
            State::Closed => {
                let my_string: String = filtered.iter().collect();
                full_filtered = my_string + &full_filtered; 
                filtered.clear();

                if mem[i] != 'm' {
                    current_state = State::Wrong;
                    i += 1;
                    continue;
                }

                filtered.push(mem[i]);
                current_state = State::M;
                i += 1;

            }
            State::Wrong => {
                filtered.clear();
                if mem[i] != 'm' {
                    current_state = State::Wrong;
                    i += 1;
                    continue;
                }
               
                filtered.push(mem[i]);
                current_state = State::M;
                i += 1;
            }
        }

    }


    println!("filtered = {:?}", full_filtered);


    sum_line_product
}

fn main() {
    let lines: Vec<Vec<char>> = read_lines_one_by_one("input");

    for line in lines.iter() {

        let line_product = valdidate_memory_2(line);
    }
}
