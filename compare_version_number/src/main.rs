use std::str::FromStr;

use compare_version_number::*;

fn main() {
    let version1: String = String::from_str("0.1").unwrap();
    let version2: String = String::from_str("1.01").unwrap();

    // let solution  =  Solution{};

    let result: i32 = Solution::compare_version(version1, version2);

    assert!(
        result==-1,
        "for version1 = {} , version2 = {} expected result is {} but got {}", "0.1", "1.01", -1, result);
}
