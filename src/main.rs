mod day1;

macro_rules! run_solution {
    ($day:ident) => {
        use $day::{solve_part1, solve_part2};

        let input = include_str!(concat!("../inputs/", stringify!($day), ".txt"));

        println!("{} p1 {:#?}", stringify!($day), solve_part1(input));
        println!("{} p2 {:#?}", stringify!($day), solve_part2(input));
    };
}

fn main() {
    run_solution!(day1);
}
