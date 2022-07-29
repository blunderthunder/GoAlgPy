use container_with_most_water::max_area;

fn main() {
    let result = max_area(Vec::from([1,8,6,2,5,4,8,3,7]));

    assert_eq!(result, 49);
}
