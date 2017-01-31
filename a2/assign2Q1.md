# Question 1

We know that solving ELEMENT UNIQUENESS takes $\Omega(n\log n)$ time. We must
show that the point depth problem has an efficient solution that has ELEMENT
UNIQUENESS as a subroutine to show that the whole problem has a $\Omega(n\log
n)$ bound.

Thus we need to reduce ELEMENT UNIQUENESS to the point depth problem.

To determine if there are unique elements using the max depth problem, one
solution would be to take all numbers and produce a series of convex hulls each
one with distance $N_i$ from the center. This will make each one of the input
numbers into a convex hull, and calculating the max depth of the set of convex
hulls will return how many unique points there are. Transforming the input
numbers into the concentric convex hulls is $\Theta(n)$ as you need a linear amount of
work for each number. Likewise, comparing the total number of points to the max
depth is also linear. Thus giving a lower bound for the max depth problem of

$$\Omega(n\log n) + \Theta(n) + \Theta(1) = \Omega(n\log n)$$
