pub struct Solution;
impl Solution {
    // Finds the all roots that are minimum height trees of the given number of nodes and the edge
    // list. Uses the properties of centered trees, see https://en.wikipedia.org/wiki/Centered_tree.
    pub fn find_min_height_trees(n: i32, edges: Vec<Vec<i32>>) -> Vec<i32> {
        let n: usize = n.try_into().unwrap();
        let mut adjacent = vec![vec![]; n];

        for edge in edges {
            if edge.len() != 2 {
                continue;
            }

            let first: usize = edge[0].try_into().unwrap();
            let second: usize = edge[1].try_into().unwrap();

            adjacent[first].push(edge[1]);
            adjacent[second].push(edge[0]);
        }

        // To find the center of a tree, which is the minimum height root(s), first find the
        // farthest node from a random node, then find the farthest node from that node. The path
        // between the second and third nodes is the diameter of the tree and the center of that
        // path are the two minimimum height tree roots.
        let a = 0;
        let (b, _) = Self::farthest(a, &adjacent, &mut vec![false; n]);
        let (_, path) = Self::farthest(b, &adjacent, &mut vec![false; n]);

        let roots = if path.len() % 2 == 0 {
            vec![path[(path.len() - 1) / 2], path[(path.len() - 1) / 2 + 1]]
        } else {
            vec![path[(path.len() - 1) / 2]]
        };

        roots
    }

    // Recursively finds the farthest node from the root using the adjacency matrix. Returns both
    // the farthest node and the path between the root and the node.
    fn farthest(root: i32, adjacent: &Vec<Vec<i32>>, seen: &mut Vec<bool>) -> (i32, Vec<i32>) {
        let mut path = vec![];
        let mut node = root;

        let root_index: usize = root.try_into().unwrap();
        seen[root_index] = true;

        for child in adjacent[root_index].clone() {
            let index: usize = child.try_into().unwrap();
            if !seen[index] {
                let (c_node, c_path) = Self::farthest(child, adjacent, seen);
                if c_path.len() > path.len() {
                    path = c_path;
                    node = c_node;
                }
            }
        }

        path.push(root);
        (node, path)
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn basic() {
        assert_eq!(
            Solution::find_min_height_trees(2, vec![vec![1, 0]]),
            vec![0, 1]
        );
    }

    #[test]
    fn provided() {
        assert_eq!(
            Solution::find_min_height_trees(4, vec![vec![1, 0], vec![1, 2], vec![1, 3]]),
            vec![1]
        );

        assert_eq!(
            Solution::find_min_height_trees(
                6,
                vec![vec![3, 0], vec![3, 1], vec![3, 2], vec![3, 4], vec![5, 4]]
            ),
            vec![3, 4]
        );
    }
}
