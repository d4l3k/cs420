# Question 1

## Heaviest-First

![](./q1p1.png)\


In the above diagram, the vertex with the heaviest weight of 5 is picked, and the
two adjacent vertices with weights 4 are discarded. This is not the maximum
independent set, since the sum is 5, instead of the set picking just the bottom
two vertices with sum 8.

## Even-Odd

![](./q1p2.png)\


The above diagram clearly shows that Even-Odd doesn't work.

## One that does work

We can solve this using dynamic programming. If we start from one side and
slowly build up, we can select the optimal set as we go. At every step we can
either include a node if it's independent, or skip it in favor of the next node.
This is similar to the largest subsequence problem.

TODO: flesh out

Ignore all vertices which have weights that are zero or negative.
