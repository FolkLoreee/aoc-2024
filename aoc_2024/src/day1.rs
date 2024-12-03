use crate::utils::file::read_lines;

pub fn run() {
    let mut dist1: Vec<i32> = Vec::new();
    let mut dist2: Vec<i32> = Vec::new();
    if let Ok(lines) = read_lines("./inputs/day1.txt") {
        for line in lines.flatten() {
            let parts = line.split(',');
            let arr: Vec<&str> = parts.collect();

            let num1: i32 = arr[0].parse().unwrap();
            let num2: i32 = arr[1].parse().unwrap();
            dist1.push(num1);
            dist2.push(num2);
        }
    }
    dist1.sort();
    dist2.sort();

    let mut sum: i32 = 0;
    for (i, num) in dist1.iter().enumerate() {
        let num2 = dist2[i];
        let diff = (num2 - num).abs();
        println!("{num2} - {num} = {diff}");
        sum += diff;
    }
    println!("sum: {sum}");
}
