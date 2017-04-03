# Question 1

## (a) Prove that this heuristic returns a vertex cover of G.

By definition, a spanning tree includes all vertices in a graph. Thus, it is a
vertex cover. All leaf nodes are connected via an edge to a non-leaf node. Thus,
if the non-leaf node is part of the vertex cover, the leaf node is covered since
they're connected by an edge. Therefore, if all non-leaf nodes of a spanning
tree are part of the vertex cover, it forms a complete vertex cover.

## (b) Prove that this heuristic returns a 2-approximation to the minimum vertex cover of G.



## (c) Describe an infinite family of graphs for which this heuristic returns a vertex cover of size 2 · OPT.

The family of all circular paths with an even number of nodes returns a vertex
cover of size $2*OPT$. Since there's no leaf nodes, that means that the vertex
cover formed by the heuristic is sized $n$. The optimal vertex cover in a path
uses every other node. Assuming an even $n$, that means the optimal vertex cover
is $i=\{2,4,6,\ldots,n\}$, with a total node count of $\frac{n}{2}$.

$\frac{n}{\frac{n}{2}} = 2$

