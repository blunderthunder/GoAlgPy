use bricks_falling_when_hit::hit_bricks;

fn main() {
    let grids: Vec<Vec<i32>> = vec![vec![1,0,0,0], vec![1,1,1,0]];
    let hits: Vec<Vec<i32>> = vec![vec![1,0]];
    println!("the result is {:?}", hit_bricks(grids, hits));

    let grids: Vec<Vec<i32>> = vec![vec![1,0,0,0],vec![1,1,0,0]];
    let hits: Vec<Vec<i32>> = vec![vec![1,1], vec![1,0]];
    println!("the result is {:?}", hit_bricks(grids, hits));

    let grids: Vec<Vec<i32>> = vec![vec![1,0,1],vec![1,1,1]];
    let hits: Vec<Vec<i32>> = vec![vec![0,0],vec![0,2],vec![1,1]];
    println!("the result is {:?}", hit_bricks(grids, hits));
    
}
