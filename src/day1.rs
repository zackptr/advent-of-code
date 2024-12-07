use itertools::Itertools;

pub fn solve_part1(input: &str) -> usize {
    let (left, right): (Vec<usize>, Vec<usize>) = input
        .lines()
        .filter_map(|ln| ln.split_once("   "))
        .map(|(l, r)| (l.parse::<usize>().unwrap(), r.parse::<usize>().unwrap()))
        .unzip();

    let total_distance: usize = (left.iter().sorted())
        .zip(right.iter().sorted())
        .map(|(n, m)| n.abs_diff(*m))
        .sum();

    total_distance
}

pub fn solve_part2(input: &str) -> usize {
    let (left, right): (Vec<usize>, Vec<usize>) = input
        .lines()
        .filter_map(|ln| ln.split_once("   "))
        .map(|(l, r)| (l.parse::<usize>().unwrap(), r.parse::<usize>().unwrap()))
        .unzip();

    let similarity_score: usize = left
        .iter()
        .map(|n| n * right.iter().filter(|m| *m == n).count())
        .sum();

    similarity_score
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solve_part1_example() {
        let example = include_str!("../examples/day1.txt");

        assert!(solve_part1(example) == 11);
    }

    #[test]
    fn test_solve_part2_example() {
        let example = include_str!("../examples/day1.txt");

        assert!(solve_part2(example) == 31);
    }
}
