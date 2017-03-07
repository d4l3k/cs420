# Question 3

## Payoff matrix

|        | Nickel | Dime |
| ------ | ------ | ---- |
| Nickel | 5      | -5   |
| Dime   | -10    | 10   |

## Describe the optimal strategy for Row using a linear program.

We can take the above matrix to be $A$.

$x = [x_1 x_2 \ldots x_n]^T$, where Row picks $i$ with probability $x_i$.

$y = [y_1 y_2 \ldots y_n]^T$, where Col picks $j$ with probability $y_j$.

$x_i \geq 0$, $y_j \geq 0$

$\sum_i x_i = 1$, $\sum_j y_j = 1$

Row wants to maximize $y^TAx$.

## What is the optimal strategy?

The optimal strategy is to switch between the two 50/50 since it makes it impossible
to predict and there is a 50% chance of losing money.

## What is the optimal strategy if instead of a nickel or a dime the players hide a a-cent coin or a b-cent coin?

The values don't matter, since for every option Row picks, either possibility
adds up to zero. Thus, all that matters is whether Row wins or loses and
randomizing between the options minimizes the number of losses.
