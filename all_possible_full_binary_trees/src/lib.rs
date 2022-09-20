// Definition for a binary tree node.
#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
  pub val: i32,
  pub left: Option<Rc<RefCell<TreeNode>>>,
  pub right: Option<Rc<RefCell<TreeNode>>>,
}
// 
impl TreeNode {
  #[inline]
  pub fn new(val: i32) -> Self {
    TreeNode {
      val,
      left: None,
      right: None
    }
  }
}
use std::collections::HashMap;
use std::rc::Rc;
use std::cell::RefCell;

pub struct Solution;

impl Solution {
    pub fn all_possible_fbt(n: i32) -> Vec<Option<Rc<RefCell<TreeNode>>>> {
      let mut cache = HashMap::from([
        (0, vec![]),
        (1, vec![Some(Rc::new(RefCell::new(TreeNode::new(0))))])
      ]);

      fn backtrack(n: i32, hash: &mut HashMap<i32, Vec<Option<Rc<RefCell<TreeNode>>>>>) -> Vec<Option<Rc<RefCell<TreeNode>>>> {
        
        if hash.contains_key(&n){
          return hash[&n].clone()
        }

        let mut res: Vec<Option<Rc<RefCell<TreeNode>>>> = Vec::new();

        for l in 0..n{
          let r = n - 1 - l;
          for lefttree in backtrack(l, hash){
            for righttree in backtrack(r, hash){
              
              res.push(Some(Rc::new(RefCell::new(TreeNode{val: 0, left: lefttree.clone(), right: righttree}))))
            }
          }
        }
        hash.insert(n, res.clone());
        return res
      }
      backtrack(n, &mut cache)
    }
}