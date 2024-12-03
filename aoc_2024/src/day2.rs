use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

pub fn run() {
    let mut safe_counter = 0;
    if let Ok(lines) = read_lines("./inputs/day2.txt") {
        for line in lines.flatten() {
            let arr: Vec<i8> = line.split(' ').map(|i| i.parse().unwrap()).collect();
            if is_increasing(&arr) && differ_by_at_most_n(&arr, 1, 3) {
                safe_counter += 1;
            }
        }
    }
    println!("total: {safe_counter}")
}

fn is_increasing(input: &[i8]) -> bool {
    let mut last = input[0];
    let mut last_diff = i8::MAX;
    for i in input.iter().skip(1) {
        let curr = *i;
        let diff = curr - last;
        if last_diff == i8::MAX {
            last_diff = diff;
        }
        if last_diff >= 0 && diff < 0 || last_diff < 0 && diff >= 0 {
            println!("diff: {diff}, previous: {last_diff}");
            return false;
        }
        last = curr;
    }
    true
}

fn differ_by_at_most_n(input: &[i8], min_n: i8, max_n: i8) -> bool {
    let mut last = input[0];
    for i in input.iter().skip(1) {
        let curr = *i;
        let diff = (curr - last).abs();
        if diff < min_n || diff > max_n {
            println!("diff: {diff}");
            return false;
        }
        last = curr;
    }
    true
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where
    P: AsRef<Path>,
{
    let file = File::open(filename).expect("error opening the file");
    Ok(io::BufReader::new(file).lines())
}
