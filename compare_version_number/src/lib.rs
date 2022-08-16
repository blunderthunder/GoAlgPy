use std::collections::HashMap;

impl Solution {
    pub fn compare_version(version1: String, version2: String) -> i32 {
        let mut mapv1: HashMap<usize, usize> = version1.split('.')
            .into_iter()
            .map(|x: &str| x.parse::<usize>().unwrap() as usize)
            .enumerate()
            .collect::<HashMap<usize, usize>>();

        let mut mapv2: HashMap<usize, usize> = version2.split('.')
            .into_iter()
            .map(|x: &str| x.parse::<usize>().unwrap() as usize)
            .enumerate()
            .collect::<HashMap<_, _>>();

        let longest_len: usize;

        if mapv1.len() > mapv2.len() {
            longest_len = mapv1.len();
        } else {
            longest_len = mapv2.len();
        }
        for idx in 0..longest_len {
            if let Some(_) = mapv1.get(&idx) {
                // do nothing 
            } else {
                mapv1.insert(idx, 0);
            }
            if let Some(_) = mapv2.get(&idx) {
                // do nothing
            } else {
                mapv2.insert(idx, 0);
            }

            if mapv1[&idx] > mapv2[&idx] {
                return 1
            } else if mapv2[&idx] > mapv1[&idx] {
                return -1
            } else {
                continue;
            }
        }

        0
    }
}

pub struct Solution;