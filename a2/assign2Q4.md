# Question 4

Assumption: The input sequence M is sorted.

We need to find M+1 points S such that the midpoint of $s_is_{i+1}$ is $M_i$.
We can model this problem as a linear system with M+1 variables.
For each value M define a linear equation in the format
$0.5 S_i + 0.5 S_{i+1} = M_i$.

The only trick is setting the initial value. Since the points M are based off of
the midpoint, that means the distance between $m_1$ and $m_2$ will be less than
or equal to the distance between $s_1$ and $m_1$. Thus we can pick an arbitrary
point $s_1$ within that range and see if it forms a valid solution.

We can then structure this into matrices of the form $Ax=b$ and solve it using
matrix inversion with a runtime of $O(n^3)$ where n is the number of output
points or M+1.

This seems rather inefficient, but does solve it and can be effectively run on
GPUs.
