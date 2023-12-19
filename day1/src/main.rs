use std::fs;

fn main() {    
    let contents = fs::read_to_string("/Users/matheus/Developer/adventofcode/day1/day1/src/input.txt")
        .expect("Something went wrong reading the file");

    let total: i32 = contents.lines().map(get_calibration_from_lines).sum();

    println!("Total: {}", total);
}

fn get_calibration_from_lines(line: &str) -> i32 {

    let mut numbers: Vec<i32> = line.chars()
    .filter_map(|c| c.to_digit(10).map(|digit| digit as i32)).collect();

    let first = numbers.get(0).copied().unwrap_or_default();
    let last = numbers.last().copied().unwrap_or_default();

    format!("{}{}", first, last).parse().unwrap_or_default()
}
