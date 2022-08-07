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

Fast and slow pointers are used in problems where there is a cycle or when the the middle and
end of a data structure need to be determined where the length isn't known. The classic use of this
pattern is to determine if there is a cycle in the data structure as the two pointers will
eventually meet if there is a cycle.

Tweaks:

- Can be used to determine the start of a linked list cycle. Use fast and slow to determine if
  there is a cycle. Reset one of the pointers to point to the head of the list, then increment
  them both by one step each until they meet again. This meeting point is the start of the cycle.
  This works because one pointer moves twice the speed of the other, so when the slow one enters
  the cycle, the fast one is already at position k (where k is the length of the pre-cycle) in the
  cycle, so when they meet, it will be a position end/start of cycle - k. This is because moving
  two pointers in a cycle, one a 1x speed and one at 2x speed, they'll meet at the starting place,
  which due to the offset of +k of one pointer means the meeting position is -k (i.e. for the fast
  pointer it moves 2 cycles + 2k while the slow one moves one cycle, so they'll meet at the fast
  pointer's start position - 2k or position -k) . So resetting one pointer to the start and moving
  k steps for both pointers means they meet at the start of they cycle.

- Use this technique in a multipart problem to determine the middle and end positions of a linked
  list or graph that needs to be split or re-ordered.
  
Some good example problems:

- 🌼 [Linked List Cycle II](https://leetcode.com/problems/linked-list-cycle-ii/)
- 🌼 [Reorder List](https://leetcode.com/problems/reorder-list/)


## Merge Intervals

This pattern is used for problems that contain intervals that should be merged or modified so that
every interval in the set is disjoint. This usually requires that either the input to be sorted,
either by the an input constraint or by sorting as the first action taken to solve the problem. It
also requires determining which intervals overlap, which can be accomplished by:

 max(A[0], B[0]) - min(A[1], B[1]) >= 0 for two intervals A and B where A[0] = start and A[1] = end

The intervals need to be iterated over checking the neighboring intervals for overlap and combining
them when they do. If there is only one list of intervals, this is accomplished by creating a
second list of merged intervals containing the first one and a separate pointer for the current
position in the merged list. Then for i=1..n, each interval[i] is checked against the current
merged interval. If there is no overlap, interval[i] is inserted into the merged list and the
pointer is increased. If there is overlap, the current merged interval is extended to the end of
interval[i]. If there are two lists to be merged, two pointers should be used to point to the
current interval in each list, then whichever interval's end is lower than the other should be
incremented (i.e if A[iA][1] < B[iB][1] { iA++ } else { iB++ }).

Tweaks:

- Having to insert an interval into the set of intervals can be easily accomplished by appending
  the new interval to the set, sorting, then merging. Alternatively if the sets are already
  ordered, the new interval can be inserted into the set by iterating the set once to find the
  right position.
  
- "Difference Algorithm": For overlapping intervals that "add up" and have a known contiguous size,
  create a final values array of the right length initialized with zeros, then for each interval,
  add the value of the interval to the start of the interval and subtract the value from the end+1.
  Then iterate over the values array adding adding the previous value to the current (i.e.
  values[i] += values[i-1])

Some good example problems:

- 🌼 [Insert Interval](https://leetcode.com/problems/insert-interval/)
- 🌼 [Interval List Intersections](https://leetcode.com/problems/interval-list-intersections/)
- 🌼 [Corporate Flight Bookins](https://leetcode.com/problems/corporate-flight-bookings/)


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

