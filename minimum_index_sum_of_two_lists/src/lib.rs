use std::collections::HashMap;

pub fn find_restaurant(list1: Vec<String>, list2: Vec<String>) -> Vec<String> {
    let mut indexes : HashMap<&String, usize>= HashMap::new();
    let mut sums : Vec<String> = vec!();

    let mut minimumsum = 2001;
    for val in list1.iter().enumerate(){
        indexes.entry(val.1).or_insert(val.0);
    }
    for val in list2.iter().enumerate(){
        if let Some(x) = indexes.get(val.1) {
            let indexsum = x + val.0;
            if indexsum == minimumsum{
                sums.push(String::from(val.1.clone()));
            } else if indexsum < minimumsum {
                minimumsum = indexsum;
                sums = vec![String::from(val.1.clone())];
            }
        }
    }
    return sums
}