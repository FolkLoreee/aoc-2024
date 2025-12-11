use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

pub fn run () {

}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where
    P: AsRef<Path>,
{
    let file = File::open(filename).expect("error opening the file");
    Ok(io::BufReader::new(file).lines())
}
