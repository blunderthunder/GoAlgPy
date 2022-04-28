fn dfs(x: i32, y: i32, grid: &mut Vec<Vec<i32>>) -> i32 {
    let ( m, n) = ( grid.len() as i32, grid[0].len() as i32);
    let scale: Vec<(i32, i32)> = vec![(-1,0), (1, 0), (0, -1), (0, 1)];
    if grid[x as usize][y as usize] != 1 {
        return 0
    } else{
        grid[x as usize][y as usize] = 2;
        let mut ans = 1;
        for ( xx, yy ) in scale.iter().map(|(pairx, pairy)| (x + pairx, y + pairy)){
            if xx >= 0 && xx < m && yy >= 0 && yy < n{
                ans += dfs(xx, yy, grid);
            }
        }
        return ans
    }
}

fn is_stable(x: i32, y: i32, grid: &mut Vec<Vec<i32>>) -> bool{
    grid[x as usize][y as usize] += 1;
    if grid[x as usize][y as usize] <= 0{
        return false
    } else if ( x == 0 && grid[x as usize][y as usize] ==1 ) || ( grid[x as usize][y as usize] == 2) {
        return true
    } else{
        let neighbour = vec![(-1,0), (1, 0), (0, -1), (0, 1)];
        let (m,n) = ( grid.len() as i32, grid[0].len() as i32);
        neighbour.iter().map(|pair| (pair.0 +x, pair.1 +y)).any(|(xx, yy)| (xx >= 0 && xx < m && yy >= 0 && yy < n && grid[xx as usize][yy as usize] == 2))
    }
}

pub fn hit_bricks(grid: Vec<Vec<i32>>, hits: Vec<Vec<i32>>) -> Vec<i32> {
    let mut fall : Vec<i32> = Vec::new();
    let mut grids : Vec<Vec<i32>> = Vec::from(grid);

    for loc in &hits{
        grids[loc[0] as usize][loc[1] as usize] -= 1
    }
    for i in 0..grids[0].len(){
        dfs(0, i as i32, &mut grids);
    }
    for loc in hits.into_iter().rev(){
        if is_stable(loc[0], loc[1], &mut grids){
            fall.push(dfs(loc[0],loc[1], &mut grids)-1)
        }else {
            fall.push(0)
        }
    }
    fall.into_iter().rev().collect()
}
