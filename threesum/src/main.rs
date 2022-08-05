use threesum::three_sum;

fn main() {
    let nums: Vec<i32> = vec![-1,0,1,2,-1,-4];
    println!("input array:  {:?}", nums);
    // check the value
    let expected_result: Vec<Vec<i32>> = three_sum(nums);

    println!("the result is : {:?}", expected_result);
    assert_eq!(vec![vec![-1,-1,2], vec![-1,0,1]] as Vec<Vec<i32>>, expected_result)
}
