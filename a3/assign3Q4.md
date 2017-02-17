# Question 4

## (a) Give an example of a matrix M that is not rearrangeable even through every column and every row contains at least one entry that equals 1.

$$\begin{bmatrix}
1 & 0 & 0 \
1 & 0 & 0 \
0 & 0 & 1
\end{bmatrix}$$


## (b) Describe an efficient algorithm that determines if M is rearrangeable.

We can use bipartite matching to determine if $M$ is rearrangeable. One way to do
this is to create a graph with one vertex for every column and one vertex for
every row. You then add edges for every element in $M$ that equals $1$ connecting
the $i$th row vertex with the $j$th column vertex. If the number of bipartite
matches equals $n$, $M$ is rearrangeable.

Since the number of vertexes is $2n$, and the worst case number of edges is
$O(n^2)$ since every element in the matrix could be $1$.
Finding a bipartite matching requires solving a max flow problem which is
$O(|V|*|E|)$. Thus, this algorithm has a worst case of $O(n^3)$ which is in
polynomial time.

