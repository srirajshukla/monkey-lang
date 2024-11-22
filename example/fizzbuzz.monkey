let start = 1;
let end = 100;

let fizzbuzz = fn (i) {
    if (and(rem(i, 5) == 0, rem(i, 3) == 0 )) {
        return "FizzBuzz";
    }
    if (rem(i, 5) == 0) {
        return "Fizz";
    }
    if (rem(i, 3) == 0) {
        return "Buzz";
    }

    return "";
};

let solveFizzBuzz = fn (s, e) {
    if (s > e) {
        return "";
    }

    let v = fizzbuzz(s);

    if (!eq(v, "")) {
        puts(join(s, " = ", v));
    }
    return solveFizzBuzz(s+1, e);
};

puts(start, end);
solveFizzBuzz(start, end);