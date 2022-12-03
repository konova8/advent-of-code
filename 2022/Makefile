all: day1/day1-1.out day1/day1-2.out day2/day2-1.out day2/day2-2.out day3/day3-1.out day3/day3-2.out day4/day4-1.out day4/day4-2.out day5/day5-1.out day5/day5-2.out day6/day6-1.out day6/day6-2.out day7/day7-1.out day7/day7-2.out day8/day8-1.out day8/day8-2.out day9/day9-1.out day9/day9-2.out day10/day10-1.out day10/day10-2.out day11/day11-1.out day11/day11-2.out day12/day12-1.out day12/day12-2.out day13/day13-1.out day13/day13-2.out day14/day14-1.out day14/day14-2.out day15/day15-1.out day15/day15-2.out day16/day16-1.out day16/day16-2.out day17/day17-1.out day17/day17-2.out day18/day18-1.out day18/day18-2.out day19/day19-1.out day19/day19-2.out day20/day20-1.out day20/day20-2.out day21/day21-1.out day21/day21-2.out day22/day22-1.out day22/day22-2.out day23/day23-1.out day23/day23-2.out day24/day24-1.out day24/day24-2.out day25/day25-1.out day25/day25-2.out

%.out: %.o
	gcc -o $@ $^

%.o: %.c
	gcc -c -o $@ $<

clean:
	rm -f */*.o */*.out 
