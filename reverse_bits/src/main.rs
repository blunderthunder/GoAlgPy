use reverse_bits::*;

fn main() {
    let given: u32 = "00000010100101000001111010011100".parse::<u32>().unwrap();
    let expected: u32 = "00111001011110000010100101000000".parse::<u32>().unwrap();

    let result = Solution::reverse_bits(given);

    assert_eq!(expected, result, "Expected value dosn't match result .");
}
