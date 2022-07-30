// Returns an index with the key or None in log(n) time.
pub fn binary_search<T: PartialEq + PartialOrd>(
    items: &[T],
    key: T,
    low: usize,
    high: usize,
) -> Option<usize> {
    if low > high {
        return None;
    }

    let middle = (low + high) / 2;

    println!("low {low} high {high} middle {middle}");

    if items[middle] == key {
        return Some(middle);
    } else if items[middle] > key {
        return binary_search(items, key, low, middle - 1);
    } else {
        return binary_search(items, key, middle + 1, high);
    }
}

// Returns the index of the right boundary of the key in log(n) time.
// Note: The high value is returned when low is more than high. When the middle value is equal
// to the key the low value is moved right of the middle.
pub fn bs_right<T: PartialEq + PartialOrd>(items: &[T], key: T, low: i32, high: i32) -> i32 {
    if low > high {
        return high;
    }

    let middle = (low + high) / 2;

    if items[middle as usize] > key {
        return bs_right(items, key, low, middle - 1);
    } else {
        return bs_right(items, key, middle + 1, high);
    }
}

// Returns the index of the left boundary of the key in log(n) time.
// Note: The low value is returned when high is less then low. When the middle value is equal
// to the key the high value is moved left of the middle.
pub fn bs_left<T: PartialEq + PartialOrd>(items: &[T], key: T, low: i32, high: i32) -> i32 {
    if high < low {
        return low;
    }

    let middle = (low + high) / 2;

    if items[middle as usize] < key {
        return bs_left(items, key, middle + 1, high);
    } else {
        return bs_left(items, key, low, middle - 1);
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn boundaries() {
        let data = vec![1, 2, 2, 2, 3, 4, 5, 5, 5, 5, 5, 6, 7];
        assert_eq!(bs_left(&data, 2, 0, (data.len() - 1) as i32), 1);
        assert_eq!(bs_right(&data, 2, 0, (data.len() - 1) as i32), 3);
        assert_eq!(bs_left(&data, 3, 0, (data.len() - 1) as i32), 4);
        assert_eq!(bs_right(&data, 3, 0, (data.len() - 1) as i32), 4);
        assert_eq!(bs_left(&data, 1, 0, (data.len() - 1) as i32), 0);
        assert_eq!(bs_right(&data, 1, 0, (data.len() - 1) as i32), 0);
        assert_eq!(bs_left(&data, 7, 0, (data.len() - 1) as i32), 12);
        assert_eq!(bs_right(&data, 7, 0, (data.len() - 1) as i32), 12);
        assert_eq!(bs_left(&data, 9, 0, (data.len() - 1) as i32), 13);
        assert_eq!(bs_right(&data, 9, 0, (data.len() - 1) as i32), 12);
        assert_eq!(bs_left(&data, -1, 0, (data.len() - 1) as i32), 0);
        assert_eq!(bs_right(&data, -1, 0, (data.len() - 1) as i32), -1);
    }
}
