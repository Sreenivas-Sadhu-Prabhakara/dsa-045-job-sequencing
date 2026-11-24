package main

import "testing"

func TestSchedule(t *testing.T) {
	cases := []struct {
		jobs                  []Job
		wantCount, wantProfit int
	}{
		{[]Job{{1, 4, 20}, {2, 1, 10}, {3, 1, 40}, {4, 1, 30}}, 2, 60},
		{[]Job{{1, 2, 100}, {2, 1, 19}, {3, 2, 27}, {4, 1, 25}, {5, 1, 15}}, 2, 127},
		{[]Job{{1, 1, 5}}, 1, 5},
		{[]Job{}, 0, 0},
		{[]Job{{1, 3, 10}, {2, 3, 20}, {3, 3, 30}}, 3, 60}, // all fit
	}
	for _, c := range cases {
		count, profit := Schedule(c.jobs)
		if count != c.wantCount || profit != c.wantProfit {
			t.Errorf("Schedule(%v) = (%d,%d), want (%d,%d)", c.jobs, count, profit, c.wantCount, c.wantProfit)
		}
	}
}
