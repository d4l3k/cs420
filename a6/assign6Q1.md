# Question 1

## a) Prove that CLIQUE-3 is in NP.

Proof via the verifier-based definition of NP.

We can use the nodes in a clique as the certificate for the problem. We can then
create a verifier for whether those nodes are actually in a clique by visiting
each of them and counting the number of times each node is connected to another
node via an edge. Verifying each node is in polynomial time since there are a
maximum number of edges and nodes in a clique. There also can only be a
polynomial number of cliques in a graph due to the limit of 3 edges on each
vertex. Therefore, the verifier runs in polynomial time and CLIQUE-3 is in NP.

## b) What is wrong with the following proof of NP-completeness for CLIQUE-3?

This isn't actually proving anything about CLIQUE-3, it's doing the reduction
incorrectly. What this is saying is that CLIQUE is at least as hard as CLIQUE-3,
not the other way around.

## c) What's wrong with this other proof? VC-3 -> CLIQUE-3

It doesn't establish why "C ⊆ V is a vertex cover in G if and only if the
complementary step V-C is a clique in G.", and thus doesn't prove correctness of
the reduction.

## d) Describe an $O(|V|^4)$-time algorithm for CLIQUE-3.

Since all in CLIQUE-3, all vertices have a maximum of 3 edges, we know the
largest a CLIQUE can be is 4 vertices. Thus, we can brute force all possible
solutions of maximum size 4.


For every node we can recursively search for all nodes that are connected to the
original node via an edge. Each search is O(|V|) and there's a maximum of 3
possible connected nodes, that means finding all cliques for a given vertex is
$O(|V|^3)$. For all nodes, it becomes $O(|V|^4)$.
