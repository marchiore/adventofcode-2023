use std::{fs, collections::HashMap};

const NUMBERS: [&str; 9] = ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"];

fn main() {    
    let contents = fs::read_to_string("/Users/matheus/Developer/adventofcode/day1/src/input.txt")
        .expect("Something went wrong reading the file");
    
    let mut total: i32 = 0;
    for line in contents.lines() {
        let mut numbers_in_lines: HashMap<i32, i32>  = HashMap::new();
        let numbers_from_line = get_calibration_from_lines(line, &mut numbers_in_lines);
        for (key, value) in numbers_in_lines.iter() {
            println!("Key: {}, Value: {}", key, value);
        }
        total = total + numbers_from_line;
        println!("{} - {} total: {}", line, numbers_from_line, total);
    }

}

fn get_calibration_from_lines(line: &str, number_in_lines: &mut HashMap<i32, i32>) -> i32 {

    for number in NUMBERS.iter() {
        match line.find(number) {
            Some(index) => {
                number_in_lines.insert(index as i32, get_number_from_string(number));
            },
            None => ()
        }
    }
    for (index, char) in line.chars().enumerate() {

        if char.is_digit(10) {
            number_in_lines.insert(index as i32, char.to_digit(10).unwrap() as i32);
        }
    }

    let primeiro: Option<&i32>;
    if let Some(value) = number_in_lines.get(number_in_lines.keys().min().unwrap()) {
        primeiro = Some(value);
    } else {
        primeiro = Some(&0);
    }

    let ultimo: Option<&i32>;
    if let Some(value) = number_in_lines.get(number_in_lines.keys().max().unwrap()) {
        ultimo = Some(value);
    } else {
        ultimo = Some(primeiro.unwrap());
    }

    return format!("{}{}", primeiro.unwrap_or(&0), ultimo.unwrap_or(primeiro.unwrap())).parse::<i32>().unwrap()
}

fn get_number_from_string(str: &str) -> i32 {
    let mut m = HashMap::new();

    m.insert("one", 1);
    m.insert("two", 2);
    m.insert("three", 3);
    m.insert("four", 4);
    m.insert("five", 5);
    m.insert("six", 6);
    m.insert("seven", 7);
    m.insert("eight", 8);
    m.insert("nine", 9);

    return *m.get(str).unwrap_or(&0);

}

fn find_all_occurrence_indices(text: &str, pattern: &str) -> Vec<usize> {
    text.match_indices(pattern)
        .map(|(index, _)| index)
        .collect()
}
