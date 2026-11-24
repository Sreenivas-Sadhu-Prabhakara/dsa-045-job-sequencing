# Job Sequencing Problem

> Striver SDE Sheet 045 · Day 8 Greedy — take jobs richest-first, and run each on the latest free day inside its deadline. Go, stdlib only, tested.

## Problem

Each job takes one day, has a deadline (day number) and a profit paid only if it completes by its deadline. One job per day. Choose and order jobs to maximize total profit; report how many run and the profit.

## Approach

Two greedy choices stacked. First, consider jobs in descending profit — a richer job never deserves losing its slot to a poorer one. Second, place each accepted job on the **latest** still-free day within its deadline: procrastinating deliberately keeps earlier days open for jobs with tighter deadlines that haven't been placed yet. The latest-slot scan is linear here; swapping it for a union-find "next free day" structure drops placement to near-O(1) when deadlines are huge — noted, not needed at sheet scale.

**Complexity:** O(n log n + n·D) time (D = max deadline), O(D) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
