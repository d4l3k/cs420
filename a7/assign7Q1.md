# Question 1

## (a) Prove that this heuristic returns a vertex cover of G.

By definition, a spanning tree includes all vertices in a graph. Thus, it is a
vertex cover. All leaf nodes are connected via an edge to a non-leaf node. Thus,
if the non-leaf node is part of the vertex cover, the leaf edge is covered since
they're connected by an edge. We must show that no two leaf nodes are connected
to each other via an edge. For two leaf nodes to be connected to each other, the
depth first search must have stopped at both nodes. One search path must have
stopped first, with the other still unvisited, however, depth first search will
expand to all possible paths before stopping. Thus, search wouldn't have stopped
with an unvisited node connected via an edge. Thus, there cannot be two leaf
nodes connected to each other. Therefore, if all non-leaf nodes of a spanning
tree are part of the vertex cover, it forms a complete vertex cover.

## (b) Prove that this heuristic returns a 2-approximation to the minimum vertex cover of G.

We need to show that $c(G) \leq 2*c_{OPT}(g)$.

For one essential property of vertex covers is that every edge must have a one
of it's vertices as part of the vertex cover. Leaf nodes of the graph, will
never be picked over non-leaf nodes of a graph due to being degree one. Thus,
the minimum vertex cover must be at least sized $\frac n2$ where $n=$number of
non-leaf nodes. Since the vertex cover produced by the heuristic will have a
maximum vertex cover size of $n$ by definition, it must be a 2-approximation to
the minimum vertex cover of G.

## (c) Describe an infinite family of graphs for which this heuristic returns a vertex cover of size 2 · OPT.

The family of all circular paths with an even number of nodes returns a vertex
cover of size $2*OPT$. Since there's no leaf nodes, that means that the vertex
cover formed by the heuristic is sized $n$. The optimal vertex cover in a path
uses every other node. Assuming an even $n$, that means the optimal vertex cover
is $i=\{2,4,6,\ldots,n\}$, with a total node count of $\frac{n}{2}$.

$\frac{n}{\frac{n}{2}} = 2$

