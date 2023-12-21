use std::{fs, collections::HashMap};

const NUMBERS: [&str; 9] = ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"];

fn main() {    
    let contents = fs::read_to_string("input.txt")
        .expect("Something went wrong reading the file");
    
    let mut total: i32 = 0;
    for line in contents.lines() {
        let mut numbers_in_lines: HashMap<i32, i32>  = HashMap::new();
        let numbers_from_line = get_calibration_from_lines(line, &mut numbers_in_lines);

        total = total + numbers_from_line;
        
    }
    println!("total: {}", total);

}

fn get_calibration_from_lines(line: &str, number_in_lines: &mut HashMap<i32, i32>) -> i32 {

    for number in NUMBERS.iter() {
        for (index, _) in line.match_indices(number) {
            number_in_lines.insert(index as i32, get_number_from_word(number));
        }
    }
    for (index, char) in line.chars().enumerate() {

        if char.is_digit(10) {
            number_in_lines.insert(index as i32, char.to_digit(10).unwrap() as i32);
        }
    }

    let primeiro = number_in_lines.get(&number_in_lines.keys().min().cloned().unwrap_or(0)).unwrap_or(&0);
    let ultimo = number_in_lines.get(&number_in_lines.keys().max().cloned().unwrap_or(0)).unwrap_or(primeiro);

    format!("{}{}", primeiro, ultimo).parse::<i32>().unwrap()
}

fn get_number_from_word(word: &str) -> i32 {
    let mut number_map: HashMap<&str, i32> = HashMap::new();

    number_map.insert("one", 1);
    number_map.insert("two", 2);
    number_map.insert("three", 3);
    number_map.insert("four", 4);
    number_map.insert("five", 5);
    number_map.insert("six", 6);
    number_map.insert("seven", 7);
    number_map.insert("eight", 8);
    number_map.insert("nine", 9);

    number_map.get(word).cloned().unwrap_or_default()   
}

