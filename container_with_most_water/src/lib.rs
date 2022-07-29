
pub fn max_area(height: Vec<i32>) -> i32 {
    match height.len() {
        0 | 1 => 0,
        2 => {
            if height[1] < height[0]{
                return height[1]
            } else {
                return height[0]
            }
        },
        _ => {
            let mut ans: i32 = 0;
            let mut startidx: usize = 0;
            let mut endidx: usize = height.len() - 1;

            loop {
                if startidx >= endidx{
                    break;
                }

                if height[startidx] > height[endidx]{
                    let newans = height[endidx] * ( endidx - startidx ) as i32;
                    if newans > ans{
                        ans = newans;
                    }
                    endidx = endidx - 1
                } else {
                    let newans = height[startidx] * ( endidx - startidx ) as i32;
                    if newans > ans{
                        ans = newans;
                    }
                    startidx = startidx + 1
                }
            }
            ans
        }
    }
}
