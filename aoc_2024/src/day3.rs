use regex::Regex;
use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

pub fn run() {
    let mut sum: i64 = 0;
    if let Ok(lines) = read_lines("./inputs/day3.txt") {
        for line in lines.flatten() {
            let re = Regex::new(r"mul\((?<val1>[0-9]+),(?<val2>[0-9]+)\)").unwrap();
            let vals: Vec<(i32, i32)> = re
                .captures_iter(line.as_str())
                .map(|caps| {
                    let val1 = caps.name("val1").unwrap().as_str().parse().unwrap();
                    let val2 = caps.name("val2").unwrap().as_str().parse().unwrap();
                    (val1, val2)
                })
                .collect();
            for val in vals {
                let mult = val.0 * val.1;
                sum += mult as i64;
            }
        }
    }
    println!("total: {sum}")
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where
    P: AsRef<Path>,
{
    let file = File::open(filename).expect("error opening the file");
    Ok(io::BufReader::new(file).lines())
}
