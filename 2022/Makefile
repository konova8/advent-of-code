all: day01/day01-1.out day01/day01-2.out day02/day02-1.out day02/day02-2.out day03/day03-1.out day03/day03-2.out day04/day04-1.out day04/day04-2.out day06/day06-1.out day06/day06-2.out

%.out: %.o
	gcc -o $@ $^

%.o: %.c
	gcc -c -o $@ $<

clean:
	rm -f */*.o */*.out 
