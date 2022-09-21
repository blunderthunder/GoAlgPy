pub struct Solution {
  prefixes: Vec<i32>,
  total: i32
}


/** 
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */

use rand::Rng;

impl Solution {

    pub fn new(w: Vec<i32>) -> Self {
        let mut total: i32 = 0;
        let mut prefixes: Vec<i32> = Vec::new();

        for weight in w.into_iter(){
          total = total +  weight;
          prefixes.push(total)
        }

        return Solution { prefixes: prefixes, total: total }
    }
    
    pub fn pick_index(&self) -> i32 {
      let mut rng = rand::thread_rng();
      let target = rng.gen_range(0, self.total);

      let mut l = 0;
      let mut r: i32 = self.prefixes.len() as i32;

      while l <= r{
        let mid = ( l + r ) / 2;
        if self.prefixes.get(mid as usize).unwrap().clone() <= target {
          l = mid + 1
        } else {
          r = mid - 1
        }
      }
      l
    }
}