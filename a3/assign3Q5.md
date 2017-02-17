# Question 5


We can solve this problem by transforming the input graph into a bipartite graph
and then using Ford-Fulkerson to solve the bipartite matching.

First create two new sets of vertexes, one for each of vertexes in the input
graph. These will form the bipartite graph with one set on each side of the
partition. The next step is for every directed edge $\{i, j\}$ in the input
graph, create a new edge between the $i$th node in the first set, and the $j$th
node in the second set.

Finally, once we have the graph we can solve the bipartite matching using
Ford-Fulkerson as discussed in class. If there is a perfect matching, there is
cycle cover of the input graph formed by the matched edges. If there is no
perfect matching, report that no cycle cover exists.

The transformation step can be done in $\Theta(|V|+|E|)$ since it requires iterating
through each vertex and edge once. The solving of the bipartite matching is
$O(|V|*|E|)$ due to the invocation of Ford-Fulkerson.
