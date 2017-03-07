# Question 2

## (a) Give a linear-programming formulation of the maximum bipartite matching problem.

$\max \sum x$ where $x$ is a vector representing the edges with $1$ being
selected, and $0$ being not selected.

Sum of all adjacent edges is less than 1.

$x_{uv} + x_{vw} \leq 1$ for all $(u, v), (v, w) \in E$.

## (b) Now dualize the linear program from part (a). What do the dual variables represent?  What does the objective function represent?


### Primal problem

Maximize $c^Tx$ subject to $Ax\leq b, x \geq 0$.

$c=[1 1 \ldots 1]$ for all values in x.

$A$ is the $n \times n$ matrix formed by $[\ldots, x_{uv},  \ldots, x_{vw}, \ldots]$
for all $(u, v), (v, w) \in E$.

$b=[1 1 \ldots 1]$.

### Dual problem

Minimize $b^Ty$ subject to $A^Ty \geq c$, $y \geq 0$.

Two edges are an adjacency pair if they are connected by a vertex.

$A^T_{ij}$ is $1$ for each adjacency pair $j$ that edge $i$ is in.

$A^Ty$ computes the number of times an adjacency pair is picked for every edge.
Since $A^Ty \geq c$, that means that every edge has to be selected as part of an
adjacency pair at least once. Minimizing $b^Ty$, minimizes the number of pairs
that are selected.
