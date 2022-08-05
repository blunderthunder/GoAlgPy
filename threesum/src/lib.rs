pub fn three_sum(nums: Vec<i32>) -> Vec<Vec<i32>> {
    let mut result: Vec<Vec<i32>> = Vec::new();

    let mut nums = nums.clone();
    nums.sort();

    let n = nums.len();

    if n < 3{
        return vec![] as Vec<Vec<i32>>
    }
    if n == 3{
        if ( nums[0] + nums[1] + nums[2] ) == 0{
            result.push(vec![nums[0], nums[1], nums[2]]);
        }
        return result
    }

    // if all the condition fail use pointer method
    for num1_idx in 0usize..(n-2){
        if num1_idx > 0 && ( nums[num1_idx] == nums[num1_idx - 1]) {
            continue
        }

        let mut num2_idx: usize = num1_idx + 1;
        let mut num3_idx: usize = n - 1;
        while num2_idx < num3_idx {
            let sum = nums[num1_idx] + nums[num2_idx] + nums[num3_idx];

            if sum == 0{
                // add to the triplets
                result.push(vec![nums[num1_idx], nums[num2_idx], nums[num3_idx]]);

                num3_idx = num3_idx - 1;

                // skip all the duplicates from right
                while ( num2_idx < num3_idx ) && (nums[num3_idx] == nums[num3_idx + 1]){
                    num3_idx = num3_idx - 1;
                }
            } else if sum > 0{
                num3_idx = num3_idx -1;
            } else {
                num2_idx = num2_idx + 1;
            }
        }
    }
    result
}