impl Solution {
    pub fn reverse_bits(mut x: u32) -> u32 {
        (0..32).into_iter().fold(0, |mut ans, _| {
            ans = (ans << 1) | (x & 1);
            x >>= 1;
            ans
        })
    }
}

pub struct Solution;