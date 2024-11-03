use std::collections::HashSet;
use std::fs::File;
use std::io::{self, BufRead};
use std::path::{Path, PathBuf};

fn resolve_path(path: &str) -> Result<PathBuf, io::Error> {
    let root = std::env::current_dir()?;
    Ok(root.join(path))
}

fn read_numbers(path: &Path) -> Result<Vec<i32>, io::Error> {
    let file = File::open(path)?;
    io::BufReader::new(file)
        .lines()
        .enumerate()
        .map(|(n, line)| {
            line.map_err(|e| {
                io::Error::new(
                    io::ErrorKind::InvalidData,
                    format!("Error reading line {}: {}", n + 1, e),
                )
            })?
            .parse::<i32>()
            .map_err(|e| {
                io::Error::new(
                    io::ErrorKind::InvalidData,
                    format!("Error parsing line {}: {}", n + 1, e),
                )
            })
        })
        .collect()
}

fn p1<T>(iter: T, sum: i32) -> Option<i32>
where
    T: Iterator<Item = i32>,
{
    let mut set = HashSet::new();
    for n in iter {
        let complement = sum - n;
        if set.contains(&complement) {
            return Some(n * complement);
        }
        set.insert(n);
    }
    None
}

fn p2(numbers: &[i32], sum: i32) -> Option<i32> {
    for (i, n) in numbers.iter().enumerate() {
        let rest = numbers
            .iter()
            .enumerate()
            .filter_map(|(j, &v)| if i != j { Some(v) } else { None });
        if let Some(result) = p1(rest, sum - n) {
            return Some(n * result);
        }
    }
    None
}

const TOTAL: i32 = 2020;

fn main() -> io::Result<()> {
    let example = resolve_path("input.txt")?;
    let numbers = read_numbers(&example)?;

    if let Some(result) = p1(numbers.iter().cloned(), TOTAL) {
        println!("P1: {}", result);
    } else {
        println!("P1: No solution");
    };

    if let Some(result) = p2(&numbers, TOTAL) {
        println!("P2: {}", result);
    } else {
        println!("P2: No solution");
    };

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_p1() {
        let numbers = vec![1721, 979, 366, 299, 675, 1456];
        assert_eq!(p1(numbers.iter().cloned(), TOTAL), Some(514579));
    }

    #[test]
    fn test_p2() {
        let numbers = vec![1721, 979, 366, 299, 675, 1456];
        assert_eq!(p2(&numbers, TOTAL), Some(241861950));
    }
}
