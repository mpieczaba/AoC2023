fn main() {
    const DIGITS: [&str; 9] = [
        "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
    ];

    const FILE: &str = include_str!("input.txt");

    let mut result_part_one = 0;
    let mut result_part_two = 0;

    FILE.lines().for_each(|line| {
        let (mut first_index, mut first_value) = line
            .chars()
            .enumerate()
            .find(|(_, c)| c.is_digit(10))
            .and_then(|(i, c)| Some((i, c.to_digit(10).unwrap())))
            .expect("at least one digit expected");

        let (mut last_index, mut last_value) = line
            .chars()
            .rev()
            .enumerate()
            .find(|(_, c)| c.is_digit(10))
            .and_then(|(i, c)| Some((line.len() - i - 1, c.to_digit(10).unwrap())))
            .expect("at least one digit expected");

        result_part_one += 10 * first_value + last_value;

        DIGITS.iter().enumerate().for_each(|(i, s)| {
            if let Some(m) = line.find(s) {
                if m < first_index {
                    first_index = m;
                    first_value = i as u32 + 1;
                }
            };

            if let Some(m) = line.rfind(s) {
                if m > last_index {
                    last_index = m;
                    last_value = i as u32 + 1;
                }
            };
        });

        result_part_two += 10 * first_value + last_value;
    });

    println!("Part one: {}", result_part_one);
    println!("Part two: {}", result_part_two);
}
