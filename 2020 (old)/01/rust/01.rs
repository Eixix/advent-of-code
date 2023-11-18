use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    if let Ok(lines) = read_lines("./small.txt") {
        // Consumes the iterator, returns an (Optional) String
        for first in lines.iter() {
            let firstInt = first.unwrap().parse::<i32>().unwrap();
            for second in lines {
                let secondInt = second.unwrap().parse::<i32>().unwrap();
                    if firstInt + secondInt == 2020 {
                        println!("{}", (firstInt * secondInt));
                    }
                }
        }
    }
}

// The output is wrapped in a Result to allow matching on errors
// Returns an Iterator to the Reader of the lines of the file.
fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}