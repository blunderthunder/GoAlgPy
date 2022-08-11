use maximum_nesting_depth_of_two_valid_parentheses::*;

fn main() {
    let given: String = String::from("(()())");
    println!("given input is : {}", given);

    let result: Vec<i32> = max_depth_after_split(given);

    println!("output is : {:?}", result )
}
