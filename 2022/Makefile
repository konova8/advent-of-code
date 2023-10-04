all: day1/day1-1.out day1/day1-2.out day2/day2-1.out day2/day2-2.out day3/day3-1.out day3/day3-2.out day4/day4-1.out day4/day4-2.out day6/day6-1.out day6/day6-2.out

%.out: %.o
	gcc -o $@ $^

%.o: %.c
	gcc -c -o $@ $<

clean:
	rm -f */*.o */*.out 
