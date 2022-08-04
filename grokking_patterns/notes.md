# Pattern Notes

## Sliding Window

These types of problems involve finding the maximum number of something with restrictions. Using
two pointers to represent the front and the back of the window that is "sliding" across the data
and keeping some information about the data seen between the two pointers is the general solution.
The window can remain a constant size, with the front pointer adding data and the back pointer
removing data as the window moves forward. The window can also be a variable size, with one pointer
always moving and the other moving when the window no longer satisfies the problem's conditions.
For example, the front pointer always increments, but when it encounters a "bad" position, the back
pointer is moved forward, modifying the stored data until the front pointer is again in a good
position. Storing of the "best" solution seen is usually required as well.

Some good example problems:

- 🌼 [Fruit Into Baskets](https://leetcode.com/problems/fruit-into-baskets/)
- 🌼 [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)
- 🌼 [Max Consecutive Ones III](https://leetcode.com/problems/max-consecutive-ones-iii/)


## Islands (Matrix Traversal)


## Two Pointers

This pattern uses two pointers that move through a data structure independently to always look for
more optimal solutions. Many of these problems first require the data to be sorted so that the
determination of which pointer is moved is based on either increasing or decreasing the running
value. The pointers can start on the same side or commonly one at the beginning and one at the end.
If more than two elements need to be tracked, such as 3sum or 4sum, then the other elements are
incremented by loops and the two pointers only work on data that comes after the looped elements.
Storing the "best" solution is also required.

Some good example problems:

- 🌼 [Subarry Product Less Than K](https://leetcode.com/problems/subarray-product-less-than-k/)
- 🌼 [Sort Colors](https://leetcode.com/problems/sort-colors/)
- 🌼 [4Sum](https://leetcode.com/problems/4sum/)


## Fast & Slow Pointers


## Merge Intervals


## Cyclic Sort


## In-place Reversal of a LinkedList


## Tree Breadth First Search


## Tree Depth First Search


## Two Heaps


## Subsets


## Modified Binary Search


## Bitwise XOR


## Top K Elements


## K-way merge


## 0/1 Knapsack


## Topological Sort


## Multi-threaded


## Misc

