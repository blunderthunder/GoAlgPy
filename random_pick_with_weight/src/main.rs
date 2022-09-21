use random_pick_with_weight::Solution;

fn main() {
    let input = vec![1,2,2,1,4];

    let sol = Solution::new(input);
    println!(" Pick index : {:?}", sol.pick_index())
}
