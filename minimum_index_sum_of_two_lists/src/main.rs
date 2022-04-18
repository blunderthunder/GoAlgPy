use minimum_index_sum_of_two_lists::find_restaurant;

fn main() {
    let v1: Vec<String> = vec!["Shogun","Tapioca Express","Burger King","KFC"].iter().map(|x| x.to_string()).collect();
    let v2: Vec<String> = vec!["Piatti","The Grill at Torrey Pines","Hungry Hunter Steakhouse","Shogun"].iter().map(|x| x.to_string()).collect();

    print!("minimum index sum of lists {:?} and {:?} is : ", v1 , v2);
    let val: Vec<String> = find_restaurant(v1, v2);

    println!("{:?}", val);
}