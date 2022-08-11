pub fn max_depth_after_split(seq: String) -> Vec<i32> {
    let mut ans: Vec<i32> = Vec::new();
    let mut depth: i32 = 0;

    for elem in seq.chars() {
        if elem == '('{
            ans.push(depth % 2);
            depth += 1;
        } else {
            depth -= 1;
            ans.push(depth % 2)
        }
    }

    ans
}