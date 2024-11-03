function to_decimal(bits) {
    value=0
    for (i=0; i < length(bits); i++) {
        bit = int(bits[i])
        if (bit > 0) { value += 2 ^ (msb - i) }
    }
    return value;
}

{
    msb = 0
    split($0, chars, "")
    for (bit=1; bit <= length($0); ++bit) {
        if (bit in sums) { sums[bit] += chars[bit] }
        else           { sums[bit] = 0 }
        if (bit > msb) { msb = bit }
    }
    codes += 1
}

END {
    for (i in sums) {
        if (sums[i] / codes > 0.5) {
            gamma[i] = 1
            epsilon[i] = 0
        } else {
            gamma[i] = 0
            epsilon[i] = 1
        }
    }
    gamma_rate=to_decimal(gamma)
    epsilon_rate=to_decimal(epsilon)
    printf("%d\n", gamma_rate * epsilon_rate)
}
